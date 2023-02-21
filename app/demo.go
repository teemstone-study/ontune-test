package app

import (
	"fmt"
	"math"
	"math/rand"
	"mgr/data"
	"strconv"
	"strings"
	"sync"
	"time"
)

type DemoHandler struct {
	Hostcount int
	Interval  int
	Agentinfo data.AgentinfoArr
	Hostinfo  data.HostinfoArr
	AvgPerf   *sync.Map
	AvgCpu    *sync.Map
	AvgDisk   *sync.Map
	AvgNet    *sync.Map
	AvgProc   *sync.Map
}

func (d *DemoHandler) Init(hostcount int) {
	d.Hostcount = hostcount
	d.AvgPerf = &sync.Map{}
	d.AvgCpu = &sync.Map{}
	d.AvgDisk = &sync.Map{}
	d.AvgNet = &sync.Map{}
	d.AvgProc = &sync.Map{}
}

func (d *DemoHandler) InitDemoAgentInfo() {
	ts := time.Now().Unix()

	var agentinfo_arr data.AgentinfoArr
	var hostinfo_arr data.HostinfoArr

	for i := 0; i < d.Hostcount; i++ {
		demo_agent_data := &data.Agentinfo{
			Agentid:           i + 1,
			Agentname:         fmt.Sprintf("DummyAgent%d", i+1),
			Enabled:           1,
			Connected:         1,
			Updated:           1,
			Shorttermperf:     2,
			Shorttermproc:     5,
			Shorttermio:       5,
			Shorttermcpu:      5,
			Longtermperf:      600,
			Longtermproc:      600,
			Longtermio:        600,
			Longtermcpu:       600,
			Model:             "-",
			Serial:            "-",
			Group:             "-",
			Ipaddress:         "localhost",
			Pscommand:         "-",
			Logevent:          "-",
			Processevent:      "-",
			Timecheck:         1,
			Disconnectedtime:  ts,
			Skipdatatypes:     0,
			Virperf:           1,
			Hypervisor:        0,
			Serviceevent:      "-",
			Installdate:       ts,
			Lastconnectedtime: ts,
		}
		agentinfo_arr.SetData(*demo_agent_data)
		d.Agentinfo = agentinfo_arr

		demo_host_data := &data.Hostinfo{
			Agentid:        i + 1,
			Hostname:       fmt.Sprintf("DummyAgent%d", i+1),
			Hostnameext:    fmt.Sprintf("DummyAgent%d", i+1),
			Os:             "-",
			Fw:             "-",
			Agentversion:   "V1",
			Model:          "-",
			Serial:         "-",
			Processorcount: rand.Intn(4),
			Processorclock: rand.Intn(3200),
			Memorysize:     rand.Intn(32000),
			Swapsize:       0,
			Poolid:         -1,
			Replication:    0,
			Smt:            0,
			Micropar:       0,
			Capped:         0,
			Ec:             -1,
			Virtualcpu:     rand.Intn(4),
			Weight:         0,
			Cpupool:        0,
			Ams:            0,
			Allip:          "localhost",
			Numanodecount:  0,
		}
		hostinfo_arr.SetData(*demo_host_data)
		d.Hostinfo = hostinfo_arr

		MapHostInfo[i+1] = fmt.Sprintf("DummyAgent%d", i+1)
	}
}

func (d *DemoHandler) GeneratePerf(interval ConfigScrape) {
	for {
		ts := time.Now().Unix()

		var agent_perf []*data.Perf = make([]*data.Perf, 0)
		for i := 0; i < d.Hostcount; i++ {
			demo_data := &data.Perf{
				Ontunetime:       ts,
				Agenttime:        ts,
				Agentid:          d.Agentinfo.Agentid[i],
				User:             rand.Intn(100),
				Sys:              rand.Intn(100),
				Wait:             rand.Intn(100),
				Idle:             rand.Intn(100),
				Processorcount:   rand.Intn(100),
				Runqueue:         rand.Intn(100),
				Blockqueue:       rand.Intn(100),
				Waitqueue:        rand.Intn(100),
				Pqueue:           rand.Intn(100),
				Pcrateuser:       rand.Intn(100),
				Pcratesys:        rand.Intn(100),
				Memorysize:       rand.Intn(10000),
				Memoryused:       rand.Intn(10000),
				Memorypinned:     rand.Intn(100),
				Memorysys:        rand.Intn(100),
				Memoryuser:       rand.Intn(100),
				Memorycache:      rand.Intn(100),
				Avm:              rand.Intn(10000),
				Pagingspacein:    rand.Intn(100),
				Pagingspaceout:   rand.Intn(100),
				Filesystemin:     rand.Intn(100),
				Filesystemout:    rand.Intn(100),
				Memoryscan:       rand.Intn(100),
				Memoryfreed:      rand.Intn(100),
				Swapsize:         rand.Intn(10000),
				Swapused:         rand.Intn(10000),
				Swapactive:       rand.Intn(100),
				Fork:             rand.Intn(100),
				Exec:             rand.Intn(100),
				Interupt:         rand.Intn(100),
				Systemcall:       rand.Intn(100),
				Contextswitch:    rand.Intn(100),
				Semaphore:        rand.Intn(100),
				Msg:              rand.Intn(100),
				Diskreadwrite:    rand.Intn(100),
				Diskiops:         rand.Intn(100),
				Networkreadwrite: rand.Intn(100),
				Networkiops:      rand.Intn(100),
				Topcommandid:     rand.Intn(100),
				Topcommandcount:  rand.Intn(100),
				Topuserid:        rand.Intn(100),
				Topcpu:           rand.Intn(100),
				Topdiskid:        rand.Intn(100),
				Topvgid:          rand.Intn(100),
				Topbusy:          rand.Intn(100),
				Maxpid:           rand.Intn(100),
				Threadcount:      rand.Intn(100),
				Pidcount:         rand.Intn(100),
				Linuxbuffer:      rand.Intn(100),
				Linuxcached:      rand.Intn(100),
				Linuxsrec:        rand.Intn(100),
				Memused_mb:       rand.Intn(100),
				Irq:              rand.Intn(100),
				Softirq:          rand.Intn(100),
				Swapused_mb:      rand.Intn(100),
				Dusm:             rand.Intn(100),
			}
			agent_perf = append(agent_perf, demo_data)

			// Store Average Value
			if val, ok := d.AvgPerf.LoadOrStore(demo_data.Agentid, make([]*data.Perf, 0)); ok {
				var avgvalue []*data.Perf = make([]*data.Perf, 0)
				avgvalue = append(avgvalue, val.([]*data.Perf)...)
				d.AvgPerf.Store(demo_data.Agentid, avgvalue)
			}
		}

		GlobalChannel.DemoPerfData <- agent_perf
		time.Sleep(time.Second * time.Duration(interval.Rate))
	}
}

func (d *DemoHandler) GenerateAvgPerf() {
	var agent_perf []*data.Perf = make([]*data.Perf, 0)
	var max_perf []*data.Perf = make([]*data.Perf, 0)

	d.AvgPerf.Range(func(key, value any) bool {
		avg_data := &data.Perf{Agentid: key.(int)}
		max_data := &data.Perf{Agentid: key.(int)}

		for _, av := range value.([]*data.Perf) {
			avg_data.Ontunetime = int64(math.Max(float64(avg_data.Ontunetime), float64(av.Ontunetime)))
			avg_data.Agenttime = int64(math.Max(float64(avg_data.Agenttime), float64(av.Ontunetime)))
			avg_data.User = avg_data.User + av.User
			avg_data.Sys = avg_data.Sys + av.Sys
			avg_data.Wait = avg_data.Wait + av.Wait
			avg_data.Idle = avg_data.Idle + av.Idle
			avg_data.Processorcount = avg_data.Processorcount + av.Processorcount
			avg_data.Runqueue = avg_data.Runqueue + av.Runqueue
			avg_data.Blockqueue = avg_data.Blockqueue + av.Blockqueue
			avg_data.Waitqueue = avg_data.Waitqueue + av.Waitqueue
			avg_data.Pqueue = avg_data.Pqueue + av.Pqueue
			avg_data.Pcrateuser = avg_data.Pcrateuser + av.Pcrateuser
			avg_data.Pcratesys = avg_data.Pcratesys + av.Pcratesys
			avg_data.Memorysize = avg_data.Memorysize + av.Memorysize
			avg_data.Memoryused = avg_data.Memoryused + av.Memoryused
			avg_data.Memorypinned = avg_data.Memorypinned + av.Memorypinned
			avg_data.Memorysys = avg_data.Memorysys + av.Memorysys
			avg_data.Memoryuser = avg_data.Memoryuser + av.Memoryuser
			avg_data.Memorycache = avg_data.Memorycache + av.Memorycache
			avg_data.Avm = avg_data.Avm + av.Avm
			avg_data.Pagingspacein = avg_data.Pagingspacein + av.Pagingspacein
			avg_data.Pagingspaceout = avg_data.Pagingspaceout + av.Pagingspaceout
			avg_data.Filesystemin = avg_data.Filesystemin + av.Filesystemin
			avg_data.Filesystemout = avg_data.Filesystemout + av.Filesystemout
			avg_data.Memoryscan = avg_data.Memoryscan + av.Memoryscan
			avg_data.Memoryfreed = avg_data.Memoryfreed + av.Memoryfreed
			avg_data.Swapsize = avg_data.Swapsize + av.Swapsize
			avg_data.Swapused = avg_data.Swapused + av.Swapused
			avg_data.Swapactive = avg_data.Swapactive + av.Swapactive
			avg_data.Fork = avg_data.Fork + av.Fork
			avg_data.Exec = avg_data.Exec + av.Exec
			avg_data.Interupt = avg_data.Interupt + av.Interupt
			avg_data.Systemcall = avg_data.Systemcall + av.Systemcall
			avg_data.Contextswitch = avg_data.Contextswitch + av.Contextswitch
			avg_data.Semaphore = avg_data.Semaphore + av.Semaphore
			avg_data.Msg = avg_data.Msg + av.Msg
			avg_data.Diskreadwrite = avg_data.Diskreadwrite + av.Diskreadwrite
			avg_data.Diskiops = avg_data.Diskiops + av.Diskiops
			avg_data.Networkreadwrite = avg_data.Networkreadwrite + av.Networkreadwrite
			avg_data.Networkiops = avg_data.Networkiops + av.Networkiops
			avg_data.Topcommandid = avg_data.Topcommandid + av.Topcommandid
			avg_data.Topcommandcount = avg_data.Topcommandcount + av.Topcommandcount
			avg_data.Topuserid = avg_data.Topuserid + av.Topuserid
			avg_data.Topcpu = avg_data.Topcpu + av.Topcpu
			avg_data.Topdiskid = avg_data.Topdiskid + av.Topdiskid
			avg_data.Topvgid = avg_data.Topvgid + av.Topvgid
			avg_data.Topbusy = avg_data.Topbusy + av.Topbusy
			avg_data.Maxpid = avg_data.Maxpid + av.Maxpid
			avg_data.Threadcount = avg_data.Threadcount + av.Threadcount
			avg_data.Pidcount = avg_data.Pidcount + av.Pidcount
			avg_data.Linuxbuffer = avg_data.Linuxbuffer + av.Linuxbuffer
			avg_data.Linuxcached = avg_data.Linuxcached + av.Linuxcached
			avg_data.Linuxsrec = avg_data.Linuxsrec + av.Linuxsrec
			avg_data.Memused_mb = avg_data.Memused_mb + av.Memused_mb
			avg_data.Irq = avg_data.Irq + av.Irq
			avg_data.Softirq = avg_data.Softirq + av.Softirq
			avg_data.Swapused_mb = avg_data.Swapused_mb + av.Swapused_mb
			avg_data.Dusm = avg_data.Dusm + av.Dusm

			max_data.Ontunetime = int64(math.Max(float64(max_data.Ontunetime), float64(av.Ontunetime)))
			max_data.Agenttime = int64(math.Max(float64(max_data.Agenttime), float64(av.Ontunetime)))
			max_data.User = int(math.Max(float64(max_data.User), float64(av.User)))
			max_data.Sys = int(math.Max(float64(max_data.Sys), float64(av.Sys)))
			max_data.Wait = int(math.Max(float64(max_data.Wait), float64(av.Wait)))
			max_data.Idle = int(math.Max(float64(max_data.Idle), float64(av.Idle)))
			max_data.Processorcount = int(math.Max(float64(max_data.Processorcount), float64(av.Processorcount)))
			max_data.Runqueue = int(math.Max(float64(max_data.Runqueue), float64(av.Runqueue)))
			max_data.Blockqueue = int(math.Max(float64(max_data.Blockqueue), float64(av.Blockqueue)))
			max_data.Waitqueue = int(math.Max(float64(max_data.Waitqueue), float64(av.Waitqueue)))
			max_data.Pqueue = int(math.Max(float64(max_data.Pqueue), float64(av.Pqueue)))
			max_data.Pcrateuser = int(math.Max(float64(max_data.Pcrateuser), float64(av.Pcrateuser)))
			max_data.Pcratesys = int(math.Max(float64(max_data.Pcratesys), float64(av.Pcratesys)))
			max_data.Memorysize = int(math.Max(float64(max_data.Memorysize), float64(av.Memorysize)))
			max_data.Memoryused = int(math.Max(float64(max_data.Memoryused), float64(av.Memoryused)))
			max_data.Memorypinned = int(math.Max(float64(max_data.Memorypinned), float64(av.Memorypinned)))
			max_data.Memorysys = int(math.Max(float64(max_data.Memorysys), float64(av.Memorysys)))
			max_data.Memoryuser = int(math.Max(float64(max_data.Memoryuser), float64(av.Memoryuser)))
			max_data.Memorycache = int(math.Max(float64(max_data.Memorycache), float64(av.Memorycache)))
			max_data.Avm = int(math.Max(float64(max_data.Avm), float64(av.Avm)))
			max_data.Pagingspacein = int(math.Max(float64(max_data.Pagingspacein), float64(av.Pagingspacein)))
			max_data.Pagingspaceout = int(math.Max(float64(max_data.Pagingspaceout), float64(av.Pagingspaceout)))
			max_data.Filesystemin = int(math.Max(float64(max_data.Filesystemin), float64(av.Filesystemin)))
			max_data.Filesystemout = int(math.Max(float64(max_data.Filesystemout), float64(av.Filesystemout)))
			max_data.Memoryscan = int(math.Max(float64(max_data.Memoryscan), float64(av.Memoryscan)))
			max_data.Memoryfreed = int(math.Max(float64(max_data.Memoryfreed), float64(av.Memoryfreed)))
			max_data.Swapsize = int(math.Max(float64(max_data.Swapsize), float64(av.Swapsize)))
			max_data.Swapused = int(math.Max(float64(max_data.Swapused), float64(av.Swapused)))
			max_data.Swapactive = int(math.Max(float64(max_data.Swapactive), float64(av.Swapactive)))
			max_data.Fork = int(math.Max(float64(max_data.Fork), float64(av.Fork)))
			max_data.Exec = int(math.Max(float64(max_data.Exec), float64(av.Exec)))
			max_data.Interupt = int(math.Max(float64(max_data.Interupt), float64(av.Interupt)))
			max_data.Systemcall = int(math.Max(float64(max_data.Systemcall), float64(av.Systemcall)))
			max_data.Contextswitch = int(math.Max(float64(max_data.Contextswitch), float64(av.Contextswitch)))
			max_data.Semaphore = int(math.Max(float64(max_data.Semaphore), float64(av.Semaphore)))
			max_data.Msg = int(math.Max(float64(max_data.Msg), float64(av.Msg)))
			max_data.Diskreadwrite = int(math.Max(float64(max_data.Diskreadwrite), float64(av.Diskreadwrite)))
			max_data.Diskiops = int(math.Max(float64(max_data.Diskiops), float64(av.Diskiops)))
			max_data.Networkreadwrite = int(math.Max(float64(max_data.Networkreadwrite), float64(av.Networkreadwrite)))
			max_data.Networkiops = int(math.Max(float64(max_data.Networkiops), float64(av.Networkiops)))
			max_data.Topcommandid = int(math.Max(float64(max_data.Topcommandid), float64(av.Topcommandid)))
			max_data.Topcommandcount = int(math.Max(float64(max_data.Topcommandcount), float64(av.Topcommandcount)))
			max_data.Topuserid = int(math.Max(float64(max_data.Topuserid), float64(av.Topuserid)))
			max_data.Topcpu = int(math.Max(float64(max_data.Topcpu), float64(av.Topcpu)))
			max_data.Topdiskid = int(math.Max(float64(max_data.Topdiskid), float64(av.Topdiskid)))
			max_data.Topvgid = int(math.Max(float64(max_data.Topvgid), float64(av.Topvgid)))
			max_data.Topbusy = int(math.Max(float64(max_data.Topbusy), float64(av.Topbusy)))
			max_data.Maxpid = int(math.Max(float64(max_data.Maxpid), float64(av.Maxpid)))
			max_data.Threadcount = int(math.Max(float64(max_data.Threadcount), float64(av.Threadcount)))
			max_data.Pidcount = int(math.Max(float64(max_data.Pidcount), float64(av.Pidcount)))
			max_data.Linuxbuffer = int(math.Max(float64(max_data.Linuxbuffer), float64(av.Linuxbuffer)))
			max_data.Linuxcached = int(math.Max(float64(max_data.Linuxcached), float64(av.Linuxcached)))
			max_data.Linuxsrec = int(math.Max(float64(max_data.Linuxsrec), float64(av.Linuxsrec)))
			max_data.Memused_mb = int(math.Max(float64(max_data.Memused_mb), float64(av.Memused_mb)))
			max_data.Irq = int(math.Max(float64(max_data.Irq), float64(av.Irq)))
			max_data.Softirq = int(math.Max(float64(max_data.Softirq), float64(av.Softirq)))
			max_data.Swapused_mb = int(math.Max(float64(max_data.Swapused_mb), float64(av.Swapused_mb)))
			max_data.Dusm = int(math.Max(float64(max_data.Dusm), float64(av.Dusm)))
		}
		if value_len := len(value.([]*data.Perf)); value_len > 0 {
			avg_data.User = avg_data.User / value_len
			avg_data.Sys = avg_data.Sys / value_len
			avg_data.Wait = avg_data.Wait / value_len
			avg_data.Idle = avg_data.Idle / value_len
			avg_data.Processorcount = avg_data.Processorcount / value_len
			avg_data.Runqueue = avg_data.Runqueue / value_len
			avg_data.Blockqueue = avg_data.Blockqueue / value_len
			avg_data.Waitqueue = avg_data.Waitqueue / value_len
			avg_data.Pqueue = avg_data.Pqueue / value_len
			avg_data.Pcrateuser = avg_data.Pcrateuser / value_len
			avg_data.Pcratesys = avg_data.Pcratesys / value_len
			avg_data.Memorysize = avg_data.Memorysize / value_len
			avg_data.Memoryused = avg_data.Memoryused / value_len
			avg_data.Memorypinned = avg_data.Memorypinned / value_len
			avg_data.Memorysys = avg_data.Memorysys / value_len
			avg_data.Memoryuser = avg_data.Memoryuser / value_len
			avg_data.Memorycache = avg_data.Memorycache / value_len
			avg_data.Avm = avg_data.Avm / value_len
			avg_data.Pagingspacein = avg_data.Pagingspacein / value_len
			avg_data.Pagingspaceout = avg_data.Pagingspaceout / value_len
			avg_data.Filesystemin = avg_data.Filesystemin / value_len
			avg_data.Filesystemout = avg_data.Filesystemout / value_len
			avg_data.Memoryscan = avg_data.Memoryscan / value_len
			avg_data.Memoryfreed = avg_data.Memoryfreed / value_len
			avg_data.Swapsize = avg_data.Swapsize / value_len
			avg_data.Swapused = avg_data.Swapused / value_len
			avg_data.Swapactive = avg_data.Swapactive / value_len
			avg_data.Fork = avg_data.Fork / value_len
			avg_data.Exec = avg_data.Exec / value_len
			avg_data.Interupt = avg_data.Interupt / value_len
			avg_data.Systemcall = avg_data.Systemcall / value_len
			avg_data.Contextswitch = avg_data.Contextswitch / value_len
			avg_data.Semaphore = avg_data.Semaphore / value_len
			avg_data.Msg = avg_data.Msg / value_len
			avg_data.Diskreadwrite = avg_data.Diskreadwrite / value_len
			avg_data.Diskiops = avg_data.Diskiops / value_len
			avg_data.Networkreadwrite = avg_data.Networkreadwrite / value_len
			avg_data.Networkiops = avg_data.Networkiops / value_len
			avg_data.Topcommandid = avg_data.Topcommandid / value_len
			avg_data.Topcommandcount = avg_data.Topcommandcount / value_len
			avg_data.Topuserid = avg_data.Topuserid / value_len
			avg_data.Topcpu = avg_data.Topcpu / value_len
			avg_data.Topdiskid = avg_data.Topdiskid / value_len
			avg_data.Topvgid = avg_data.Topvgid / value_len
			avg_data.Topbusy = avg_data.Topbusy / value_len
			avg_data.Maxpid = avg_data.Maxpid / value_len
			avg_data.Threadcount = avg_data.Threadcount / value_len
			avg_data.Pidcount = avg_data.Pidcount / value_len
			avg_data.Linuxbuffer = avg_data.Linuxbuffer / value_len
			avg_data.Linuxcached = avg_data.Linuxcached / value_len
			avg_data.Linuxsrec = avg_data.Linuxsrec / value_len
			avg_data.Memused_mb = avg_data.Memused_mb / value_len
			avg_data.Irq = avg_data.Irq / value_len
			avg_data.Softirq = avg_data.Softirq / value_len
			avg_data.Swapused_mb = avg_data.Swapused_mb / value_len
			avg_data.Dusm = avg_data.Dusm / value_len
		}
		agent_perf = append(agent_perf, avg_data)
		max_perf = append(max_perf, max_data)

		return true
	})

	GlobalChannel.DemoAvgPerfData <- agent_perf
	GlobalChannel.DemoAvgMaxPerfData <- max_perf
	time.Sleep(time.Millisecond * time.Duration(1))

	// Init
	d.AvgPerf = &sync.Map{}
}

func (d *DemoHandler) GenerateCpu(interval ConfigScrape) {
	for {
		ts := time.Now().Unix()

		var agent_cpu []*data.Cpu = make([]*data.Cpu, 0)
		for i := 0; i < d.Hostcount; i++ {
			for j := 0; j < CPU_CORE; j++ {
				demo_data := &data.Cpu{
					Ontunetime:    ts,
					Agenttime:     ts,
					Agentid:       d.Agentinfo.Agentid[i],
					Index:         j,
					User:          rand.Intn(100),
					Sys:           rand.Intn(100),
					Wait:          rand.Intn(100),
					Idle:          rand.Intn(100),
					Runqueue:      rand.Intn(100),
					Fork:          rand.Intn(100),
					Exec:          rand.Intn(100),
					Interupt:      rand.Intn(100),
					Systemcall:    rand.Intn(100),
					Contextswitch: rand.Intn(100),
				}
				agent_cpu = append(agent_cpu, demo_data)

				// Store Average Value
				key := fmt.Sprintf("%d_%d", demo_data.Agentid, demo_data.Index)
				if val, ok := d.AvgCpu.LoadOrStore(key, make([]*data.Cpu, 0)); ok {
					var avgvalue []*data.Cpu = make([]*data.Cpu, 0)
					avgvalue = append(avgvalue, val.([]*data.Cpu)...)
					d.AvgCpu.Store(key, avgvalue)
				}
			}
		}

		GlobalChannel.DemoCpuData <- agent_cpu
		time.Sleep(time.Second * time.Duration(interval.Rate))
	}
}

func (d *DemoHandler) GenerateAvgCpu() {
	var agent_cpu []*data.Cpu = make([]*data.Cpu, 0)
	var max_cpu []*data.Cpu = make([]*data.Cpu, 0)

	d.AvgCpu.Range(func(key, value any) bool {
		slice := strings.Split(key.(string), "_")
		agentid, err := strconv.Atoi(slice[0])
		ErrorCheck(err)
		index, err := strconv.Atoi(slice[1])
		ErrorCheck(err)

		avg_data := &data.Cpu{Agentid: agentid, Index: index}
		max_data := &data.Cpu{Agentid: agentid, Index: index}

		for _, av := range value.([]*data.Cpu) {
			avg_data.Ontunetime = int64(math.Max(float64(avg_data.Ontunetime), float64(av.Ontunetime)))
			avg_data.Agenttime = int64(math.Max(float64(avg_data.Agenttime), float64(av.Ontunetime)))
			avg_data.User = avg_data.User + av.User
			avg_data.Sys = avg_data.Sys + av.Sys
			avg_data.Wait = avg_data.Wait + av.Wait
			avg_data.Idle = avg_data.Idle + av.Idle
			avg_data.Runqueue = avg_data.Runqueue + av.Runqueue
			avg_data.Fork = avg_data.Fork + av.Fork
			avg_data.Exec = avg_data.Exec + av.Exec
			avg_data.Interupt = avg_data.Interupt + av.Interupt
			avg_data.Systemcall = avg_data.Systemcall + av.Systemcall
			avg_data.Contextswitch = avg_data.Contextswitch + av.Contextswitch

			max_data.Ontunetime = int64(math.Max(float64(max_data.Ontunetime), float64(av.Ontunetime)))
			max_data.Agenttime = int64(math.Max(float64(max_data.Agenttime), float64(av.Ontunetime)))
			max_data.User = int(math.Max(float64(max_data.User), float64(av.User)))
			max_data.Sys = int(math.Max(float64(max_data.Sys), float64(av.Sys)))
			max_data.Wait = int(math.Max(float64(max_data.Wait), float64(av.Wait)))
			max_data.Idle = int(math.Max(float64(max_data.Idle), float64(av.Idle)))
			max_data.Runqueue = int(math.Max(float64(max_data.Runqueue), float64(av.Runqueue)))
			max_data.Fork = int(math.Max(float64(max_data.Fork), float64(av.Fork)))
			max_data.Exec = int(math.Max(float64(max_data.Exec), float64(av.Exec)))
			max_data.Interupt = int(math.Max(float64(max_data.Interupt), float64(av.Interupt)))
			max_data.Systemcall = int(math.Max(float64(max_data.Systemcall), float64(av.Systemcall)))
			max_data.Contextswitch = int(math.Max(float64(max_data.Contextswitch), float64(av.Contextswitch)))

		}
		if value_len := len(value.([]*data.Cpu)); value_len > 0 {
			avg_data.User = avg_data.User / value_len
			avg_data.Sys = avg_data.Sys / value_len
			avg_data.Wait = avg_data.Wait / value_len
			avg_data.Idle = avg_data.Idle / value_len
			avg_data.Runqueue = avg_data.Runqueue / value_len
			avg_data.Fork = avg_data.Fork / value_len
			avg_data.Exec = avg_data.Exec / value_len
			avg_data.Interupt = avg_data.Interupt / value_len
			avg_data.Systemcall = avg_data.Systemcall / value_len
			avg_data.Contextswitch = avg_data.Contextswitch / value_len
		}
		agent_cpu = append(agent_cpu, avg_data)
		max_cpu = append(max_cpu, max_data)

		return true
	})

	GlobalChannel.DemoAvgCpuData <- agent_cpu
	GlobalChannel.DemoAvgMaxCpuData <- max_cpu
	time.Sleep(time.Millisecond * time.Duration(1))

	// Init
	d.AvgCpu = &sync.Map{}
}

func (d *DemoHandler) GenerateDisk(interval ConfigScrape) {
	for {
		ts := time.Now().Unix()

		var agent_disk []*data.Disk = make([]*data.Disk, 0)
		for i := 0; i < d.Hostcount; i++ {
			for ioid := range data.DISK_IONAME {
				demo_data := &data.Disk{
					Ontunetime:   ts,
					Agenttime:    ts,
					Agentid:      d.Agentinfo.Agentid[i],
					Ionameid:     ioid,
					Readrate:     rand.Intn(100),
					Writerate:    rand.Intn(100),
					Iops:         rand.Intn(100),
					Busy:         rand.Intn(100),
					Descid:       rand.Intn(50),
					Readsvctime:  rand.Intn(100),
					Writesvctime: rand.Intn(100),
				}
				agent_disk = append(agent_disk, demo_data)

				// Store Average Value
				key := fmt.Sprintf("%d_%d", demo_data.Agentid, demo_data.Ionameid)
				if val, ok := d.AvgDisk.LoadOrStore(key, make([]*data.Disk, 0)); ok {
					var avgvalue []*data.Disk = make([]*data.Disk, 0)
					avgvalue = append(avgvalue, val.([]*data.Disk)...)
					d.AvgDisk.Store(key, avgvalue)
				}
			}
		}

		GlobalChannel.DemoDiskData <- agent_disk
		time.Sleep(time.Second * time.Duration(interval.Rate))
	}
}

func (d *DemoHandler) GenerateAvgDisk() {
	var agent_disk []*data.Disk = make([]*data.Disk, 0)
	var max_disk []*data.Disk = make([]*data.Disk, 0)

	d.AvgDisk.Range(func(key, value any) bool {
		slice := strings.Split(key.(string), "_")
		agentid, err := strconv.Atoi(slice[0])
		ErrorCheck(err)
		ioid, err := strconv.Atoi(slice[1])
		ErrorCheck(err)

		avg_data := &data.Disk{Agentid: agentid, Ionameid: ioid}
		max_data := &data.Disk{Agentid: agentid, Ionameid: ioid}

		for _, av := range value.([]*data.Disk) {
			avg_data.Ontunetime = int64(math.Max(float64(avg_data.Ontunetime), float64(av.Ontunetime)))
			avg_data.Agenttime = int64(math.Max(float64(avg_data.Agenttime), float64(av.Ontunetime)))
			avg_data.Descid = av.Descid
			avg_data.Readrate = avg_data.Readrate + av.Readrate
			avg_data.Writerate = avg_data.Writerate + av.Writerate
			avg_data.Iops = avg_data.Iops + av.Iops
			avg_data.Busy = avg_data.Busy + av.Busy
			avg_data.Descid = avg_data.Descid + av.Descid
			avg_data.Readsvctime = avg_data.Readsvctime + av.Readsvctime
			avg_data.Writesvctime = avg_data.Writesvctime + av.Writesvctime

			max_data.Ontunetime = int64(math.Max(float64(max_data.Ontunetime), float64(av.Ontunetime)))
			max_data.Agenttime = int64(math.Max(float64(max_data.Agenttime), float64(av.Ontunetime)))
			max_data.Descid = av.Descid
			max_data.Readrate = int(math.Max(float64(max_data.Readrate), float64(av.Readrate)))
			max_data.Writerate = int(math.Max(float64(max_data.Writerate), float64(av.Writerate)))
			max_data.Iops = int(math.Max(float64(max_data.Iops), float64(av.Iops)))
			max_data.Busy = int(math.Max(float64(max_data.Busy), float64(av.Busy)))
			max_data.Descid = int(math.Max(float64(max_data.Descid), float64(av.Descid)))
			max_data.Readsvctime = int(math.Max(float64(max_data.Readsvctime), float64(av.Readsvctime)))
			max_data.Writesvctime = int(math.Max(float64(max_data.Writesvctime), float64(av.Writesvctime)))

		}
		if value_len := len(value.([]*data.Disk)); value_len > 0 {
			avg_data.Readrate = avg_data.Readrate / value_len
			avg_data.Writerate = avg_data.Writerate / value_len
			avg_data.Iops = avg_data.Iops / value_len
			avg_data.Busy = avg_data.Busy / value_len
			avg_data.Descid = avg_data.Descid / value_len
			avg_data.Readsvctime = avg_data.Readsvctime / value_len
			avg_data.Writesvctime = avg_data.Writesvctime / value_len
		}
		agent_disk = append(agent_disk, avg_data)
		max_disk = append(max_disk, max_data)

		return true
	})

	GlobalChannel.DemoAvgDiskData <- agent_disk
	GlobalChannel.DemoAvgMaxDiskData <- max_disk
	time.Sleep(time.Millisecond * time.Duration(1))

	// Init
	d.AvgDisk = &sync.Map{}
}

func (d *DemoHandler) GenerateNet(interval ConfigScrape) {
	for {
		ts := time.Now().Unix()

		var agent_net []*data.Net = make([]*data.Net, 0)
		for i := 0; i < d.Hostcount; i++ {
			for ioid := range data.NET_IONAME {
				demo_data := &data.Net{
					Ontunetime: ts,
					Agenttime:  ts,
					Agentid:    d.Agentinfo.Agentid[i],
					Ionameid:   ioid,
					Readrate:   rand.Intn(100),
					Writerate:  rand.Intn(100),
					Readiops:   rand.Intn(100),
					Writeiops:  rand.Intn(100),
					Errorps:    rand.Intn(100),
					Collision:  rand.Intn(100),
				}
				agent_net = append(agent_net, demo_data)

				// Store Average Value
				key := fmt.Sprintf("%d_%d", demo_data.Agentid, demo_data.Ionameid)
				if val, ok := d.AvgNet.LoadOrStore(key, make([]*data.Net, 0)); ok {
					var avgvalue []*data.Net = make([]*data.Net, 0)
					avgvalue = append(avgvalue, val.([]*data.Net)...)
					d.AvgNet.Store(key, avgvalue)
				}
			}
		}

		GlobalChannel.DemoNetData <- agent_net
		time.Sleep(time.Second * time.Duration(interval.Rate))
	}
}

func (d *DemoHandler) GenerateAvgNet() {
	var agent_net []*data.Net = make([]*data.Net, 0)
	var max_net []*data.Net = make([]*data.Net, 0)

	d.AvgNet.Range(func(key, value any) bool {
		slice := strings.Split(key.(string), "_")
		agentid, err := strconv.Atoi(slice[0])
		ErrorCheck(err)
		ioid, err := strconv.Atoi(slice[1])
		ErrorCheck(err)

		avg_data := &data.Net{Agentid: agentid, Ionameid: ioid}
		max_data := &data.Net{Agentid: agentid, Ionameid: ioid}

		for _, av := range value.([]*data.Net) {
			avg_data.Ontunetime = int64(math.Max(float64(avg_data.Ontunetime), float64(av.Ontunetime)))
			avg_data.Agenttime = int64(math.Max(float64(avg_data.Agenttime), float64(av.Ontunetime)))
			avg_data.Readrate = avg_data.Readrate + av.Readrate
			avg_data.Writerate = avg_data.Writerate + av.Writerate
			avg_data.Readiops = avg_data.Readiops + av.Readiops
			avg_data.Writeiops = avg_data.Writeiops + av.Writeiops
			avg_data.Errorps = avg_data.Errorps + av.Errorps
			avg_data.Collision = avg_data.Collision + av.Collision

			max_data.Ontunetime = int64(math.Max(float64(max_data.Ontunetime), float64(av.Ontunetime)))
			max_data.Agenttime = int64(math.Max(float64(max_data.Agenttime), float64(av.Ontunetime)))
			max_data.Readrate = int(math.Max(float64(max_data.Readrate), float64(av.Readrate)))
			max_data.Writerate = int(math.Max(float64(max_data.Writerate), float64(av.Writerate)))
			max_data.Readiops = int(math.Max(float64(max_data.Readiops), float64(av.Readiops)))
			max_data.Writeiops = int(math.Max(float64(max_data.Writeiops), float64(av.Writeiops)))
			max_data.Errorps = int(math.Max(float64(max_data.Errorps), float64(av.Errorps)))
			max_data.Collision = int(math.Max(float64(max_data.Collision), float64(av.Collision)))
		}
		if value_len := len(value.([]*data.Net)); value_len > 0 {
			avg_data.Readrate = avg_data.Readrate / value_len
			avg_data.Writerate = avg_data.Writerate / value_len
			avg_data.Readiops = avg_data.Readiops / value_len
			avg_data.Writeiops = avg_data.Writeiops / value_len
			avg_data.Errorps = avg_data.Errorps / value_len
			avg_data.Collision = avg_data.Collision / value_len
		}
		agent_net = append(agent_net, avg_data)
		max_net = append(agent_net, max_data)

		return true
	})

	GlobalChannel.DemoAvgNetData <- agent_net
	GlobalChannel.DemoAvgMaxNetData <- max_net

	time.Sleep(time.Millisecond * time.Duration(1))

	// Init
	d.AvgNet = &sync.Map{}
}

func (d *DemoHandler) GenerateProc(interval ConfigScrape) {
	for {
		ts := time.Now().Unix()

		var agent_proc []*data.Pid = make([]*data.Pid, 0)
		for i := 0; i < d.Hostcount; i++ {
			// 위 구조까지 포함하면 5중 For문이지만,
			// 기본값으로 pc-3, pu-1, pa-3 이므로 실제 Loop 회수는 9회임
			for _, pc := range PROCCMD_ARR {
				for _, pu := range PROCUSER_ARR {
					for _, pa := range PROCARG_ARR {
						demo_data := &data.Pid{
							Ontunetime: ts,
							Agenttime:  ts,
							Agentid:    d.Agentinfo.Agentid[i],
							Pid:        rand.Intn(100),
							Ppid:       rand.Intn(100),
							Uid:        rand.Intn(100),
							Cmdid:      pc,
							Userid:     pu,
							Argid:      pa,
							Usr:        rand.Intn(100),
							Sys:        rand.Intn(100),
							Usrsys:     rand.Intn(100),
							Sz:         rand.Intn(100),
							Rss:        rand.Intn(100),
							Vmem:       rand.Intn(100),
							Chario:     rand.Intn(100),
							Processcnt: rand.Intn(100),
							Threadcnt:  rand.Intn(100),
							Handlecnt:  rand.Intn(100),
							Stime:      rand.Intn(100),
							Pvbytes:    rand.Intn(100),
							Pgpool:     rand.Intn(100),
						}
						agent_proc = append(agent_proc, demo_data)

						// Store Average Value
						key := fmt.Sprintf("%d_%d_%d_%d", demo_data.Agentid, demo_data.Cmdid, demo_data.Userid, demo_data.Argid)
						if val, ok := d.AvgProc.LoadOrStore(key, make([]*data.Pid, 0)); ok {
							var avgvalue []*data.Pid = make([]*data.Pid, 0)
							avgvalue = append(avgvalue, val.([]*data.Pid)...)
							d.AvgProc.Store(key, avgvalue)
						}
					}
				}
			}
		}

		GlobalChannel.DemoProcData <- agent_proc
		time.Sleep(time.Second * time.Duration(interval.Rate))
	}
}

func (d *DemoHandler) GenerateAvgProc() {
	var agent_proc []*data.Pid = make([]*data.Pid, 0)
	var max_proc []*data.Pid = make([]*data.Pid, 0)

	d.AvgProc.Range(func(key, value any) bool {
		slice := strings.Split(key.(string), "_")
		agentid, err := strconv.Atoi(slice[0])
		ErrorCheck(err)
		cmdid, err := strconv.Atoi(slice[1])
		ErrorCheck(err)
		userid, err := strconv.Atoi(slice[2])
		ErrorCheck(err)
		argid, err := strconv.Atoi(slice[3])
		ErrorCheck(err)

		avg_data := &data.Pid{
			Agentid: agentid,
			Cmdid:   cmdid,
			Userid:  userid,
			Argid:   argid,
		}

		max_data := &data.Pid{
			Agentid: agentid,
			Cmdid:   cmdid,
			Userid:  userid,
			Argid:   argid,
		}

		// PID, PPID, UID는 Key 값은 아니지만 평균값으로 낼 데이터도 아니므로 최대값으로 잡음
		for _, av := range value.([]*data.Pid) {
			avg_data.Ontunetime = int64(math.Max(float64(avg_data.Ontunetime), float64(av.Ontunetime)))
			avg_data.Agenttime = int64(math.Max(float64(avg_data.Agenttime), float64(av.Ontunetime)))
			avg_data.Pid = int(math.Max(float64(avg_data.Pid), float64(av.Pid)))
			avg_data.Ppid = int(math.Max(float64(avg_data.Ppid), float64(av.Ppid)))
			avg_data.Uid = int(math.Max(float64(avg_data.Uid), float64(av.Uid)))
			avg_data.Usr = avg_data.Usr + av.Usr
			avg_data.Sys = avg_data.Sys + av.Sys
			avg_data.Usrsys = avg_data.Usrsys + av.Usrsys
			avg_data.Sz = avg_data.Sz + av.Sz
			avg_data.Rss = avg_data.Rss + av.Rss
			avg_data.Vmem = avg_data.Vmem + av.Vmem
			avg_data.Chario = avg_data.Chario + av.Chario
			avg_data.Processcnt = avg_data.Processcnt + av.Processcnt
			avg_data.Threadcnt = avg_data.Threadcnt + av.Threadcnt
			avg_data.Handlecnt = avg_data.Handlecnt + av.Handlecnt
			avg_data.Stime = avg_data.Stime + av.Stime
			avg_data.Pvbytes = avg_data.Pvbytes + av.Pvbytes
			avg_data.Pgpool = avg_data.Pgpool + av.Pgpool

			max_data.Ontunetime = int64(math.Max(float64(max_data.Ontunetime), float64(av.Ontunetime)))
			max_data.Agenttime = int64(math.Max(float64(max_data.Agenttime), float64(av.Ontunetime)))
			max_data.Pid = int(math.Max(float64(max_data.Pid), float64(av.Pid)))
			max_data.Ppid = int(math.Max(float64(max_data.Ppid), float64(av.Ppid)))
			max_data.Uid = int(math.Max(float64(max_data.Uid), float64(av.Uid)))
			max_data.Usr = int(math.Max(float64(max_data.Usr), float64(av.Usr)))
			max_data.Sys = int(math.Max(float64(max_data.Sys), float64(av.Sys)))
			max_data.Usrsys = int(math.Max(float64(max_data.Usrsys), float64(av.Usrsys)))
			max_data.Sz = int(math.Max(float64(max_data.Sz), float64(av.Sz)))
			max_data.Rss = int(math.Max(float64(max_data.Rss), float64(av.Rss)))
			max_data.Vmem = int(math.Max(float64(max_data.Vmem), float64(av.Vmem)))
			max_data.Chario = int(math.Max(float64(max_data.Chario), float64(av.Chario)))
			max_data.Processcnt = int(math.Max(float64(max_data.Processcnt), float64(av.Processcnt)))
			max_data.Threadcnt = int(math.Max(float64(max_data.Threadcnt), float64(av.Threadcnt)))
			max_data.Handlecnt = int(math.Max(float64(max_data.Handlecnt), float64(av.Handlecnt)))
			max_data.Stime = int(math.Max(float64(max_data.Stime), float64(av.Stime)))
			max_data.Pvbytes = int(math.Max(float64(max_data.Pvbytes), float64(av.Pvbytes)))
			max_data.Pgpool = int(math.Max(float64(max_data.Pgpool), float64(av.Pgpool)))
		}
		if value_len := len(value.([]*data.Pid)); value_len > 0 {
			avg_data.Usr = avg_data.Usr / value_len
			avg_data.Sys = avg_data.Sys / value_len
			avg_data.Usrsys = avg_data.Usrsys / value_len
			avg_data.Sz = avg_data.Sz / value_len
			avg_data.Rss = avg_data.Rss / value_len
			avg_data.Vmem = avg_data.Vmem / value_len
			avg_data.Chario = avg_data.Chario / value_len
			avg_data.Processcnt = avg_data.Processcnt / value_len
			avg_data.Threadcnt = avg_data.Threadcnt / value_len
			avg_data.Handlecnt = avg_data.Handlecnt / value_len
			avg_data.Stime = avg_data.Stime / value_len
			avg_data.Pvbytes = avg_data.Pvbytes / value_len
			avg_data.Pgpool = avg_data.Pgpool / value_len
		}
		agent_proc = append(agent_proc, avg_data)
		max_proc = append(agent_proc, max_data)

		return true
	})

	GlobalChannel.DemoAvgProcData <- agent_proc
	GlobalChannel.DemoAvgMaxProcData <- max_proc
	time.Sleep(time.Millisecond * time.Duration(1))

	// Init
	d.AvgProc = &sync.Map{}
}
