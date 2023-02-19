package app

import (
	"fmt"
	"math"
	"math/rand"
	"mgr/data"
	"time"
)

type DemoHandler struct {
	Hostcount    int
	Interval     int
	Agentinfo    data.AgentinfoArr
	Hostinfo     data.HostinfoArr
	AvgBasicPerf map[int][]*data.Basicperf
}

func (d *DemoHandler) Init(hostcount int) {
	d.Hostcount = hostcount
	d.AvgBasicPerf = make(map[int][]*data.Basicperf)

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

		GlobalMutex.DemoAvgBasicM.Lock()
		var agent_basicperf []data.Basicperf = make([]data.Basicperf, 0)
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
			agent_basicperf = append(agent_basicperf, *demo_data)

			if _, ok := d.AvgBasicPerf[demo_data.Agentid]; !ok {
				d.AvgBasicPerf[demo_data.Agentid] = make([]*data.Basicperf, 0)
			}
			d.AvgBasicPerf[demo_data.Agentid] = append(d.AvgBasicPerf[demo_data.Agentid], demo_data)
		}
		GlobalMutex.DemoAvgBasicM.Unlock()

		GlobalChannel.DemoBasicData <- agent_basicperf
		time.Sleep(time.Second * time.Duration(interval.Rate))
	}
}

func (d *DemoHandler) GenerateAvgBasicPerf() {
	var agent_basicperf []data.Basicperf = make([]data.Basicperf, 0)

	GlobalMutex.DemoAvgBasicM.Lock()
	for aid, adata := range d.AvgBasicPerf {
		avg_data := &data.Basicperf{
			Agentid: aid,
		}
		for _, aadata := range adata {
			avg_data.Ontunetime = int64(math.Max(float64(avg_data.Ontunetime), float64(aadata.Ontunetime)))
			avg_data.Agenttime = int64(math.Max(float64(avg_data.Agenttime), float64(aadata.Ontunetime)))
			avg_data.User = avg_data.User + aadata.User
			avg_data.Sys = avg_data.Sys + aadata.Sys
			avg_data.Wait = avg_data.Wait + aadata.Wait
			avg_data.Idle = avg_data.Idle + aadata.Idle
			avg_data.Processorcount = avg_data.Processorcount + aadata.Processorcount
			avg_data.Runqueue = avg_data.Runqueue + aadata.Runqueue
			avg_data.Blockqueue = avg_data.Blockqueue + aadata.Blockqueue
			avg_data.Waitqueue = avg_data.Waitqueue + aadata.Waitqueue
			avg_data.Pqueue = avg_data.Pqueue + aadata.Pqueue
			avg_data.Pcrateuser = avg_data.Pcrateuser + aadata.Pcrateuser
			avg_data.Pcratesys = avg_data.Pcratesys + aadata.Pcratesys
			avg_data.Memorysize = avg_data.Memorysize + aadata.Memorysize
			avg_data.Memoryused = avg_data.Memoryused + aadata.Memoryused
			avg_data.Memorypinned = avg_data.Memorypinned + aadata.Memorypinned
			avg_data.Memorysys = avg_data.Memorysys + aadata.Memorysys
			avg_data.Memoryuser = avg_data.Memoryuser + aadata.Memoryuser
			avg_data.Memorycache = avg_data.Memorycache + aadata.Memorycache
			avg_data.Avm = avg_data.Avm + aadata.Avm
			avg_data.Pagingspacein = avg_data.Pagingspacein + aadata.Pagingspacein
			avg_data.Pagingspaceout = avg_data.Pagingspaceout + aadata.Pagingspaceout
			avg_data.Filesystemin = avg_data.Filesystemin + aadata.Filesystemin
			avg_data.Filesystemout = avg_data.Filesystemout + aadata.Filesystemout
			avg_data.Memoryscan = avg_data.Memoryscan + aadata.Memoryscan
			avg_data.Memoryfreed = avg_data.Memoryfreed + aadata.Memoryfreed
			avg_data.Swapsize = avg_data.Swapsize + aadata.Swapsize
			avg_data.Swapused = avg_data.Swapused + aadata.Swapused
			avg_data.Swapactive = avg_data.Swapactive + aadata.Swapactive
			avg_data.Fork = avg_data.Fork + aadata.Fork
			avg_data.Exec = avg_data.Exec + aadata.Exec
			avg_data.Interupt = avg_data.Interupt + aadata.Interupt
			avg_data.Systemcall = avg_data.Systemcall + aadata.Systemcall
			avg_data.Contextswitch = avg_data.Contextswitch + aadata.Contextswitch
			avg_data.Semaphore = avg_data.Semaphore + aadata.Semaphore
			avg_data.Msg = avg_data.Msg + aadata.Msg
			avg_data.Diskreadwrite = avg_data.Diskreadwrite + aadata.Diskreadwrite
			avg_data.Diskiops = avg_data.Diskiops + aadata.Diskiops
			avg_data.Networkreadwrite = avg_data.Networkreadwrite + aadata.Networkreadwrite
			avg_data.Networkiops = avg_data.Networkiops + aadata.Networkiops
			avg_data.Topcommandid = avg_data.Topcommandid + aadata.Topcommandid
			avg_data.Topcommandcount = avg_data.Topcommandcount + aadata.Topcommandcount
			avg_data.Topuserid = avg_data.Topuserid + aadata.Topuserid
			avg_data.Topcpu = avg_data.Topcpu + aadata.Topcpu
			avg_data.Topdiskid = avg_data.Topdiskid + aadata.Topdiskid
			avg_data.Topvgid = avg_data.Topvgid + aadata.Topvgid
			avg_data.Topbusy = avg_data.Topbusy + aadata.Topbusy
			avg_data.Maxpid = avg_data.Maxpid + aadata.Maxpid
			avg_data.Threadcount = avg_data.Threadcount + aadata.Threadcount
			avg_data.Pidcount = avg_data.Pidcount + aadata.Pidcount
			avg_data.Linuxbuffer = avg_data.Linuxbuffer + aadata.Linuxbuffer
			avg_data.Linuxcached = avg_data.Linuxcached + aadata.Linuxcached
			avg_data.Linuxsrec = avg_data.Linuxsrec + aadata.Linuxsrec
			avg_data.Memused_mb = avg_data.Memused_mb + aadata.Memused_mb
			avg_data.Irq = avg_data.Irq + aadata.Irq
			avg_data.Softirq = avg_data.Softirq + aadata.Softirq
			avg_data.Swapused_mb = avg_data.Swapused_mb + aadata.Swapused_mb
			avg_data.Dusm = avg_data.Dusm + aadata.Dusm
		}
		avg_data.User = avg_data.User / len(adata)
		avg_data.Sys = avg_data.Sys / len(adata)
		avg_data.Wait = avg_data.Wait / len(adata)
		avg_data.Idle = avg_data.Idle / len(adata)
		avg_data.Processorcount = avg_data.Processorcount / len(adata)
		avg_data.Runqueue = avg_data.Runqueue / len(adata)
		avg_data.Blockqueue = avg_data.Blockqueue / len(adata)
		avg_data.Waitqueue = avg_data.Waitqueue / len(adata)
		avg_data.Pqueue = avg_data.Pqueue / len(adata)
		avg_data.Pcrateuser = avg_data.Pcrateuser / len(adata)
		avg_data.Pcratesys = avg_data.Pcratesys / len(adata)
		avg_data.Memorysize = avg_data.Memorysize / len(adata)
		avg_data.Memoryused = avg_data.Memoryused / len(adata)
		avg_data.Memorypinned = avg_data.Memorypinned / len(adata)
		avg_data.Memorysys = avg_data.Memorysys / len(adata)
		avg_data.Memoryuser = avg_data.Memoryuser / len(adata)
		avg_data.Memorycache = avg_data.Memorycache / len(adata)
		avg_data.Avm = avg_data.Avm / len(adata)
		avg_data.Pagingspacein = avg_data.Pagingspacein / len(adata)
		avg_data.Pagingspaceout = avg_data.Pagingspaceout / len(adata)
		avg_data.Filesystemin = avg_data.Filesystemin / len(adata)
		avg_data.Filesystemout = avg_data.Filesystemout / len(adata)
		avg_data.Memoryscan = avg_data.Memoryscan / len(adata)
		avg_data.Memoryfreed = avg_data.Memoryfreed / len(adata)
		avg_data.Swapsize = avg_data.Swapsize / len(adata)
		avg_data.Swapused = avg_data.Swapused / len(adata)
		avg_data.Swapactive = avg_data.Swapactive / len(adata)
		avg_data.Fork = avg_data.Fork / len(adata)
		avg_data.Exec = avg_data.Exec / len(adata)
		avg_data.Interupt = avg_data.Interupt / len(adata)
		avg_data.Systemcall = avg_data.Systemcall / len(adata)
		avg_data.Contextswitch = avg_data.Contextswitch / len(adata)
		avg_data.Semaphore = avg_data.Semaphore / len(adata)
		avg_data.Msg = avg_data.Msg / len(adata)
		avg_data.Diskreadwrite = avg_data.Diskreadwrite / len(adata)
		avg_data.Diskiops = avg_data.Diskiops / len(adata)
		avg_data.Networkreadwrite = avg_data.Networkreadwrite / len(adata)
		avg_data.Networkiops = avg_data.Networkiops / len(adata)
		avg_data.Topcommandid = avg_data.Topcommandid / len(adata)
		avg_data.Topcommandcount = avg_data.Topcommandcount / len(adata)
		avg_data.Topuserid = avg_data.Topuserid / len(adata)
		avg_data.Topcpu = avg_data.Topcpu / len(adata)
		avg_data.Topdiskid = avg_data.Topdiskid / len(adata)
		avg_data.Topvgid = avg_data.Topvgid / len(adata)
		avg_data.Topbusy = avg_data.Topbusy / len(adata)
		avg_data.Maxpid = avg_data.Maxpid / len(adata)
		avg_data.Threadcount = avg_data.Threadcount / len(adata)
		avg_data.Pidcount = avg_data.Pidcount / len(adata)
		avg_data.Linuxbuffer = avg_data.Linuxbuffer / len(adata)
		avg_data.Linuxcached = avg_data.Linuxcached / len(adata)
		avg_data.Linuxsrec = avg_data.Linuxsrec / len(adata)
		avg_data.Memused_mb = avg_data.Memused_mb / len(adata)
		avg_data.Irq = avg_data.Irq / len(adata)
		avg_data.Softirq = avg_data.Softirq / len(adata)
		avg_data.Swapused_mb = avg_data.Swapused_mb / len(adata)
		avg_data.Dusm = avg_data.Dusm / len(adata)

		agent_basicperf = append(agent_basicperf, *avg_data)
	}

	GlobalChannel.DemoAvgBasicData <- agent_basicperf
	time.Sleep(time.Millisecond * time.Duration(1))

	// Init
	d.AvgBasicPerf = make(map[int][]*data.Basicperf)
	GlobalMutex.DemoAvgBasicM.Unlock()
}
