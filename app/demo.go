package app

import (
	"fmt"
	"math"
	"math/rand"
	"mgr/data"
	"sync"
	"time"
)

type DemoHandler struct {
	Hostcount    int
	Interval     int
	Agentinfo    data.AgentinfoArr
	Hostinfo     data.HostinfoArr
	AvgBasicPerf *sync.Map
}

func (d *DemoHandler) Init(hostcount int) {
	d.Hostcount = hostcount
	d.AvgBasicPerf = &sync.Map{}

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
			Shorttermbasic:    2,
			Shorttermproc:     5,
			Shorttermio:       5,
			Shorttermcpu:      5,
			Longtermbasic:     600,
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
			Virbasicperf:      1,
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

		MapHostifo[i+1] = fmt.Sprintf("DummyAgent%d", i+1)
	}
}

func (d *DemoHandler) GenerateBasicPerf(interval ConfigScrape) {
	for {
		ts := time.Now().Unix()

		var agent_basicperf []*data.Basicperf = make([]*data.Basicperf, 0)
		for i := 0; i < d.Hostcount; i++ {
			demo_data := &data.Basicperf{
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
			agent_basicperf = append(agent_basicperf, demo_data)

			// Store Average Value
			if val, ok := d.AvgBasicPerf.LoadOrStore(demo_data.Agentid, make([]*data.Basicperf, 0)); ok {
				var avgvalue []*data.Basicperf = make([]*data.Basicperf, 0)
				avgvalue = append(avgvalue, val.([]*data.Basicperf)...)
				d.AvgBasicPerf.Store(demo_data.Agentid, avgvalue)
			}
		}

		GlobalChannel.DemoBasicData <- agent_basicperf
		time.Sleep(time.Second * time.Duration(interval.Rate))
	}
}

func (d *DemoHandler) GenerateAvgBasicPerf() {
	var agent_basicperf []*data.Basicperf = make([]*data.Basicperf, 0)

	d.AvgBasicPerf.Range(func(key, value any) bool {
		avg_data := &data.Basicperf{
			Agentid: key.(int),
		}
		for _, av := range value.([]*data.Basicperf) {
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
		}
		if value_len := len(value.([]*data.Basicperf)); value_len > 0 {
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
		agent_basicperf = append(agent_basicperf, avg_data)

		return true
	})

	GlobalChannel.DemoAvgBasicData <- agent_basicperf
	time.Sleep(time.Millisecond * time.Duration(1))

	// Init
	d.AvgBasicPerf = &sync.Map{}
}
