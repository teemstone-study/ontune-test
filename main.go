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
	go demo_handler.GeneratePerf(hostintervals.Perf)
	go demo_handler.GenerateCpu(hostintervals.CPU)
	go demo_handler.GenerateDisk(hostintervals.Disk)
	go demo_handler.GenerateNet(hostintervals.Net)
	go demo_handler.GenerateProc(hostintervals.Proc)
	go demo_handler.GenerateDf(hostintervals.DF)

	for {
		select {
		case perf_data := <-app.GlobalChannel.DemoPerfData:
			process_handler.ReceivePerf(perf_data)
		case avg_perf_data := <-app.GlobalChannel.DemoAvgPerfData:
			process_handler.ReceiveAvgPerf(avg_perf_data, "avgperf")
		case avgmax_perf_data := <-app.GlobalChannel.DemoAvgMaxPerfData:
			process_handler.ReceiveAvgPerf(avgmax_perf_data, "avgmaxperf")
		case cpu_data := <-app.GlobalChannel.DemoCpuData:
			process_handler.ReceiveCpu(cpu_data)
		case avg_cpu_data := <-app.GlobalChannel.DemoAvgCpuData:
			process_handler.ReceiveAvgCpu(avg_cpu_data, "avgcpu")
		case avgmax_cpu_data := <-app.GlobalChannel.DemoAvgMaxCpuData:
			process_handler.ReceiveAvgCpu(avgmax_cpu_data, "avgmaxcpu")
		case disk_data := <-app.GlobalChannel.DemoDiskData:
			process_handler.ReceiveDisk(disk_data)
		case avg_disk_data := <-app.GlobalChannel.DemoAvgDiskData:
			process_handler.ReceiveAvgDisk(avg_disk_data, "avgdisk")
		case avgmax_disk_data := <-app.GlobalChannel.DemoAvgMaxDiskData:
			process_handler.ReceiveAvgDisk(avgmax_disk_data, "avgmaxdisk")
		case net_data := <-app.GlobalChannel.DemoNetData:
			process_handler.ReceiveNet(net_data)
		case avg_net_data := <-app.GlobalChannel.DemoAvgNetData:
			process_handler.ReceiveAvgNet(avg_net_data, "avgnet")
		case avgmax_net_data := <-app.GlobalChannel.DemoAvgMaxNetData:
			process_handler.ReceiveAvgNet(avgmax_net_data, "avgmaxnet")
		case proc_data := <-app.GlobalChannel.DemoProcData:
			process_handler.ReceiveProc(proc_data)
		case avg_proc_data := <-app.GlobalChannel.DemoAvgProcData:
			process_handler.ReceiveAvgProc(avg_proc_data, "avgpid")
		case avgmax_proc_data := <-app.GlobalChannel.DemoAvgMaxProcData:
			process_handler.ReceiveAvgProc(avgmax_proc_data, "avgmaxpid")
		case avg_df_data := <-app.GlobalChannel.DemoAvgDfData:
			process_handler.ReceiveAvgDf(avg_df_data)
		case request_avg_flag := <-app.GlobalChannel.AverageRequest:
			switch request_avg_flag {
			case "perf":
				go demo_handler.GenerateAvgPerf()
			case "proc":
				go demo_handler.GenerateAvgProc()
			case "disk":
				go demo_handler.GenerateAvgDisk()
			case "net":
				go demo_handler.GenerateAvgNet()
			case "cpu":
				go demo_handler.GenerateAvgCpu()
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
