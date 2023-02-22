package app

import (
	"fmt"
	"mgr/data"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type DBHandler struct {
	db     *sqlx.DB
	dbtype string
	name   string
	demo   DemoHandler
}

var (
	info_tables          = []string{"tableinfo", "ontuneinfo", "agentinfo", "hostinfo"}
	metric_single_tables = []string{"lastrealtimeperf", "lastperf", "deviceid", "descid", "dfnameid", "lvnameid"}
	metric_ref_tables    = []string{"proccmd", "procuserid", "procargid"}
	metric_avg_tables    = []string{"avgperf", "avgcpu", "avgdisk", "avgnet", "avgpid", "avgproc", "avgmaxperf", "avgmaxcpu", "avgmaxdisk", "avgmaxnet", "avgmaxpid", "avgmaxproc", "avgdf"}
	metric_tables        = []string{"realtimeperf", "realtimecpu", "realtimedisk", "realtimenet", "realtimepid", "realtimeproc"}
)

func (d *DBHandler) CheckTable() {
	d.CheckTableInfo()

	if d.dbtype == "pg" {
		d.CheckTableRef()
		d.CheckTableMetric()
		d.CheckTableAvgMetric()
	}
}

func (d *DBHandler) CheckTableInfo() {
	single_tables := make([]string, 0)
	single_tables = append(single_tables, info_tables...)
	single_tables = append(single_tables, metric_single_tables...)

	for _, tb := range single_tables {
		var exist_count int
		err := d.db.QueryRow("SELECT COUNT(*) FROM pg_tables where tablename=$1", tb).Scan(&exist_count)
		ErrorFatal(err)

		if exist_count == 0 {
			tx := d.db.MustBegin()
			switch tb {
			case "tableinfo":
				tx.MustExec(data.TableinfoStmt)
			case "ontuneinfo":
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.OntuneinfoStmt, "int8"))
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.OntuneinfoStmt, "timestamptz"))
				}
				tx.MustExec(data.InsertOntuneinfoStmt, getOntuneinfo(d.dbtype)...)
			case "agentinfo":
				tx.MustExec(data.AgentinfoStmt)
			case "hostinfo":
				tx.MustExec(data.HostinfoStmt)
			case "lastrealtimeperf":
				tx.MustExec(data.LastrealtimeperfStmt)
			case "lastperf":
				tx.MustExec(data.LastperfStmt)
			default:
				tx.MustExec(fmt.Sprintf(data.RefStmt, tb))
			}
			tx.MustExec(data.InsertTableinfoStmt, tb, time.Now().Unix())
			LogWrite("log", fmt.Sprintf("%s %s table creation is completed\n", d.name, tb))
			err = tx.Commit()
			ErrorTx(err, tx)
		}
	}

	if d.dbtype == "ts" {
		for _, tbname := range metric_ref_tables {
			if !d.existTableinfo(tbname) {
				_, err := d.db.Exec(fmt.Sprintf(data.RefStmt, tbname))
				ErrorPq(err)

				tx := d.db.MustBegin()
				tx.MustExec(data.InsertTableinfoStmt, tbname, time.Now().Unix())
				LogWrite("log", fmt.Sprintf("%s %s table creation is completed\n", d.name, tbname))
				err = tx.Commit()
				ErrorTx(err, tx)
			}
		}

		mtables := metric_tables
		mtables = append(mtables, metric_avg_tables...)

		for _, tb := range mtables {
			var exist_count int
			err := d.db.QueryRow("SELECT COUNT(*) FROM pg_tables where tablename=$1", tb).Scan(&exist_count)
			ErrorFatal(err)

			if exist_count == 0 {
				tx := d.db.MustBegin()
				var ontunetime_type = "timestamptz"
				if tb[len(tb)-4:] == "perf" {
					tx.MustExec(fmt.Sprintf(data.MetricPerfStmt, tb, ontunetime_type))
				} else if tb[len(tb)-3:] == "cpu" {
					tx.MustExec(fmt.Sprintf(data.MetricCpuStmt, tb, ontunetime_type))
				} else if tb[len(tb)-4:] == "disk" {
					tx.MustExec(fmt.Sprintf(data.MetricDiskStmt, tb, ontunetime_type))
				} else if tb[len(tb)-3:] == "net" {
					tx.MustExec(fmt.Sprintf(data.MetricNetStmt, tb, ontunetime_type))
				} else if tb[len(tb)-3:] == "pid" {
					tx.MustExec(fmt.Sprintf(data.MetricPidStmt, tb, ontunetime_type))
				} else if tb[len(tb)-4:] == "proc" {
					tx.MustExec(fmt.Sprintf(data.MetricProcStmt, tb, ontunetime_type))
				} else if tb[len(tb)-2:] == "df" {
					tx.MustExec(fmt.Sprintf(data.MetricDfStmt, tb, ontunetime_type))
				}

				tx.MustExec(fmt.Sprintf(data.MetricHypertable, tb))
				tx.MustExec(data.InsertTableinfoStmt, tb, time.Now().Unix())
				LogWrite("log", fmt.Sprintf("%s %s table creation is completed\n", d.name, tb))
				err = tx.Commit()
				ErrorTx(err, tx)
			}
		}
	}
}

func (d *DBHandler) CheckTableRef() {
	d.createRefTables(false)
}

func (d *DBHandler) CheckTableAvgMetric() {
	d.createMetricTables(false, "day")
}

func (d *DBHandler) CheckTableMetric() {
	d.createMetricTables(false, "hour")
}

func (d *DBHandler) CheckDailyTableRef() {
	var rTableCreate bool
	for {
		hour := time.Now().Hour()
		if hour >= 23 && !rTableCreate {
			rTableCreate = true
			d.createRefTables(true)
			d.createMetricTables(true, "day")
		} else if hour < 23 {
			rTableCreate = false
		}
		time.Sleep(time.Second * time.Duration(10))
	}
}

func (d *DBHandler) createRefTables(check_flag bool) {
	for _, tbname := range metric_ref_tables {
		tablename := d.getTableName("day", tbname, check_flag)
		LogWrite("log", fmt.Sprintf("%s\n", tablename))
		if !d.existTableinfo(tablename) {
			_, err := d.db.Exec(fmt.Sprintf(data.RefStmt, tablename))
			ErrorPq(err)

			tx := d.db.MustBegin()
			tx.MustExec(data.InsertTableinfoStmt, tablename, time.Now().Unix())
			LogWrite("log", fmt.Sprintf("%s %s table creation is completed\n", d.name, tablename))
			err = tx.Commit()
			ErrorTx(err, tx)
		}
	}
}

func (d *DBHandler) CheckHourTableMetric() {
	var mTableCreate bool
	for {
		min := time.Now().Minute()
		if min >= 50 && !mTableCreate {
			mTableCreate = true
			d.createMetricTables(true, "hour")
		} else if min < 50 {
			mTableCreate = false
		}
		time.Sleep(time.Second * time.Duration(10))
	}
}

func (d *DBHandler) createMetricTables(check_flag bool, metric_type string) {
	var tables []string
	if metric_type == "day" {
		tables = metric_avg_tables
	} else {
		tables = metric_tables
	}
	for _, tbname := range tables {
		tablename := d.getTableName(metric_type, tbname, check_flag)
		LogWrite("log", fmt.Sprintf("%s\n", tablename))
		if !d.existTableinfo(tablename) {
			var ontunetime_type string = "integer"

			if tbname[len(tbname)-4:] == "perf" {
				_, err := d.db.Exec(fmt.Sprintf(data.MetricPerfStmt, tablename, ontunetime_type))
				ErrorPq(err)
			} else if tbname[len(tbname)-3:] == "cpu" {
				_, err := d.db.Exec(fmt.Sprintf(data.MetricCpuStmt, tablename, ontunetime_type))
				ErrorPq(err)
			} else if tbname[len(tbname)-4:] == "disk" {
				_, err := d.db.Exec(fmt.Sprintf(data.MetricDiskStmt, tablename, ontunetime_type))
				ErrorPq(err)
			} else if tbname[len(tbname)-3:] == "net" {
				_, err := d.db.Exec(fmt.Sprintf(data.MetricNetStmt, tablename, ontunetime_type))
				ErrorPq(err)
			} else if tbname[len(tbname)-3:] == "pid" {
				_, err := d.db.Exec(fmt.Sprintf(data.MetricPidStmt, tablename, ontunetime_type))
				ErrorPq(err)
			} else if tbname[len(tbname)-4:] == "proc" {
				_, err := d.db.Exec(fmt.Sprintf(data.MetricProcStmt, tablename, ontunetime_type))
				ErrorPq(err)
			} else if tbname[len(tbname)-2:] == "df" {
				_, err := d.db.Exec(fmt.Sprintf(data.MetricDfStmt, tablename, ontunetime_type))
				ErrorPq(err)
			}

			tx := d.db.MustBegin()
			tx.MustExec(data.InsertTableinfoStmt, tablename, time.Now().Unix())
			LogWrite("log", fmt.Sprintf("%s %s table creation is completed\n", d.name, tablename))
			err := tx.Commit()
			ErrorTx(err, tx)

		}
	}
}

func (d *DBHandler) getTableName(tabletype string, tablename string, check_flag bool) string {
	if d.dbtype == "pg" {
		now := time.Now()
		var format string
		if tabletype == "hour" {
			format = HOURLY_DATE_FORMAT
		} else if tabletype == "day" {
			format = DAILY_DATE_FORMAT
		}

		var createHour string
		if check_flag {
			createHour = now.Add(time.Hour * time.Duration(1)).Format(format)
		} else {
			createHour = now.Format(format)
		}
		metrictablename := tablename + "_" + createHour
		return metrictablename
	} else if d.dbtype == "ts" {
		return tablename
	}
	return ""
}

func (d *DBHandler) existTableinfo(tablename string) bool {
	var cnt int
	err := d.db.QueryRow(data.SELECT_COUNT_TABLEINFO, tablename).Scan(&cnt)
	ErrorPq(err)

	if cnt == 1 {
		return true
	} else {
		return false
	}
}

func (d *DBHandler) UpdateOntuneinfo() {
	ticker := time.NewTicker(time.Second * time.Duration(1))
	for range ticker.C {
		tx := d.db.MustBegin()

		if d.dbtype == "pg" {
			GlobalOntunetime = time.Now().Unix()
			tx.MustExec(data.UpdateOntuneinfoStmt, GlobalOntunetime)
		} else if d.dbtype == "ts" {
			GlobalOntunetimets = time.Now()
			tx.MustExec(data.UpdateOntuneinfoStmt, GlobalOntunetimets)
		}
		err := tx.Commit()
		ErrorTx(err, tx)
	}
}

func (d *DBHandler) InsertAgentData(agent_data *sync.Map, wg *sync.WaitGroup, idx int) {
	agent_data.Range(func(key, value any) bool {
		if key == "lastperf" || key == "lastrealtimeperf" {
			return true
		}

		val := value.(*sync.Map)
		if value == nil || contains(key) || GetMapSize(val) == 0 {
			return true
		}

		// Copy and Init
		insert_agent_data := CopyMap(val)
		agent_data.LoadAndDelete(key)

		k := key.(string)
		if len(k) >= 12 && k[:12] == "realtimeperf" {
			var realtimeperf_arr data.PerfArr = data.PerfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Perf)
				realtimeperf_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()
			args := realtimeperf_arr.GetArgs(d.dbtype)

			if GetMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertPerfPg, d.getTableName("hour", "realtimeperf", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertPerfTs, d.getTableName("hour", "realtimeperf", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("debug", fmt.Sprintf("rperf %d %s %s %d %d", idx, k, d.dbtype, len(realtimeperf_arr.Ontunetime), GlobalOntunetime))
			}
		} else if len(k) >= 11 && k[:11] == "realtimepid" {
			var realtimepid_arr data.PidArr = data.PidArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Pid)
				realtimepid_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()
			args := realtimepid_arr.GetArgs(d.dbtype)

			if GetMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertPidPg, d.getTableName("hour", "realtimepid", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertPidTs, d.getTableName("hour", "realtimepid", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("debug", fmt.Sprintf("rpid %d %s %s %d %d", idx, k, d.dbtype, len(realtimepid_arr.Ontunetime), GlobalOntunetime))
			}
		} else if len(k) >= 12 && k[:12] == "realtimeproc" {
			var realtimeproc_arr data.ProcArr = data.ProcArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Proc)
				realtimeproc_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()
			args := realtimeproc_arr.GetArgs(d.dbtype)

			if GetMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertProcPg, d.getTableName("hour", "realtimeproc", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertProcTs, d.getTableName("hour", "realtimeproc", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("debug", fmt.Sprintf("rproc %d %s %s %d %d", idx, k, d.dbtype, len(realtimeproc_arr.Ontunetime), GlobalOntunetime))
			}
		} else if len(k) >= 12 && k[:12] == "realtimedisk" {
			var realtimedisk_arr data.DiskArr = data.DiskArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Disk)
				realtimedisk_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()
			args := realtimedisk_arr.GetArgs(d.dbtype)

			if GetMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertDiskPg, d.getTableName("hour", "realtimedisk", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertDiskTs, d.getTableName("hour", "realtimedisk", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("debug", fmt.Sprintf("rdisk %d %s %s %d %d", idx, k, d.dbtype, len(realtimedisk_arr.Ontunetime), GlobalOntunetime))
			}
		} else if len(k) >= 11 && k[:11] == "realtimenet" {
			var realtimenet_arr data.NetArr = data.NetArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Net)
				realtimenet_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()
			args := realtimenet_arr.GetArgs(d.dbtype)

			if GetMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertNetPg, d.getTableName("hour", "realtimenet", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertNetTs, d.getTableName("hour", "realtimenet", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("debug", fmt.Sprintf("rnet %d %s %s %d %d", idx, k, d.dbtype, len(realtimenet_arr.Ontunetime), GlobalOntunetime))
			}
		} else if len(k) >= 11 && k[:11] == "realtimecpu" {
			var realtimecpu_arr data.CpuArr = data.CpuArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Cpu)
				realtimecpu_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()
			args := realtimecpu_arr.GetArgs(d.dbtype)

			if GetMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertCpuPg, d.getTableName("hour", "realtimecpu", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertCpuTs, d.getTableName("hour", "realtimecpu", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("debug", fmt.Sprintf("rcpu %d %s %s %d %d", idx, k, d.dbtype, len(realtimecpu_arr.Ontunetime), GlobalOntunetime))
			}
		} else if len(k) >= 7 && k[:7] == "avgperf" {
			var avgperf_arr data.PerfArr = data.PerfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Perf)
				avgperf_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertPerfPg, d.getTableName("day", "avgperf", false)), avgperf_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertPerfTs, d.getTableName("day", "avgperf", false)), avgperf_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("aperf %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 6 && k[:6] == "avgpid" {
			var avgpid_arr data.PidArr = data.PidArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Pid)
				avgpid_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertPidPg, d.getTableName("day", "avgpid", false)), avgpid_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertPidTs, d.getTableName("day", "avgpid", false)), avgpid_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("apid %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 7 && k[:7] == "avgproc" {
			var avgproc_arr data.ProcArr = data.ProcArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Proc)
				avgproc_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertProcPg, d.getTableName("day", "avgproc", false)), avgproc_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertProcTs, d.getTableName("day", "avgproc", false)), avgproc_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("aproc %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 5 && k[:5] == "avgdf" {
			var avgdf_arr data.DfArr = data.DfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Df)
				avgdf_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertDfPg, d.getTableName("day", "avgdf", false)), avgdf_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertDfTs, d.getTableName("day", "avgdf", false)), avgdf_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("adf %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 7 && k[:7] == "avgdisk" {
			var avgdisk_arr data.DiskArr = data.DiskArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Disk)
				avgdisk_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertDiskPg, d.getTableName("day", "avgdisk", false)), avgdisk_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertDiskTs, d.getTableName("day", "avgdisk", false)), avgdisk_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("adisk %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 6 && k[:6] == "avgnet" {
			var avgnet_arr data.NetArr = data.NetArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Net)
				avgnet_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertNetPg, d.getTableName("day", "avgnet", false)), avgnet_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertNetTs, d.getTableName("day", "avgnet", false)), avgnet_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("anet %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 6 && k[:6] == "avgcpu" {
			var avgcpu_arr data.CpuArr = data.CpuArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Cpu)
				avgcpu_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertCpuPg, d.getTableName("day", "avgcpu", false)), avgcpu_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertCpuTs, d.getTableName("day", "avgcpu", false)), avgcpu_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("acpu %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 10 && k[:10] == "avgmaxperf" {
			var avgperf_arr data.PerfArr = data.PerfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Perf)
				avgperf_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertPerfPg, d.getTableName("day", "avgmaxperf", false)), avgperf_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertPerfTs, d.getTableName("day", "avgmaxperf", false)), avgperf_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("amperf %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 9 && k[:9] == "avgmaxpid" {
			var avgpid_arr data.PidArr = data.PidArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Pid)
				avgpid_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertPidPg, d.getTableName("day", "avgmaxpid", false)), avgpid_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertPidTs, d.getTableName("day", "avgmaxpid", false)), avgpid_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("ampid %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 10 && k[:10] == "avgmaxproc" {
			var avgproc_arr data.ProcArr = data.ProcArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Proc)
				avgproc_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertProcPg, d.getTableName("day", "avgmaxproc", false)), avgproc_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertProcTs, d.getTableName("day", "avgmaxproc", false)), avgproc_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("amproc %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 10 && k[:10] == "avgmaxdisk" {
			var avgdisk_arr data.DiskArr = data.DiskArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Disk)
				avgdisk_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertDiskPg, d.getTableName("day", "avgmaxdisk", false)), avgdisk_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertDiskTs, d.getTableName("day", "avgmaxdisk", false)), avgdisk_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("amdisk %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 9 && k[:9] == "avgmaxnet" {
			var avgnet_arr data.NetArr = data.NetArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Net)
				avgnet_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertNetPg, d.getTableName("day", "avgmaxnet", false)), avgnet_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertNetTs, d.getTableName("day", "avgmaxnet", false)), avgnet_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("amnet %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 9 && k[:9] == "avgmaxcpu" {
			var avgcpu_arr data.CpuArr = data.CpuArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Cpu)
				avgcpu_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertCpuPg, d.getTableName("day", "avgmaxcpu", false)), avgcpu_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertCpuTs, d.getTableName("day", "avgmaxcpu", false)), avgcpu_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("debug", fmt.Sprintf("amcpu %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		}

		time.Sleep(time.Millisecond * time.Duration(1))
		return true
	})

	defer func() {
		if c := recover(); c != nil {
			fmt.Println("recover data")
		}
		//LogWrite("debug", fmt.Sprintf("waitgroup %d", idx))
		wg.Done()
	}()
}

func (d *DBHandler) InsertLastPerfData(agent_data *sync.Map, wg *sync.WaitGroup) {
	var lastperf_arr data.LastperfArr = data.LastperfArr{}
	agent_data.Range(func(key, value any) bool {
		lpdata := value.(*data.Lastperf)
		lastperf_arr.SetData(*lpdata)

		return true
	})

	tx := d.db.MustBegin()
	tx.MustExec(data.TruncateLastperf)
	tx.MustExec(data.InsertLastperf, lastperf_arr.GetArgs()...)
	err := tx.Commit()
	ErrorTx(err, tx)

	//LogWrite("debug", "waitgroup -2")
	defer wg.Done()
}

func (d *DBHandler) InsertLastRealtimePerfData(agent_data *sync.Map, wg *sync.WaitGroup) {
	if lrtp_data, ok := agent_data.Load("lastrealtimeperf"); ok && GetMapSize(lrtp_data.(*sync.Map)) > 0 {
		var lastrealtimeperf_arr data.LastrealtimeperfArr = data.LastrealtimeperfArr{}
		lrtp_data.(*sync.Map).Range(func(key, value any) bool {
			lrtpdata := value.(*data.Lastrealtimeperf)
			lastrealtimeperf_arr.SetData(*lrtpdata)

			return true
		})

		tx := d.db.MustBegin()
		tx.MustExec(data.TruncateLastrealtimeperf)
		tx.MustExec(data.InsertLastrealtimeperf, lastrealtimeperf_arr.GetArgs()...)
		err := tx.Commit()
		//LogWrite("debug", fmt.Sprintf("lrtp %s %d", d.dbtype, GlobalOntunetime))
		ErrorTx(err, tx)
	}
	//LogWrite("debug", "waitgroup -1")
	defer wg.Done()
}

func (d *DBHandler) InitExecAgentHostInfo() {
	var exist_count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM agentinfo where _enabled=1 and _agentname like 'Dummy%'").Scan(&exist_count)
	ErrorFatal(err)

	if exist_count < d.demo.Hostcount {
		tx := d.db.MustBegin()
		tx.MustExec(data.DeleteAgentinfoDummy)
		tx.MustExec(data.DeleteHostinfoDummy)
		tx.MustExec(data.InsertAgentinfo, d.demo.Agentinfo.GetArgs()...)
		tx.MustExec(data.InsertHostinfo, d.demo.Hostinfo.GetArgs()...)
		err = tx.Commit()
		ErrorTx(err, tx)
	}
}

func (d *DBHandler) InitDeviceid() {
	var exist_count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM deviceid").Scan(&exist_count)
	ErrorFatal(err)

	if exist_count < len(data.DEVICE_IDS) {
		keys := make([]int, 0)
		for k := range data.DEVICE_IDS {
			keys = append(keys, k)
		}

		tx := d.db.MustBegin()
		tx.MustExec(fmt.Sprintf(data.TruncateRef, "deviceid"))
		tx.MustExec(fmt.Sprintf(data.InsertRef, "deviceid"), pq.Array(keys), pq.StringArray(data.DEVICE_IDS))
		err = tx.Commit()
		ErrorTx(err, tx)
	}
}

func (d *DBHandler) InitDescid() {
	var exist_count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM descid").Scan(&exist_count)
	ErrorFatal(err)

	if exist_count < len(data.VOLUME_GROUPS) {
		keys := make([]int, 0)
		for k := range data.VOLUME_GROUPS {
			keys = append(keys, k)
		}

		tx := d.db.MustBegin()
		tx.MustExec(fmt.Sprintf(data.TruncateRef, "descid"))
		tx.MustExec(fmt.Sprintf(data.InsertRef, "descid"), pq.Array(keys), pq.StringArray(data.VOLUME_GROUPS))
		err = tx.Commit()
		ErrorTx(err, tx)
	}
}

func (d *DBHandler) InitDfnameid() {
	var exist_count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM dfnameid").Scan(&exist_count)
	ErrorFatal(err)

	if exist_count < len(data.DFNAME_DATA) {
		keys := make([]int, 0)
		for k := range data.DFNAME_DATA {
			keys = append(keys, k)
		}

		tx := d.db.MustBegin()
		tx.MustExec(fmt.Sprintf(data.TruncateRef, "dfnameid"))
		tx.MustExec(fmt.Sprintf(data.InsertRef, "dfnameid"), pq.Array(keys), pq.StringArray(data.VOLUME_GROUPS))
		err = tx.Commit()
		ErrorTx(err, tx)
	}
}

func (d *DBHandler) InitLvnameid() {
	var exist_count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM lvnameid").Scan(&exist_count)
	ErrorFatal(err)

	if exist_count < len(data.LVNAME_DATA) {
		keys := make([]int, 0)
		for k := range data.LVNAME_DATA {
			keys = append(keys, k)
		}

		tx := d.db.MustBegin()
		tx.MustExec(fmt.Sprintf(data.TruncateRef, "lvnameid"))
		tx.MustExec(fmt.Sprintf(data.InsertRef, "lvnameid"), pq.Array(keys), pq.StringArray(data.VOLUME_GROUPS))
		err = tx.Commit()
		ErrorTx(err, tx)
	}
}

func (d *DBHandler) InitRef() {
	for _, v := range metric_ref_tables {
		tablename := d.getTableName("day", v, false)
		var exist_count int
		err := d.db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tablename)).Scan(&exist_count)
		ErrorFatal(err)

		var total_count int
		var total_data []string
		switch v {
		case "proccmd":
			total_count = 1000
			total_data = data.PROCCMD_DATA
		case "procuserid":
			total_count = 100
			total_data = data.PROCUSERID_DATA
		case "procargid":
			total_count = 1000
			total_data = data.PROCARGID_DATA
		}

		if exist_count < total_count {
			keys := make([]int, 0)
			for k := range total_data {
				keys = append(keys, k)
			}

			tx := d.db.MustBegin()
			tx.MustExec(fmt.Sprintf(data.TruncateRef, tablename))
			tx.MustExec(fmt.Sprintf(data.InsertRef, tablename), pq.Array(keys), pq.StringArray(total_data))
			err = tx.Commit()
			ErrorTx(err, tx)
		}
	}
}

func DBConnection(dbinfo ConfigDbInfo) *sqlx.DB {
	conn := dbinfo.Datasource()
	db, err := sqlx.Connect("postgres", conn)
	ErrorFatal(err)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(5)
	return db
}

func DBInit(dbinfo ConfigDbInfo, DemoHandler DemoHandler) *DBHandler {
	d := &DBHandler{
		db:     DBConnection(dbinfo),
		dbtype: dbinfo.Name[:2],
		name:   dbinfo.Name,
		demo:   DemoHandler,
	}
	d.CheckTable()
	d.InitExecAgentHostInfo()
	d.InitDeviceid()
	d.InitDescid()
	d.InitDfnameid()
	d.InitLvnameid()
	d.InitRef()

	if d.dbtype == "pg" {
		go d.CheckHourTableMetric()
		go d.CheckDailyTableRef()
	}
	go d.UpdateOntuneinfo()

	return d
}
