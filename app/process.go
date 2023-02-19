package app

import (
	"math"
	"mgr/data"
	"strconv"
	"time"
)

type ProcessHandler struct {
	AgentTableNames []map[string]string
	AgentData       map[string]map[int]interface{}
	Ontunetime      int64
	LPtime          int64
	LPCount         int
	AvgBasictime    int64
	AvgProctime     int64
	AvgIotime       int64
	AvgCputime      int64
}

func (a *ProcessHandler) Init() {
	a.Ontunetime = time.Now().Unix()
	a.LPtime = time.Now().Unix()
	a.AvgBasictime = time.Now().Unix()
	a.AvgProctime = time.Now().Unix()
	a.AvgIotime = time.Now().Unix()
	a.AvgCputime = time.Now().Unix()

	a.AgentData = make(map[string]map[int]interface{})
	a.InitTableNames()

	InitMap(&a.AgentData)
}

func InitMap(m *map[string]map[int]interface{}) {
	(*m)["lastrealtimeperf"] = make(map[int]interface{})
	(*m)["lastperf"] = make(map[int]interface{})

	(*m)["realtimeperf1"] = make(map[int]interface{})
	(*m)["realtimeperf2"] = make(map[int]interface{})
	(*m)["realtimeperf3"] = make(map[int]interface{})
	(*m)["realtimeperf4"] = make(map[int]interface{})
	(*m)["realtimeperf5"] = make(map[int]interface{})
	(*m)["realtimeperf6"] = make(map[int]interface{})
	(*m)["realtimeperf7"] = make(map[int]interface{})
	(*m)["realtimeperf8"] = make(map[int]interface{})
	(*m)["realtimeperf9"] = make(map[int]interface{})
	(*m)["realtimeperf0"] = make(map[int]interface{})

	(*m)["avgperf1"] = make(map[int]interface{})
	(*m)["avgperf2"] = make(map[int]interface{})
	(*m)["avgperf3"] = make(map[int]interface{})
	(*m)["avgperf4"] = make(map[int]interface{})
	(*m)["avgperf5"] = make(map[int]interface{})
	(*m)["avgperf6"] = make(map[int]interface{})
	(*m)["avgperf7"] = make(map[int]interface{})
	(*m)["avgperf8"] = make(map[int]interface{})
	(*m)["avgperf9"] = make(map[int]interface{})
	(*m)["avgperf0"] = make(map[int]interface{})
}

func (a *ProcessHandler) InitTableNames() {
	a.AgentTableNames = make([]map[string]string, 10)

	for i := 0; i < 10; i++ {
		a.AgentTableNames[i] = make(map[string]string)
		a.AgentTableNames[i]["realtimeperf"] = "realtimeperf" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgperf"] = "avgperf" + strconv.Itoa(i)
	}
}

func (a *ProcessHandler) ReceiveBasicPerf(perf_data []data.Basicperf) {
	GlobalMutex.ProcDataM.Lock()
	for _, p := range perf_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimeperf"]
		if _, ok := a.AgentData[tablename][p.Agentid]; !ok {
			a.AgentData[tablename][p.Agentid] = data.Basicperf{}
		}
		a.AgentData[tablename][p.Agentid] = p

		a.SetLastrealtimeperf("basic", p)
		a.SetLastperf(p)
	}
	GlobalMutex.ProcDataM.Unlock()
	a.LPCount = a.LPCount + 1
}

func (a *ProcessHandler) ReceiveAvgBasicPerf(perf_data []data.Basicperf) {
	GlobalMutex.ProcDataM.Lock()
	for _, p := range perf_data {
		tablename := a.AgentTableNames[p.Agentid%10]["avgperf"]
		a.AgentData[tablename][p.Agentid] = p
	}
	GlobalMutex.ProcDataM.Unlock()
}

func (a *ProcessHandler) SetLastrealtimeperf(item_type string, agent_data interface{}) {
	tablename := "lastrealtimeperf"

	switch item_type {
	case "basic":
		src := agent_data.(data.Basicperf)
		if _, ok := a.AgentData[tablename][src.Agentid]; !ok {
			a.AgentData[tablename][src.Agentid] = data.Lastrealtimeperf{}
		}
		tgt := a.AgentData[tablename][src.Agentid].(data.Lastrealtimeperf)

		// Overwrite
		tgt.Ontunetime = src.Ontunetime
		tgt.Agentid = src.Agentid
		tgt.Hostname = MapHostifo[src.Agentid]
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
		tgt.Topcpu = src.Topcpu
		tgt.Topbusy = src.Topbusy
		tgt.Diskiops = src.Diskiops
		tgt.Networkiops = src.Networkiops

		a.AgentData[tablename][src.Agentid] = tgt
	}
}

func (a *ProcessHandler) SetLastperf(agent_data interface{}) {
	tablename := "lastperf"

	src := agent_data.(data.Basicperf)
	if _, ok := a.AgentData[tablename][src.Agentid]; !ok {
		a.AgentData[tablename][src.Agentid] = data.Lastperf{}
	}
	tgt := a.AgentData[tablename][src.Agentid].(data.Lastperf)

	// Add and Replace
	tgt.Ontunetime = int64(math.Max(float64(tgt.Ontunetime), float64(src.Ontunetime)))
	tgt.Hostname = MapHostifo[src.Agentid]
	tgt.User = tgt.User + src.User
	tgt.Sys = tgt.Sys + src.Sys
	tgt.Wait = tgt.Wait + src.Wait
	tgt.Idle = tgt.Idle + src.Idle
	tgt.Avm = tgt.Avm + src.Avm
	tgt.Memoryused = tgt.Memoryused + src.Memoryused
	tgt.Filecache = tgt.Filecache + src.Memorycache
	tgt.Swapused = tgt.Swapused + src.Swapused
	tgt.Topcpu = tgt.Topcpu + src.Topcpu
	tgt.Topbusy = tgt.Topbusy + src.Topbusy
	tgt.Runqueue = tgt.Runqueue + src.Runqueue
	tgt.Blockqueue = tgt.Blockqueue + src.Blockqueue
	tgt.Pagingspacein = tgt.Pagingspacein + src.Pagingspacein
	tgt.Pagingspaceout = tgt.Pagingspaceout + src.Pagingspaceout
	tgt.Filesystemin = tgt.Filesystemin + src.Filesystemin
	tgt.Filesystemout = tgt.Filesystemout + src.Filesystemout
	tgt.Memoryscan = tgt.Memoryscan + src.Memoryscan
	tgt.Memoryfreed = tgt.Memoryfreed + src.Memoryfreed
	tgt.Swapactive = tgt.Swapactive + src.Swapactive
	tgt.Fork = tgt.Fork + src.Fork
	tgt.Exec = tgt.Exec + src.Exec
	tgt.Interupt = tgt.Interupt + src.Interupt
	tgt.Systemcall = tgt.Systemcall + src.Systemcall
	tgt.Contextswitch = tgt.Contextswitch + src.Contextswitch
	tgt.Semaphore = tgt.Semaphore + src.Semaphore
	tgt.Msg = tgt.Msg + src.Msg
	tgt.Diskiops = tgt.Diskiops + src.Diskiops
	tgt.Networkiops = tgt.Networkiops + src.Networkiops

	a.AgentData[tablename][src.Agentid] = tgt
}

func (a *ProcessHandler) ProcessData() {
	for {
		newtime := time.Now().Unix()
		if newtime >= a.LPtime+LASTPERF_TIMER {
			a.ProcessLastperf()
			GlobalChannel.LastPerfData <- a.AgentData["lastperf"]
			a.LPtime = newtime

			// Lastperf Init
			<-GlobalChannel.LastperfCopyDone
			a.LPCount = 0
			GlobalMutex.ProcDataM.Lock()
			a.AgentData["lastperf"] = make(map[int]interface{})
			GlobalMutex.ProcDataM.Unlock()

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
			GlobalMutex.ProcDataM.Lock()
			for i := 0; i < 10; i++ {
				realtimetablename := a.AgentTableNames[i]["realtimeperf"]
				a.AgentData[realtimetablename] = make(map[int]interface{})

				avgtablename := a.AgentTableNames[i]["avgperf"]
				a.AgentData[avgtablename] = make(map[int]interface{})
			}
			GlobalMutex.ProcDataM.Unlock()
			<-GlobalChannel.AgentInsertDone
		}
		time.Sleep(time.Millisecond * time.Duration(1))
	}
}

func (a *ProcessHandler) ProcessLastperf() {
	if a.LPCount == 0 {
		GlobalMutex.ProcDataM.Lock()
		a.AgentData["lastperf"] = make(map[int]interface{})
		GlobalMutex.ProcDataM.Unlock()
		return
	}

	for k, v := range a.AgentData["lastperf"] {
		lpdata := v.(data.Lastperf)
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

		GlobalMutex.ProcDataM.Lock()
		a.AgentData["lastperf"][k] = lpdata
		GlobalMutex.ProcDataM.Unlock()
	}
}

func (a *ProcessHandler) RequestAvgData(intervals ConfigHost) {
	for {
		newtime := time.Now().Unix()
		if newtime >= a.AvgBasictime+int64(intervals.Perf.Avg) {
			GlobalChannel.AverageRequest <- "basic"
			a.AvgBasictime = newtime
		}
		time.Sleep(time.Millisecond * time.Duration(1))
	}
}
