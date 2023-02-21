package app

import (
	"fmt"
	"math"
	"mgr/data"
	"strconv"
	"sync"
	"time"
)

type ProcessHandler struct {
	AgentTableNames []map[string]string
	AgentData       *sync.Map
	Ontunetime      int64
	LPtime          int64
	LPCount         int
	AvgBasictime    int64
	AvgProctime     int64
	AvgDisktime     int64
	AvgNettime      int64
	AvgCputime      int64
}

func (a *ProcessHandler) Init() {
	a.Ontunetime = time.Now().Unix()
	a.LPtime = time.Now().Unix()
	a.AvgBasictime = time.Now().Unix()
	a.AvgProctime = time.Now().Unix()
	a.AvgDisktime = time.Now().Unix()
	a.AvgNettime = time.Now().Unix()
	a.AvgCputime = time.Now().Unix()

	a.AgentData = &sync.Map{}
	a.InitTableNames()
}

func (a *ProcessHandler) InitTableNames() {
	a.AgentTableNames = make([]map[string]string, 10)

	for i := 0; i < 10; i++ {
		a.AgentTableNames[i] = make(map[string]string)
		a.AgentTableNames[i]["realtimeperf"] = "realtimeperf" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgperf"] = "avgperf" + strconv.Itoa(i)
		a.AgentTableNames[i]["realtimecpu"] = "realtimecpu" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgcpu"] = "avgcpu" + strconv.Itoa(i)
		a.AgentTableNames[i]["realtimedisk"] = "realtimedisk" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgdisk"] = "avgdisk" + strconv.Itoa(i)
		a.AgentTableNames[i]["realtimenet"] = "realtimenet" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgnet"] = "avgnet" + strconv.Itoa(i)
	}
}

func (a *ProcessHandler) ReceiveBasicPerf(perf_data []*data.Basicperf) {
	for _, p := range perf_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimeperf"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			val.(*sync.Map).Store(p.Agentid, p)
		}

		a.SetLastrealtimeperf("basic", p)
		a.SetLastperf(p, p.Agentid, p.Ontunetime)
	}
	a.LPCount = a.LPCount + 1
}

func (a *ProcessHandler) ReceiveAvgBasicPerf(perf_data []*data.Basicperf) {
	for _, p := range perf_data {
		tablename := a.AgentTableNames[p.Agentid%10]["avgperf"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			val.(*sync.Map).Store(p.Agentid, p)
		}
	}
}

func (a *ProcessHandler) ReceiveCpuPerf(cpu_data []*data.Cpuperf) {
	for _, p := range cpu_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimecpu"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Index)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveAvgCpuPerf(cpu_data []*data.Cpuperf) {
	for _, p := range cpu_data {
		tablename := a.AgentTableNames[p.Agentid%10]["avgcpu"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Index)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveDiskPerf(disk_data []*data.Diskperf) {
	for _, p := range disk_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimedisk"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Ionameid)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveAvgDiskPerf(disk_data []*data.Diskperf) {
	for _, p := range disk_data {
		tablename := a.AgentTableNames[p.Agentid%10]["avgdisk"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Ionameid)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveNetPerf(net_data []*data.Netperf) {
	for _, p := range net_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimenet"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Ionameid)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveAvgNetPerf(net_data []*data.Netperf) {
	for _, p := range net_data {
		tablename := a.AgentTableNames[p.Agentid%10]["avgnet"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Ionameid)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) SetLastrealtimeperf(item_type string, agent_data interface{}) {
	tablename := "lastrealtimeperf"

	switch item_type {
	case "basic":
		src := agent_data.(*data.Basicperf)
		if aval, aok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); aok {
			if val, ok := aval.(*sync.Map).LoadOrStore(src.Agentid, &data.Lastrealtimeperf{}); ok {
				tgt := val.(*data.Lastrealtimeperf)

				// Overwrite
				tgt.Ontunetime = src.Ontunetime
				tgt.Agentid = src.Agentid
				tgt.Hostname = MapHostInfo[src.Agentid]
				tgt.User = src.User
				tgt.Sys = src.Sys
				tgt.Wait = src.Wait
				tgt.Idle = src.Idle
				tgt.Memoryused = src.Memoryused
				tgt.Filecache = src.Memorycache
				tgt.Memorysize = src.Memorysize
				tgt.Avm = src.Avm
				tgt.Swapused = src.Swapused
				tgt.Swapsize = src.Swapsize
				tgt.Diskiorate = src.Diskreadwrite
				tgt.Networkiorate = src.Networkreadwrite
				tgt.Topcpu = src.Topcpu
				tgt.Topbusy = src.Topbusy
				tgt.Diskiops = src.Diskiops
				tgt.Networkiops = src.Networkiops
			}
		}
	}
}

func (a *ProcessHandler) SetLastperf(agent_data interface{}, agent_id int, ontunetime int64) {
	tablename := "lastperf"

	if aval, aok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); aok {
		if val, ok := aval.(*sync.Map).LoadOrStore(agent_id, &data.Lastperf{}); ok {
			tgt := val.(*data.Lastperf)
			tgt.Ontunetime = int64(math.Max(float64(tgt.Ontunetime), float64(ontunetime)))
			tgt.Hostname = MapHostInfo[agent_id]

			// Add and Replace
			if src_basic, ok := agent_data.(*data.Basicperf); ok {
				tgt.User = tgt.User + src_basic.User
				tgt.Sys = tgt.Sys + src_basic.Sys
				tgt.Wait = tgt.Wait + src_basic.Wait
				tgt.Idle = tgt.Idle + src_basic.Idle
				tgt.Avm = tgt.Avm + src_basic.Avm
				tgt.Memoryused = tgt.Memoryused + src_basic.Memoryused
				tgt.Filecache = tgt.Filecache + src_basic.Memorycache
				tgt.Swapused = tgt.Swapused + src_basic.Swapused
				tgt.Diskiorate = tgt.Diskiorate + src_basic.Diskreadwrite
				tgt.Networkiorate = tgt.Networkiorate + src_basic.Networkreadwrite
				tgt.Topcpu = tgt.Topcpu + src_basic.Topcpu
				tgt.Topbusy = tgt.Topbusy + src_basic.Topbusy
				tgt.Runqueue = tgt.Runqueue + src_basic.Runqueue
				tgt.Blockqueue = tgt.Blockqueue + src_basic.Blockqueue
				tgt.Pagingspacein = tgt.Pagingspacein + src_basic.Pagingspacein
				tgt.Pagingspaceout = tgt.Pagingspaceout + src_basic.Pagingspaceout
				tgt.Filesystemin = tgt.Filesystemin + src_basic.Filesystemin
				tgt.Filesystemout = tgt.Filesystemout + src_basic.Filesystemout
				tgt.Memoryscan = tgt.Memoryscan + src_basic.Memoryscan
				tgt.Memoryfreed = tgt.Memoryfreed + src_basic.Memoryfreed
				tgt.Swapactive = tgt.Swapactive + src_basic.Swapactive
				tgt.Fork = tgt.Fork + src_basic.Fork
				tgt.Exec = tgt.Exec + src_basic.Exec
				tgt.Interupt = tgt.Interupt + src_basic.Interupt
				tgt.Systemcall = tgt.Systemcall + src_basic.Systemcall
				tgt.Contextswitch = tgt.Contextswitch + src_basic.Contextswitch
				tgt.Semaphore = tgt.Semaphore + src_basic.Semaphore
				tgt.Msg = tgt.Msg + src_basic.Msg
				tgt.Diskiops = tgt.Diskiops + src_basic.Diskiops
				tgt.Networkiops = tgt.Networkiops + src_basic.Networkiops
			}

		}
	}
}

func (a *ProcessHandler) ProcessData() {
	for {
		newtime := time.Now().Unix()
		if newtime >= a.LPtime+LASTPERF_TIMER {
			a.ProcessLastperf()
			if val, ok := a.AgentData.LoadOrStore("lastperf", &sync.Map{}); ok {
				GlobalChannel.LastPerfData <- val.(*sync.Map)
				a.LPtime = newtime
			}

			// Lastperf Init
			<-GlobalChannel.LastperfCopyDone
			a.LPCount = 0
			a.AgentData.Store("lastperf", &sync.Map{})

			// DBInsert Process Done
			<-GlobalChannel.LastperfInsertDone
		}
		time.Sleep(time.Millisecond * time.Duration(1))

		if newtime >= a.Ontunetime+PROCESS_TIMER {
			GlobalChannel.AgentData <- a.AgentData
			a.Ontunetime = newtime
			time.Sleep(time.Millisecond * time.Duration(1))

			// Realtime, AvgPerf Init
			<-GlobalChannel.AgentCopyDone
			for i := 0; i < 10; i++ {
				realtimetablename := a.AgentTableNames[i]["realtimeperf"]
				a.AgentData.Store(realtimetablename, &sync.Map{})

				avgtablename := a.AgentTableNames[i]["avgperf"]
				a.AgentData.Store(avgtablename, &sync.Map{})

				realtimetablename = a.AgentTableNames[i]["realtimecpu"]
				a.AgentData.Store(realtimetablename, &sync.Map{})

				avgtablename = a.AgentTableNames[i]["avgcpu"]
				a.AgentData.Store(avgtablename, &sync.Map{})

				realtimetablename = a.AgentTableNames[i]["realtimedisk"]
				a.AgentData.Store(realtimetablename, &sync.Map{})

				avgtablename = a.AgentTableNames[i]["avgdisk"]
				a.AgentData.Store(avgtablename, &sync.Map{})

				realtimetablename = a.AgentTableNames[i]["realtimenet"]
				a.AgentData.Store(realtimetablename, &sync.Map{})

				avgtablename = a.AgentTableNames[i]["avgnet"]
				a.AgentData.Store(avgtablename, &sync.Map{})

			}
			<-GlobalChannel.AgentInsertDone
		}
		time.Sleep(time.Millisecond * time.Duration(1))
	}
}

func (a *ProcessHandler) ProcessLastperf() {
	if agentdata, ok := a.AgentData.LoadOrStore("lastperf", &sync.Map{}); ok {
		agentdata.(*sync.Map).Range(func(k, v any) bool {
			lpdata := v.(*data.Lastperf)
			lpdata.User = lpdata.User / a.LPCount
			lpdata.Sys = lpdata.Sys / a.LPCount
			lpdata.Wait = lpdata.Wait / a.LPCount
			lpdata.Idle = lpdata.Idle / a.LPCount
			lpdata.Avm = lpdata.Avm / a.LPCount
			lpdata.Memoryused = lpdata.Memoryused / a.LPCount
			lpdata.Filecache = lpdata.Filecache / a.LPCount
			lpdata.Swapused = lpdata.Swapused / a.LPCount
			lpdata.Diskiorate = lpdata.Diskiorate / a.LPCount
			lpdata.Networkiorate = lpdata.Networkiorate / a.LPCount
			lpdata.Topcpu = lpdata.Topcpu / a.LPCount
			lpdata.Topbusy = lpdata.Topbusy / a.LPCount
			lpdata.Runqueue = lpdata.Runqueue / a.LPCount
			lpdata.Blockqueue = lpdata.Blockqueue / a.LPCount
			lpdata.Pagingspacein = lpdata.Pagingspacein / a.LPCount
			lpdata.Pagingspaceout = lpdata.Pagingspaceout / a.LPCount
			lpdata.Filesystemin = lpdata.Filesystemin / a.LPCount
			lpdata.Filesystemout = lpdata.Filesystemout / a.LPCount
			lpdata.Memoryscan = lpdata.Memoryscan / a.LPCount
			lpdata.Memoryfreed = lpdata.Memoryfreed / a.LPCount
			lpdata.Swapactive = lpdata.Swapactive / a.LPCount
			lpdata.Fork = lpdata.Fork / a.LPCount
			lpdata.Exec = lpdata.Exec / a.LPCount
			lpdata.Interupt = lpdata.Interupt / a.LPCount
			lpdata.Systemcall = lpdata.Systemcall / a.LPCount
			lpdata.Contextswitch = lpdata.Contextswitch / a.LPCount
			lpdata.Semaphore = lpdata.Semaphore / a.LPCount
			lpdata.Msg = lpdata.Msg / a.LPCount
			lpdata.Diskiops = lpdata.Diskiops / a.LPCount
			lpdata.Networkiops = lpdata.Networkiops / a.LPCount
			lpdata.Networkreadrate = lpdata.Networkreadrate / a.LPCount
			lpdata.Networkwriterate = lpdata.Networkwriterate / a.LPCount

			return true
		})
	}
}

func (a *ProcessHandler) RequestAvgData(intervals ConfigHost) {
	for {
		newtime := time.Now().Unix()
		if newtime >= a.AvgBasictime+int64(intervals.Perf.Avg) {
			GlobalChannel.AverageRequest <- "basic"
			a.AvgBasictime = newtime
		}
		if newtime >= a.AvgBasictime+int64(intervals.CPU.Avg) {
			GlobalChannel.AverageRequest <- "cpu"
			a.AvgCputime = newtime
		}
		if newtime >= a.AvgBasictime+int64(intervals.Disk.Avg) {
			GlobalChannel.AverageRequest <- "disk"
			a.AvgDisktime = newtime
		}
		if newtime >= a.AvgBasictime+int64(intervals.Net.Avg) {
			GlobalChannel.AverageRequest <- "net"
			a.AvgNettime = newtime
		}
		time.Sleep(time.Millisecond * time.Duration(1))
	}
}
