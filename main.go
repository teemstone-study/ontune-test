package main

import (
	"fmt"
	"mgr/app"
	"sync"
	"time"
)

// Config: Yaml Configuration Processing
// Demo: Demo Agent Thread
// Processing: Main Thread
// DB: Database Thread

func main() {
	app.LogWrite("log", "---")
	config := app.GetConfig("config.yml")
	hostcount := config.Demo.HostCount

	var hostintervals app.ConfigHost = config.Host

	demo_handler := app.DemoHandler{}
	demo_handler.Init(hostcount)
	demo_handler.InitDemoAgentInfo()

	process_handler := app.ProcessHandler{}
	process_handler.Init()

	db_handler := make([]app.DBHandler, 0)
	for _, dbinfo := range config.Database {
		dh := *app.DBInit(dbinfo, demo_handler)
		db_handler = append(db_handler, dh)
	}

	go process_handler.ProcessData()
	go process_handler.RequestAvgData(hostintervals)
	go demo_handler.GenerateBasicPerf(hostintervals.Perf)

	for {
		select {
		case basic_data := <-app.GlobalChannel.DemoBasicData:
			process_handler.ReceiveBasicPerf(basic_data)
		case avg_basic_data := <-app.GlobalChannel.DemoAvgBasicData:
			process_handler.ReceiveAvgBasicPerf(avg_basic_data)
		case request_avg_flag := <-app.GlobalChannel.AverageRequest:
			switch request_avg_flag {
			case "basic":
				go demo_handler.GenerateAvgBasicPerf()
			case "proc":
				fmt.Println("proc")
			}
		case agent_data := <-app.GlobalChannel.AgentData:
			// Deep Copy
			receive_data := app.CopyAgentMap(agent_data)
			app.GlobalChannel.AgentCopyDone <- true

			for _, dh := range db_handler {
				var wg sync.WaitGroup
				wg.Add(app.MAX_THREAD + 1)

				// Deep Copy
				insert_data := app.CopyAgentMap(receive_data)

				go dh.InsertLastRealtimePerfData(insert_data, &wg)

				for i := 0; i < app.MAX_THREAD; i++ {
					go dh.InsertAgentData(insert_data, &wg, i)
				}
				wg.Wait()
				app.AgentDataProcess = &sync.Map{}
			}

			app.GlobalChannel.AgentInsertDone <- true
			app.LogWrite("log", fmt.Sprintf("agent insert completed %v", time.Now()))
		case lastperf_data := <-app.GlobalChannel.LastPerfData:
			receive_data := app.CopyMap(lastperf_data)
			app.GlobalChannel.LastperfCopyDone <- true

			for _, dh := range db_handler {
				insert_data := app.CopyMap(receive_data)

				var wg sync.WaitGroup
				wg.Add(1)
				go dh.InsertLastPerfData(insert_data, &wg)
				wg.Wait()
			}
			app.GlobalChannel.LastperfInsertDone <- true
			app.LogWrite("log", fmt.Sprintf("lastperf insert completed %v", time.Now()))

		}
		time.Sleep(time.Millisecond * time.Duration(1))
	}
}
