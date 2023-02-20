package app

import (
	"fmt"
	"mgr/data"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
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
	metric_tables = []string{"realtimeperf", "avgperf"}
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
				if tb[len(tb)-4:] == "perf" {
					var ontunetime_type = "timestamptz"
					tx.MustExec(fmt.Sprintf(data.MetricPerfStmt, tb, ontunetime_type))
					tx.MustExec(fmt.Sprintf(data.MetricPerfHypertable, tb))
				}
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
		if value == nil || contains(key) || getMapSize(val) == 0 {
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

			if getMapSize(val) > 0 {
				if d.dbtype == "pg" {
					tx.MustExec(fmt.Sprintf(data.InsertRealtimeperfPg, d.getMetricTableName("realtimeperf", false)), args...)
				} else if d.dbtype == "ts" {
					tx.MustExec(fmt.Sprintf(data.InsertRealtimeperfTs, d.getMetricTableName("realtimeperf", false)), args...)
				}
				err := tx.Commit()
				ErrorTx(err, tx)
				LogWrite("log", fmt.Sprintf("realtime %d %s %s %d %d", idx, k, d.dbtype, len(realtimeperf_arr.Ontunetime), GlobalOntunetime))
			}

		} else if len(k) >= 12 && k[:12] == "realtimeproc" {
			fmt.Println("proc")
		} else if len(k) >= 10 && k[:10] == "realtimeio" {
			fmt.Println("io")
		} else if len(k) >= 11 && k[:11] == "realtimecpu" {
			fmt.Println("cpu")
		} else if len(k) >= 7 && k[:7] == "avgperf" {
			var avgperf_arr data.BasicperfArr = data.BasicperfArr{}
			insert_agent_data.Range(func(k, v any) bool {
				agent_struct_data := v.(*data.Basicperf)
				avgperf_arr.SetData(*agent_struct_data)

				return true
			})

			tx := d.db.MustBegin()

			if d.dbtype == "pg" {
				tx.MustExec(fmt.Sprintf(data.InsertRealtimeperfPg, d.getMetricTableName("avgperf", false)), avgperf_arr.GetArgs(d.dbtype)...)
			} else if d.dbtype == "ts" {
				tx.MustExec(fmt.Sprintf(data.InsertRealtimeperfTs, d.getMetricTableName("avgperf", false)), avgperf_arr.GetArgs(d.dbtype)...)
			}
			err := tx.Commit()
			LogWrite("log", fmt.Sprintf("avg %d %s %s %d", idx, k, d.dbtype, GlobalOntunetime))
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
	if lrtp_data, ok := agent_data.Load("lastrealtimeperf"); ok && getMapSize(lrtp_data.(*sync.Map)) > 0 {
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

	go d.CheckHourTableMetric()
	go d.UpdateOntuneinfo()

	return d
}
