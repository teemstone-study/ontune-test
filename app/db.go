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
	metric_single_tables = []string{"lastrealtimeperf", "lastperf", "deviceid", "descid"}
	//metric_ref_tables    = []string{"proccmd", "procuserid", "procargid"}
	metric_tables = []string{"realtimeperf", "avgperf", "realtimecpu", "avgcpu", "realtimedisk", "avgdisk", "realtimenet", "avgnet"}
)

func (d *DBHandler) CheckTable() {
	d.CheckTableInfo()
	d.CheckTableMetric()
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
			case "deviceid":
				tx.MustExec(data.DeviceidStmt)
			case "descid":
				tx.MustExec(data.DescidStmt)
			}
			tx.MustExec(data.InsertTableinfoStmt, tb, time.Now().Unix())
			LogWrite("log", fmt.Sprintf("%s %s table creation is completed\n", d.name, tb))
			err = tx.Commit()
			ErrorTx(err, tx)
		}
	}

	if d.dbtype == "ts" {
		for _, tb := range metric_tables {
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

func (d *DBHandler) CheckTableMetric() {
	d.createMetricTables(false)
}

func (d *DBHandler) CheckHourTableMetric() {
	var mTableCreate bool
	if d.dbtype == "pg" {
		for {
			min := time.Now().Minute()
			if min >= 55 && !mTableCreate {
				mTableCreate = true
				d.createMetricTables(true)
			} else if min != 55 {
				mTableCreate = false
			}
			time.Sleep(time.Second * time.Duration(10))
		}
	}
}

func (d *DBHandler) createMetricTables(check_flag bool) {
	for _, tbname := range metric_tables {
		tablename := d.getMetricTableName(tbname, check_flag)
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
			}
		}
	}
}

func (d *DBHandler) getMetricTableName(tablename string, check_flag bool) string {
	if d.dbtype == "pg" {
		now := time.Now()

		var createHour string
		if check_flag {
			createHour = now.Add(10 * time.Minute).Format(DATE_FORMAT)
		} else {
			createHour = now.Format(DATE_FORMAT)
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
		agent_data.Store(key, &sync.Map{})

		k := key.(string)
		if len(k) >= 12 && k[:12] == "realtimeperf" {
			var realtimeperf_arr data.BasicperfArr = data.BasicperfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Basicperf)
				realtimeperf_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()
			args := realtimeperf_arr.GetArgs(d.dbtype)

			if GetMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertPerfPg, d.getMetricTableName("realtimeperf", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertPerfTs, d.getMetricTableName("realtimeperf", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("log", fmt.Sprintf("rperf %d %s %s %d %d", idx, k, d.dbtype, len(realtimeperf_arr.Ontunetime), GlobalOntunetime))
			}

		} else if len(k) >= 12 && k[:12] == "realtimeproc" {
			fmt.Println("proc")
		} else if len(k) >= 12 && k[:12] == "realtimedisk" {
			var realtimedisk_arr data.DiskperfArr = data.DiskperfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Diskperf)
				realtimedisk_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()
			args := realtimedisk_arr.GetArgs(d.dbtype)

			if GetMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertDiskPg, d.getMetricTableName("realtimedisk", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertDiskTs, d.getMetricTableName("realtimedisk", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("log", fmt.Sprintf("rdisk %d %s %s %d %d", idx, k, d.dbtype, len(realtimedisk_arr.Ontunetime), GlobalOntunetime))
			}
		} else if len(k) >= 11 && k[:11] == "realtimenet" {
			var realtimenet_arr data.NetperfArr = data.NetperfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Netperf)
				realtimenet_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()
			args := realtimenet_arr.GetArgs(d.dbtype)

			if GetMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertNetPg, d.getMetricTableName("realtimenet", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertNetTs, d.getMetricTableName("realtimenet", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("log", fmt.Sprintf("rnet %d %s %s %d %d", idx, k, d.dbtype, len(realtimenet_arr.Ontunetime), GlobalOntunetime))
			}
		} else if len(k) >= 11 && k[:11] == "realtimecpu" {
			var realtimecpu_arr data.CpuperfArr = data.CpuperfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Cpuperf)
				realtimecpu_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()
			args := realtimecpu_arr.GetArgs(d.dbtype)

			if GetMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertCpuPg, d.getMetricTableName("realtimecpu", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertCpuTs, d.getMetricTableName("realtimecpu", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("log", fmt.Sprintf("rcpu %d %s %s %d %d", idx, k, d.dbtype, len(realtimecpu_arr.Ontunetime), GlobalOntunetime))
			}
		} else if len(k) >= 7 && k[:7] == "avgperf" {
			var avgperf_arr data.BasicperfArr = data.BasicperfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Basicperf)
				avgperf_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertPerfPg, d.getMetricTableName("avgperf", false)), avgperf_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertPerfTs, d.getMetricTableName("avgperf", false)), avgperf_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("log", fmt.Sprintf("aperf %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 7 && k[:7] == "avgdisk" {
			var avgdisk_arr data.DiskperfArr = data.DiskperfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Diskperf)
				avgdisk_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertDiskPg, d.getMetricTableName("avgdisk", false)), avgdisk_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertDiskTs, d.getMetricTableName("avgdisk", false)), avgdisk_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("log", fmt.Sprintf("adisk %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 6 && k[:6] == "avgnet" {
			var avgnet_arr data.NetperfArr = data.NetperfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Netperf)
				avgnet_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertNetPg, d.getMetricTableName("avgnet", false)), avgnet_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertNetTs, d.getMetricTableName("avgnet", false)), avgnet_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("log", fmt.Sprintf("anet %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		} else if len(k) >= 6 && k[:6] == "avgcpu" {
			var avgcpu_arr data.CpuperfArr = data.CpuperfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Cpuperf)
				avgcpu_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertCpuPg, d.getMetricTableName("avgcpu", false)), avgcpu_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertCpuTs, d.getMetricTableName("avgcpu", false)), avgcpu_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("log", fmt.Sprintf("acpu %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
			ErrorTx(err, tx)
		}

		time.Sleep(time.Millisecond * time.Duration(1))
		return true
	})

	defer func() {
		if c := recover(); c != nil {
			fmt.Println("recover data")
		}
		//LogWrite("log", fmt.Sprintf("waitgroup %d", idx))
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

	//LogWrite("log", "waitgroup -2")
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
		//LogWrite("log", fmt.Sprintf("lrtp %s %d", d.dbtype, GlobalOntunetime))
		ErrorTx(err, tx)
	}
	//LogWrite("log", "waitgroup -1")
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

	if exist_count < len(DEVICE_IDS) {
		keys := make([]int, 0)
		for k := range DEVICE_IDS {
			keys = append(keys, k)
		}

		tx := d.db.MustBegin()
		tx.MustExec(data.TruncateDeviceid)
		tx.MustExec(data.InsertDeviceid, pq.Array(keys), pq.StringArray(DEVICE_IDS))
		err = tx.Commit()
		ErrorTx(err, tx)
	}
}

func (d *DBHandler) InitDescid() {
	var exist_count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM descid").Scan(&exist_count)
	ErrorFatal(err)

	if exist_count < len(VOLUME_GROUPS) {
		keys := make([]int, 0)
		for k := range VOLUME_GROUPS {
			keys = append(keys, k)
		}

		tx := d.db.MustBegin()
		tx.MustExec(data.TruncateDescid)
		tx.MustExec(data.InsertDescid, pq.Array(keys), pq.StringArray(VOLUME_GROUPS))
		err = tx.Commit()
		ErrorTx(err, tx)
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

	go d.CheckHourTableMetric()
	go d.UpdateOntuneinfo()

	return d
}
