package app

import (
	"fmt"
	"math"
	"math/rand"
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
	AvgPerftime     int64
	AvgProctime     int64
	AvgDisktime     int64
	AvgNettime      int64
	AvgCputime      int64
}

func (a *ProcessHandler) Init() {
	a.Ontunetime = time.Now().Unix()
	a.LPtime = time.Now().Unix()
	a.AvgPerftime = time.Now().Unix()
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
		a.AgentTableNames[i]["avgmaxperf"] = "avgmaxperf" + strconv.Itoa(i)
		a.AgentTableNames[i]["realtimecpu"] = "realtimecpu" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgcpu"] = "avgcpu" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgmaxcpu"] = "avgmaxcpu" + strconv.Itoa(i)
		a.AgentTableNames[i]["realtimedisk"] = "realtimedisk" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgdisk"] = "avgdisk" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgmaxdisk"] = "avgmaxdisk" + strconv.Itoa(i)
		a.AgentTableNames[i]["realtimenet"] = "realtimenet" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgnet"] = "avgnet" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgmaxnet"] = "avgmaxnet" + strconv.Itoa(i)
		a.AgentTableNames[i]["realtimepid"] = "realtimepid" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgpid"] = "avgpid" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgmaxpid"] = "avgmaxpid" + strconv.Itoa(i)
		a.AgentTableNames[i]["realtimeproc"] = "realtimeproc" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgproc"] = "avgproc" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgmaxproc"] = "avgmaxproc" + strconv.Itoa(i)
		a.AgentTableNames[i]["avgdf"] = "avgdf" + strconv.Itoa(i)
	}
}

func (a *ProcessHandler) ReceivePerf(perf_data []*data.Perf) {
	for _, p := range perf_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimeperf"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			val.(*sync.Map).Store(p.Agentid, p)
		}

		a.SetLastrealtimeperf("perf", p)
		a.SetLastperf(p, p.Agentid, p.Ontunetime)
	}
	a.LPCount = a.LPCount + 1
}

func (a *ProcessHandler) ReceiveAvgPerf(perf_data []*data.Perf, tablename string) {
	for _, p := range perf_data {
		tablename := a.AgentTableNames[p.Agentid%10][tablename]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			val.(*sync.Map).Store(p.Agentid, p)
		}
	}
}

func (a *ProcessHandler) ReceiveCpu(cpu_data []*data.Cpu) {
	for _, p := range cpu_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimecpu"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Index)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveAvgCpu(cpu_data []*data.Cpu, tablename string) {
	for _, p := range cpu_data {
		tablename := a.AgentTableNames[p.Agentid%10][tablename]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Index)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveDisk(disk_data []*data.Disk) {
	for _, p := range disk_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimedisk"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Ionameid)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveAvgDisk(disk_data []*data.Disk, tablename string) {
	for _, p := range disk_data {
		tablename := a.AgentTableNames[p.Agentid%10][tablename]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Ionameid)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveNet(net_data []*data.Net) {
	for _, p := range net_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimenet"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Ionameid)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveAvgNet(net_data []*data.Net, tablename string) {
	for _, p := range net_data {
		tablename := a.AgentTableNames[p.Agentid%10][tablename]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d", p.Agentid, p.Ionameid)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveProc(pid_data []*data.Pid) {
	for _, p := range pid_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimepid"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d_%d_%d", p.Agentid, p.Cmdid, p.Userid, p.Argid)
			val.(*sync.Map).Store(key, p)
		}
	}

	proc_data := MakeProc(pid_data)

	for _, p := range proc_data {
		tablename := a.AgentTableNames[p.Agentid%10]["realtimeproc"]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d_%d", p.Agentid, p.Cmdid, p.Userid)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) ReceiveAvgProc(pid_data []*data.Pid, tablename string) {
	for _, p := range pid_data {
		tablename := a.AgentTableNames[p.Agentid%10][tablename]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d_%d_%d", p.Agentid, p.Cmdid, p.Userid, p.Argid)
			val.(*sync.Map).Store(key, p)
		}
	}

	var proctablename string
	if tablename == "avgpid" {
		proctablename = "avgproc"
	} else if tablename == "avgmaxpid" {
		proctablename = "avgmaxproc"
	} else {
		return
	}
	proc_data := MakeProc(pid_data)

	for _, p := range proc_data {
		tablename := a.AgentTableNames[p.Agentid%10][proctablename]
		if val, ok := a.AgentData.LoadOrStore(tablename, &sync.Map{}); ok {
			key := fmt.Sprintf("%d_%d_%d", p.Agentid, p.Cmdid, p.Userid)
			val.(*sync.Map).Store(key, p)
		}
	}
}

func (a *ProcessHandler) SetLastrealtimeperf(item_type string, agent_data interface{}) {
	tablename := "lastrealtimeperf"

	switch item_type {
	case "perf":
		src := agent_data.(*data.Perf)
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
			if src_perf, ok := agent_data.(*data.Perf); ok {
				tgt.User = tgt.User + src_perf.User
				tgt.Sys = tgt.Sys + src_perf.Sys
				tgt.Wait = tgt.Wait + src_perf.Wait
				tgt.Idle = tgt.Idle + src_perf.Idle
				tgt.Avm = tgt.Avm + src_perf.Avm
				tgt.Memoryused = tgt.Memoryused + src_perf.Memoryused
				tgt.Filecache = tgt.Filecache + src_perf.Memorycache
				tgt.Swapused = tgt.Swapused + src_perf.Swapused
				tgt.Diskiorate = tgt.Diskiorate + src_perf.Diskreadwrite
				tgt.Networkiorate = tgt.Networkiorate + src_perf.Networkreadwrite
				tgt.Topcpu = tgt.Topcpu + src_perf.Topcpu
				tgt.Topbusy = tgt.Topbusy + src_perf.Topbusy
				tgt.Runqueue = tgt.Runqueue + src_perf.Runqueue
				tgt.Blockqueue = tgt.Blockqueue + src_perf.Blockqueue
				tgt.Pagingspacein = tgt.Pagingspacein + src_perf.Pagingspacein
				tgt.Pagingspaceout = tgt.Pagingspaceout + src_perf.Pagingspaceout
				tgt.Filesystemin = tgt.Filesystemin + src_perf.Filesystemin
				tgt.Filesystemout = tgt.Filesystemout + src_perf.Filesystemout
				tgt.Memoryscan = tgt.Memoryscan + src_perf.Memoryscan
				tgt.Memoryfreed = tgt.Memoryfreed + src_perf.Memoryfreed
				tgt.Swapactive = tgt.Swapactive + src_perf.Swapactive
				tgt.Fork = tgt.Fork + src_perf.Fork
				tgt.Exec = tgt.Exec + src_perf.Exec
				tgt.Interupt = tgt.Interupt + src_perf.Interupt
				tgt.Systemcall = tgt.Systemcall + src_perf.Systemcall
				tgt.Contextswitch = tgt.Contextswitch + src_perf.Contextswitch
				tgt.Semaphore = tgt.Semaphore + src_perf.Semaphore
				tgt.Msg = tgt.Msg + src_perf.Msg
				tgt.Diskiops = tgt.Diskiops + src_perf.Diskiops
				tgt.Networkiops = tgt.Networkiops + src_perf.Networkiops
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

				avgmaxtablename := a.AgentTableNames[i]["avgmaxperf"]
				a.AgentData.Store(avgmaxtablename, &sync.Map{})

				realtimetablename = a.AgentTableNames[i]["realtimecpu"]
				a.AgentData.Store(realtimetablename, &sync.Map{})

				avgtablename = a.AgentTableNames[i]["avgcpu"]
				a.AgentData.Store(avgtablename, &sync.Map{})

				avgmaxtablename = a.AgentTableNames[i]["avgmaxcpu"]
				a.AgentData.Store(avgmaxtablename, &sync.Map{})

				realtimetablename = a.AgentTableNames[i]["realtimedisk"]
				a.AgentData.Store(realtimetablename, &sync.Map{})

				avgtablename = a.AgentTableNames[i]["avgdisk"]
				a.AgentData.Store(avgtablename, &sync.Map{})

				avgmaxtablename = a.AgentTableNames[i]["avgmaxdisk"]
				a.AgentData.Store(avgmaxtablename, &sync.Map{})

				realtimetablename = a.AgentTableNames[i]["realtimenet"]
				a.AgentData.Store(realtimetablename, &sync.Map{})

				avgtablename = a.AgentTableNames[i]["avgnet"]
				a.AgentData.Store(avgtablename, &sync.Map{})

				avgmaxtablename = a.AgentTableNames[i]["avgmaxnet"]
				a.AgentData.Store(avgmaxtablename, &sync.Map{})

				realtimetablename = a.AgentTableNames[i]["realtimepid"]
				a.AgentData.Store(realtimetablename, &sync.Map{})

				avgtablename = a.AgentTableNames[i]["avgpid"]
				a.AgentData.Store(avgtablename, &sync.Map{})

				avgmaxtablename = a.AgentTableNames[i]["avgmaxpid"]
				a.AgentData.Store(avgmaxtablename, &sync.Map{})

				realtimetablename = a.AgentTableNames[i]["realtimeproc"]
				a.AgentData.Store(realtimetablename, &sync.Map{})

				avgtablename = a.AgentTableNames[i]["avgproc"]
				a.AgentData.Store(avgtablename, &sync.Map{})

				avgmaxtablename = a.AgentTableNames[i]["avgmaxproc"]
				a.AgentData.Store(avgmaxtablename, &sync.Map{})
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
		if newtime >= a.AvgPerftime+int64(intervals.Perf.Avg) {
			GlobalChannel.AverageRequest <- "perf"
			a.AvgPerftime = newtime
		}
		if newtime >= a.AvgCputime+int64(intervals.CPU.Avg) {
			GlobalChannel.AverageRequest <- "cpu"
			a.AvgCputime = newtime
		}
		if newtime >= a.AvgDisktime+int64(intervals.Disk.Avg) {
			GlobalChannel.AverageRequest <- "disk"
			a.AvgDisktime = newtime
		}
		if newtime >= a.AvgNettime+int64(intervals.Net.Avg) {
			GlobalChannel.AverageRequest <- "net"
			a.AvgNettime = newtime
		}
		if newtime >= a.AvgProctime+int64(intervals.Proc.Avg) {
			// Process ID Value는 Avg 타임에 한번씩 변경하는 식으로 일단 구성할 것..
			PROCCMD_ARR = []int{rand.Intn(len(data.PROCCMD_DATA)), rand.Intn(len(data.PROCCMD_DATA)), rand.Intn(len(data.PROCCMD_DATA))}
			PROCUSER_ARR = []int{rand.Intn(len(data.PROCCMD_DATA))}
			PROCARG_ARR = []int{rand.Intn(len(data.PROCCMD_DATA)), rand.Intn(len(data.PROCCMD_DATA)), rand.Intn(len(data.PROCCMD_DATA))}

			GlobalChannel.AverageRequest <- "proc"
			a.AvgProctime = newtime
		}
		time.Sleep(time.Millisecond * time.Duration(1))
	}
}
