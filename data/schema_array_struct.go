package data

import (
	"time"

	"github.com/lib/pq"
)

type AgentinfoArr struct {
	Agentid           []int
	Agentname         []string
	Enabled           []int
	Connected         []int
	Updated           []int
	Shorttermbasic    []int
	Shorttermproc     []int
	Shorttermio       []int
	Shorttermcpu      []int
	Longtermbasic     []int
	Longtermproc      []int
	Longtermio        []int
	Longtermcpu       []int
	Model             []string
	Serial            []string
	Group             []string
	Ipaddress         []string
	Pscommand         []string
	Logevent          []string
	Processevent      []string
	Timecheck         []int
	Disconnectedtime  []int64
	Skipdatatypes     []int
	Virbasicperf      []int
	Hypervisor        []int
	Serviceevent      []string
	Installdate       []int64
	Lastconnectedtime []int64
}

func (a *AgentinfoArr) SetData(i Agentinfo) {
	a.Agentid = append(a.Agentid, i.Agentid)
	a.Agentname = append(a.Agentname, i.Agentname)
	a.Enabled = append(a.Enabled, i.Enabled)
	a.Connected = append(a.Connected, i.Connected)
	a.Updated = append(a.Updated, i.Updated)
	a.Shorttermbasic = append(a.Shorttermbasic, i.Shorttermbasic)
	a.Shorttermproc = append(a.Shorttermproc, i.Shorttermproc)
	a.Shorttermio = append(a.Shorttermio, i.Shorttermio)
	a.Shorttermcpu = append(a.Shorttermcpu, i.Shorttermcpu)
	a.Longtermbasic = append(a.Longtermbasic, i.Longtermbasic)
	a.Longtermproc = append(a.Longtermproc, i.Longtermproc)
	a.Longtermio = append(a.Longtermio, i.Longtermio)
	a.Longtermcpu = append(a.Longtermcpu, i.Longtermcpu)
	a.Model = append(a.Model, i.Model)
	a.Serial = append(a.Serial, i.Serial)
	a.Group = append(a.Group, i.Group)
	a.Ipaddress = append(a.Ipaddress, i.Ipaddress)
	a.Pscommand = append(a.Pscommand, i.Pscommand)
	a.Logevent = append(a.Logevent, i.Logevent)
	a.Processevent = append(a.Processevent, i.Processevent)
	a.Timecheck = append(a.Timecheck, i.Timecheck)
	a.Disconnectedtime = append(a.Disconnectedtime, i.Disconnectedtime)
	a.Skipdatatypes = append(a.Skipdatatypes, i.Skipdatatypes)
	a.Virbasicperf = append(a.Virbasicperf, i.Virbasicperf)
	a.Hypervisor = append(a.Hypervisor, i.Hypervisor)
	a.Serviceevent = append(a.Serviceevent, i.Serviceevent)
	a.Installdate = append(a.Installdate, i.Installdate)
	a.Lastconnectedtime = append(a.Lastconnectedtime, i.Lastconnectedtime)
}

func (a *AgentinfoArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(a.Agentid))
	data = append(data, pq.StringArray(a.Agentname))
	data = append(data, pq.Array(a.Enabled))
	data = append(data, pq.Array(a.Connected))
	data = append(data, pq.Array(a.Updated))
	data = append(data, pq.Array(a.Shorttermbasic))
	data = append(data, pq.Array(a.Shorttermproc))
	data = append(data, pq.Array(a.Shorttermio))
	data = append(data, pq.Array(a.Shorttermcpu))
	data = append(data, pq.Array(a.Longtermbasic))
	data = append(data, pq.Array(a.Longtermproc))
	data = append(data, pq.Array(a.Longtermio))
	data = append(data, pq.Array(a.Longtermcpu))
	data = append(data, pq.StringArray(a.Model))
	data = append(data, pq.StringArray(a.Serial))
	data = append(data, pq.StringArray(a.Group))
	data = append(data, pq.StringArray(a.Ipaddress))
	data = append(data, pq.StringArray(a.Pscommand))
	data = append(data, pq.StringArray(a.Logevent))
	data = append(data, pq.StringArray(a.Processevent))
	data = append(data, pq.Array(a.Timecheck))
	data = append(data, pq.Array(a.Disconnectedtime))
	data = append(data, pq.Array(a.Skipdatatypes))
	data = append(data, pq.Array(a.Virbasicperf))
	data = append(data, pq.Array(a.Hypervisor))
	data = append(data, pq.StringArray(a.Serviceevent))
	data = append(data, pq.Array(a.Installdate))
	data = append(data, pq.Array(a.Lastconnectedtime))

	return data
}

type HostinfoArr struct {
	Agentid        []int
	Hostname       []string
	Hostnameext    []string
	Os             []string
	Fw             []string
	Agentversion   []string
	Model          []string
	Serial         []string
	Processorcount []int
	Processorclock []int
	Memorysize     []int
	Swapsize       []int
	Poolid         []int
	Replication    []int
	Smt            []int
	Micropar       []int
	Capped         []int
	Ec             []int
	Virtualcpu     []int
	Weight         []int
	Cpupool        []int
	Ams            []int
	Allip          []string
	Numanodecount  []int
}

func (a *HostinfoArr) SetData(h Hostinfo) {
	a.Agentid = append(a.Agentid, h.Agentid)
	a.Hostname = append(a.Hostname, h.Hostname)
	a.Hostnameext = append(a.Hostnameext, h.Hostnameext)
	a.Os = append(a.Os, h.Os)
	a.Fw = append(a.Fw, h.Fw)
	a.Agentversion = append(a.Agentversion, h.Agentversion)
	a.Model = append(a.Model, h.Model)
	a.Serial = append(a.Serial, h.Serial)
	a.Processorcount = append(a.Processorcount, h.Processorcount)
	a.Processorclock = append(a.Processorclock, h.Processorclock)
	a.Memorysize = append(a.Memorysize, h.Memorysize)
	a.Swapsize = append(a.Swapsize, h.Swapsize)
	a.Poolid = append(a.Poolid, h.Poolid)
	a.Replication = append(a.Replication, h.Replication)
	a.Smt = append(a.Smt, h.Smt)
	a.Micropar = append(a.Micropar, h.Micropar)
	a.Capped = append(a.Capped, h.Capped)
	a.Ec = append(a.Ec, h.Ec)
	a.Virtualcpu = append(a.Virtualcpu, h.Virtualcpu)
	a.Weight = append(a.Weight, h.Weight)
	a.Cpupool = append(a.Cpupool, h.Cpupool)
	a.Ams = append(a.Ams, h.Ams)
	a.Allip = append(a.Allip, h.Allip)
	a.Numanodecount = append(a.Numanodecount, h.Numanodecount)
}

func (a *HostinfoArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(a.Agentid))
	data = append(data, pq.StringArray(a.Hostname))
	data = append(data, pq.StringArray(a.Hostnameext))
	data = append(data, pq.StringArray(a.Os))
	data = append(data, pq.StringArray(a.Fw))
	data = append(data, pq.StringArray(a.Agentversion))
	data = append(data, pq.StringArray(a.Model))
	data = append(data, pq.StringArray(a.Serial))
	data = append(data, pq.Array(a.Processorcount))
	data = append(data, pq.Array(a.Processorclock))
	data = append(data, pq.Array(a.Memorysize))
	data = append(data, pq.Array(a.Swapsize))
	data = append(data, pq.Array(a.Poolid))
	data = append(data, pq.Array(a.Replication))
	data = append(data, pq.Array(a.Smt))
	data = append(data, pq.Array(a.Micropar))
	data = append(data, pq.Array(a.Capped))
	data = append(data, pq.Array(a.Ec))
	data = append(data, pq.Array(a.Virtualcpu))
	data = append(data, pq.Array(a.Weight))
	data = append(data, pq.Array(a.Cpupool))
	data = append(data, pq.Array(a.Ams))
	data = append(data, pq.StringArray(a.Allip))
	data = append(data, pq.Array(a.Numanodecount))

	return data
}

type BasicperfArr struct {
	Ontunetime       []int64
	Agenttime        []int64
	Agentid          []int
	User             []int
	Sys              []int
	Wait             []int
	Idle             []int
	Processorcount   []int
	Runqueue         []int
	Blockqueue       []int
	Waitqueue        []int
	Pqueue           []int
	Pcrateuser       []int
	Pcratesys        []int
	Memorysize       []int
	Memoryused       []int
	Memorypinned     []int
	Memorysys        []int
	Memoryuser       []int
	Memorycache      []int
	Avm              []int
	Pagingspacein    []int
	Pagingspaceout   []int
	Filesystemin     []int
	Filesystemout    []int
	Memoryscan       []int
	Memoryfreed      []int
	Swapsize         []int
	Swapused         []int
	Swapactive       []int
	Fork             []int
	Exec             []int
	Interupt         []int
	Systemcall       []int
	Contextswitch    []int
	Semaphore        []int
	Msg              []int
	Diskreadwrite    []int
	Diskiops         []int
	Networkreadwrite []int
	Networkiops      []int
	Topcommandid     []int
	Topcommandcount  []int
	Topuserid        []int
	Topcpu           []int
	Topdiskid        []int
	Topvgid          []int
	Topbusy          []int
	Maxpid           []int
	Threadcount      []int
	Pidcount         []int
	Linuxbuffer      []int
	Linuxcached      []int
	Linuxsrec        []int
	Memused_mb       []int
	Irq              []int
	Softirq          []int
	Swapused_mb      []int
	Dusm             []int
}

func (a *BasicperfArr) SetData(b Basicperf) {
	a.Ontunetime = append(a.Ontunetime, b.Ontunetime)
	a.Agenttime = append(a.Agenttime, b.Agenttime)
	a.Agentid = append(a.Agentid, b.Agentid)
	a.User = append(a.User, b.User)
	a.Sys = append(a.Sys, b.Sys)
	a.Wait = append(a.Wait, b.Wait)
	a.Idle = append(a.Idle, b.Idle)
	a.Processorcount = append(a.Processorcount, b.Processorcount)
	a.Runqueue = append(a.Runqueue, b.Runqueue)
	a.Blockqueue = append(a.Blockqueue, b.Blockqueue)
	a.Waitqueue = append(a.Waitqueue, b.Waitqueue)
	a.Pqueue = append(a.Pqueue, b.Pqueue)
	a.Pcrateuser = append(a.Pcrateuser, b.Pcrateuser)
	a.Pcratesys = append(a.Pcratesys, b.Pcratesys)
	a.Memorysize = append(a.Memorysize, b.Memorysize)
	a.Memoryused = append(a.Memoryused, b.Memoryused)
	a.Memorypinned = append(a.Memorypinned, b.Memorypinned)
	a.Memorysys = append(a.Memorysys, b.Memorysys)
	a.Memoryuser = append(a.Memoryuser, b.Memoryuser)
	a.Memorycache = append(a.Memorycache, b.Memorycache)
	a.Avm = append(a.Avm, b.Avm)
	a.Pagingspacein = append(a.Pagingspacein, b.Pagingspacein)
	a.Pagingspaceout = append(a.Pagingspaceout, b.Pagingspaceout)
	a.Filesystemin = append(a.Filesystemin, b.Filesystemin)
	a.Filesystemout = append(a.Filesystemout, b.Filesystemout)
	a.Memoryscan = append(a.Memoryscan, b.Memoryscan)
	a.Memoryfreed = append(a.Memoryfreed, b.Memoryfreed)
	a.Swapsize = append(a.Swapsize, b.Swapsize)
	a.Swapused = append(a.Swapused, b.Swapused)
	a.Swapactive = append(a.Swapactive, b.Swapactive)
	a.Fork = append(a.Fork, b.Fork)
	a.Exec = append(a.Exec, b.Exec)
	a.Interupt = append(a.Interupt, b.Interupt)
	a.Systemcall = append(a.Systemcall, b.Systemcall)
	a.Contextswitch = append(a.Contextswitch, b.Contextswitch)
	a.Semaphore = append(a.Semaphore, b.Semaphore)
	a.Msg = append(a.Msg, b.Msg)
	a.Diskreadwrite = append(a.Diskreadwrite, b.Diskreadwrite)
	a.Diskiops = append(a.Diskiops, b.Diskiops)
	a.Networkreadwrite = append(a.Networkreadwrite, b.Networkreadwrite)
	a.Networkiops = append(a.Networkiops, b.Networkiops)
	a.Topcommandid = append(a.Topcommandid, b.Topcommandid)
	a.Topcommandcount = append(a.Topcommandcount, b.Topcommandcount)
	a.Topuserid = append(a.Topuserid, b.Topuserid)
	a.Topcpu = append(a.Topcpu, b.Topcpu)
	a.Topdiskid = append(a.Topdiskid, b.Topdiskid)
	a.Topvgid = append(a.Topvgid, b.Topvgid)
	a.Topbusy = append(a.Topbusy, b.Topbusy)
	a.Maxpid = append(a.Maxpid, b.Maxpid)
	a.Threadcount = append(a.Threadcount, b.Threadcount)
	a.Pidcount = append(a.Pidcount, b.Pidcount)
	a.Linuxbuffer = append(a.Linuxbuffer, b.Linuxbuffer)
	a.Linuxcached = append(a.Linuxcached, b.Linuxcached)
	a.Linuxsrec = append(a.Linuxsrec, b.Linuxsrec)
	a.Memused_mb = append(a.Memused_mb, b.Memused_mb)
	a.Irq = append(a.Irq, b.Irq)
	a.Softirq = append(a.Softirq, b.Softirq)
	a.Swapused_mb = append(a.Swapused_mb, b.Swapused_mb)
	a.Dusm = append(a.Dusm, b.Dusm)
}

func (a *BasicperfArr) GetArgs(dbtype string) []interface{} {
	data := make([]interface{}, 0)

	if dbtype == "pg" {
		data = append(data, pq.Array(a.Ontunetime))
	} else if dbtype == "ts" {
		var ontunetimets []time.Time = make([]time.Time, 0)
		for _, ot := range a.Ontunetime {
			ontunetimets = append(ontunetimets, time.Unix(ot, 0))
		}
		data = append(data, pq.Array(ontunetimets))
	}
	data = append(data, pq.Array(a.Agenttime))
	data = append(data, pq.Array(a.Agentid))
	data = append(data, pq.Array(a.User))
	data = append(data, pq.Array(a.Sys))
	data = append(data, pq.Array(a.Wait))
	data = append(data, pq.Array(a.Idle))
	data = append(data, pq.Array(a.Processorcount))
	data = append(data, pq.Array(a.Runqueue))
	data = append(data, pq.Array(a.Blockqueue))
	data = append(data, pq.Array(a.Waitqueue))
	data = append(data, pq.Array(a.Pqueue))
	data = append(data, pq.Array(a.Pcrateuser))
	data = append(data, pq.Array(a.Pcratesys))
	data = append(data, pq.Array(a.Memorysize))
	data = append(data, pq.Array(a.Memoryused))
	data = append(data, pq.Array(a.Memorypinned))
	data = append(data, pq.Array(a.Memorysys))
	data = append(data, pq.Array(a.Memoryuser))
	data = append(data, pq.Array(a.Memorycache))
	data = append(data, pq.Array(a.Avm))
	data = append(data, pq.Array(a.Pagingspacein))
	data = append(data, pq.Array(a.Pagingspaceout))
	data = append(data, pq.Array(a.Filesystemin))
	data = append(data, pq.Array(a.Filesystemout))
	data = append(data, pq.Array(a.Memoryscan))
	data = append(data, pq.Array(a.Memoryfreed))
	data = append(data, pq.Array(a.Swapsize))
	data = append(data, pq.Array(a.Swapused))
	data = append(data, pq.Array(a.Swapactive))
	data = append(data, pq.Array(a.Fork))
	data = append(data, pq.Array(a.Exec))
	data = append(data, pq.Array(a.Interupt))
	data = append(data, pq.Array(a.Systemcall))
	data = append(data, pq.Array(a.Contextswitch))
	data = append(data, pq.Array(a.Semaphore))
	data = append(data, pq.Array(a.Msg))
	data = append(data, pq.Array(a.Diskreadwrite))
	data = append(data, pq.Array(a.Diskiops))
	data = append(data, pq.Array(a.Networkreadwrite))
	data = append(data, pq.Array(a.Networkiops))
	data = append(data, pq.Array(a.Topcommandid))
	data = append(data, pq.Array(a.Topcommandcount))
	data = append(data, pq.Array(a.Topuserid))
	data = append(data, pq.Array(a.Topcpu))
	data = append(data, pq.Array(a.Topdiskid))
	data = append(data, pq.Array(a.Topvgid))
	data = append(data, pq.Array(a.Topbusy))
	data = append(data, pq.Array(a.Maxpid))
	data = append(data, pq.Array(a.Threadcount))
	data = append(data, pq.Array(a.Pidcount))
	data = append(data, pq.Array(a.Linuxbuffer))
	data = append(data, pq.Array(a.Linuxcached))
	data = append(data, pq.Array(a.Linuxsrec))
	data = append(data, pq.Array(a.Memused_mb))
	data = append(data, pq.Array(a.Irq))
	data = append(data, pq.Array(a.Softirq))
	data = append(data, pq.Array(a.Swapused_mb))
	data = append(data, pq.Array(a.Dusm))

	return data
}

type ProcperffArr struct {
	Ontunetime    []int64
	Agentid       []int
	Hostname      []string
	User          []int
	Sys           []int
	Wait          []int
	Idle          []int
	Memoryused    []int
	Filecache     []int
	Memorysize    []int
	Avm           []int
	Swapused      []int
	Swapsize      []int
	Diskiorate    []int
	Networkiorate []int
	Topproc       []string
	Topuser       []string
	Topproccount  []int
	Topcpu        []int
	Topdisk       []string
	Topvg         []string
	Topbusy       []int
	Maxcpu        []int
	Maxmem        []int
	Maxswap       []int
	Maxdisk       []int
	Diskiops      []int
	Networkiops   []int
}

func (a *ProcperffArr) SetData(p Procperf) {
	a.Ontunetime = append(a.Ontunetime, p.Ontunetime)
	a.Agentid = append(a.Agentid, p.Agentid)
	a.Hostname = append(a.Hostname, p.Hostname)
	a.User = append(a.User, p.User)
	a.Sys = append(a.Sys, p.Sys)
	a.Wait = append(a.Wait, p.Wait)
	a.Idle = append(a.Idle, p.Idle)
	a.Memoryused = append(a.Memoryused, p.Memoryused)
	a.Filecache = append(a.Filecache, p.Filecache)
	a.Memorysize = append(a.Memorysize, p.Memorysize)
	a.Avm = append(a.Avm, p.Avm)
	a.Swapused = append(a.Swapused, p.Swapused)
	a.Swapsize = append(a.Swapsize, p.Swapsize)
	a.Diskiorate = append(a.Diskiorate, p.Diskiorate)
	a.Networkiorate = append(a.Networkiorate, p.Networkiorate)
	a.Topproc = append(a.Topproc, p.Topproc)
	a.Topuser = append(a.Topuser, p.Topuser)
	a.Topproccount = append(a.Topproccount, p.Topproccount)
	a.Topcpu = append(a.Topcpu, p.Topcpu)
	a.Topdisk = append(a.Topdisk, p.Topdisk)
	a.Topvg = append(a.Topvg, p.Topvg)
	a.Topbusy = append(a.Topbusy, p.Topbusy)
	a.Maxcpu = append(a.Maxcpu, p.Maxcpu)
	a.Maxmem = append(a.Maxmem, p.Maxmem)
	a.Maxswap = append(a.Maxswap, p.Maxswap)
	a.Maxdisk = append(a.Maxdisk, p.Maxdisk)
	a.Diskiops = append(a.Diskiops, p.Diskiops)
	a.Networkiops = append(a.Networkiops, p.Networkiops)
}

func (a *ProcperffArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(a.Ontunetime))
	data = append(data, pq.Array(a.Agentid))
	data = append(data, pq.StringArray(a.Hostname))
	data = append(data, pq.Array(a.User))
	data = append(data, pq.Array(a.Sys))
	data = append(data, pq.Array(a.Wait))
	data = append(data, pq.Array(a.Idle))
	data = append(data, pq.Array(a.Memoryused))
	data = append(data, pq.Array(a.Filecache))
	data = append(data, pq.Array(a.Memorysize))
	data = append(data, pq.Array(a.Avm))
	data = append(data, pq.Array(a.Swapused))
	data = append(data, pq.Array(a.Swapsize))
	data = append(data, pq.Array(a.Diskiorate))
	data = append(data, pq.Array(a.Networkiorate))
	data = append(data, pq.StringArray(a.Topproc))
	data = append(data, pq.StringArray(a.Topuser))
	data = append(data, pq.Array(a.Topproccount))
	data = append(data, pq.Array(a.Topcpu))
	data = append(data, pq.StringArray(a.Topdisk))
	data = append(data, pq.StringArray(a.Topvg))
	data = append(data, pq.Array(a.Topbusy))
	data = append(data, pq.Array(a.Maxcpu))
	data = append(data, pq.Array(a.Maxmem))
	data = append(data, pq.Array(a.Maxswap))
	data = append(data, pq.Array(a.Maxdisk))
	data = append(data, pq.Array(a.Diskiops))
	data = append(data, pq.Array(a.Networkiops))

	return data
}

type IoperfArr struct {
	Ontunetime    []int64
	Agentid       []int
	Hostname      []string
	User          []int
	Sys           []int
	Wait          []int
	Idle          []int
	Memoryused    []int
	Filecache     []int
	Memorysize    []int
	Avm           []int
	Swapused      []int
	Swapsize      []int
	Diskiorate    []int
	Networkiorate []int
	Topproc       []string
	Topuser       []string
	Topproccount  []int
	Topcpu        []int
	Topdisk       []string
	Topvg         []string
	Topbusy       []int
	Maxcpu        []int
	Maxmem        []int
	Maxswap       []int
	Maxdisk       []int
	Diskiops      []int
	Networkiops   []int
}

func (a *IoperfArr) SetData(i Ioperf) {
	a.Ontunetime = append(a.Ontunetime, i.Ontunetime)
	a.Agentid = append(a.Agentid, i.Agentid)
	a.Hostname = append(a.Hostname, i.Hostname)
	a.User = append(a.User, i.User)
	a.Sys = append(a.Sys, i.Sys)
	a.Wait = append(a.Wait, i.Wait)
	a.Idle = append(a.Idle, i.Idle)
	a.Memoryused = append(a.Memoryused, i.Memoryused)
	a.Filecache = append(a.Filecache, i.Filecache)
	a.Memorysize = append(a.Memorysize, i.Memorysize)
	a.Avm = append(a.Avm, i.Avm)
	a.Swapused = append(a.Swapused, i.Swapused)
	a.Swapsize = append(a.Swapsize, i.Swapsize)
	a.Diskiorate = append(a.Diskiorate, i.Diskiorate)
	a.Networkiorate = append(a.Networkiorate, i.Networkiorate)
	a.Topproc = append(a.Topproc, i.Topproc)
	a.Topuser = append(a.Topuser, i.Topuser)
	a.Topproccount = append(a.Topproccount, i.Topproccount)
	a.Topcpu = append(a.Topcpu, i.Topcpu)
	a.Topdisk = append(a.Topdisk, i.Topdisk)
	a.Topvg = append(a.Topvg, i.Topvg)
	a.Topbusy = append(a.Topbusy, i.Topbusy)
	a.Maxcpu = append(a.Maxcpu, i.Maxcpu)
	a.Maxmem = append(a.Maxmem, i.Maxmem)
	a.Maxswap = append(a.Maxswap, i.Maxswap)
	a.Maxdisk = append(a.Maxdisk, i.Maxdisk)
	a.Diskiops = append(a.Diskiops, i.Diskiops)
	a.Networkiops = append(a.Networkiops, i.Networkiops)
}

func (a *IoperfArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(a.Ontunetime))
	data = append(data, pq.Array(a.Agentid))
	data = append(data, pq.StringArray(a.Hostname))
	data = append(data, pq.Array(a.User))
	data = append(data, pq.Array(a.Sys))
	data = append(data, pq.Array(a.Wait))
	data = append(data, pq.Array(a.Idle))
	data = append(data, pq.Array(a.Memoryused))
	data = append(data, pq.Array(a.Filecache))
	data = append(data, pq.Array(a.Memorysize))
	data = append(data, pq.Array(a.Avm))
	data = append(data, pq.Array(a.Swapused))
	data = append(data, pq.Array(a.Swapsize))
	data = append(data, pq.Array(a.Diskiorate))
	data = append(data, pq.Array(a.Networkiorate))
	data = append(data, pq.StringArray(a.Topproc))
	data = append(data, pq.StringArray(a.Topuser))
	data = append(data, pq.Array(a.Topproccount))
	data = append(data, pq.Array(a.Topcpu))
	data = append(data, pq.StringArray(a.Topdisk))
	data = append(data, pq.StringArray(a.Topvg))
	data = append(data, pq.Array(a.Topbusy))
	data = append(data, pq.Array(a.Maxcpu))
	data = append(data, pq.Array(a.Maxmem))
	data = append(data, pq.Array(a.Maxswap))
	data = append(data, pq.Array(a.Maxdisk))
	data = append(data, pq.Array(a.Diskiops))
	data = append(data, pq.Array(a.Networkiops))

	return data
}

type CpuperfArr struct {
	Ontunetime    []int64
	Agentid       []int
	Hostname      []string
	User          []int
	Sys           []int
	Wait          []int
	Idle          []int
	Memoryused    []int
	Filecache     []int
	Memorysize    []int
	Avm           []int
	Swapused      []int
	Swapsize      []int
	Diskiorate    []int
	Networkiorate []int
	Topproc       []string
	Topuser       []string
	Topproccount  []int
	Topcpu        []int
	Topdisk       []string
	Topvg         []string
	Topbusy       []int
	Maxcpu        []int
	Maxmem        []int
	Maxswap       []int
	Maxdisk       []int
	Diskiops      []int
	Networkiops   []int
}

func (a *CpuperfArr) SetData(c Cpuperf) {
	a.Ontunetime = append(a.Ontunetime, c.Ontunetime)
	a.Agentid = append(a.Agentid, c.Agentid)
	a.Hostname = append(a.Hostname, c.Hostname)
	a.User = append(a.User, c.User)
	a.Sys = append(a.Sys, c.Sys)
	a.Wait = append(a.Wait, c.Wait)
	a.Idle = append(a.Idle, c.Idle)
	a.Memoryused = append(a.Memoryused, c.Memoryused)
	a.Filecache = append(a.Filecache, c.Filecache)
	a.Memorysize = append(a.Memorysize, c.Memorysize)
	a.Avm = append(a.Avm, c.Avm)
	a.Swapused = append(a.Swapused, c.Swapused)
	a.Swapsize = append(a.Swapsize, c.Swapsize)
	a.Diskiorate = append(a.Diskiorate, c.Diskiorate)
	a.Networkiorate = append(a.Networkiorate, c.Networkiorate)
	a.Topproc = append(a.Topproc, c.Topproc)
	a.Topuser = append(a.Topuser, c.Topuser)
	a.Topproccount = append(a.Topproccount, c.Topproccount)
	a.Topcpu = append(a.Topcpu, c.Topcpu)
	a.Topdisk = append(a.Topdisk, c.Topdisk)
	a.Topvg = append(a.Topvg, c.Topvg)
	a.Topbusy = append(a.Topbusy, c.Topbusy)
	a.Maxcpu = append(a.Maxcpu, c.Maxcpu)
	a.Maxmem = append(a.Maxmem, c.Maxmem)
	a.Maxswap = append(a.Maxswap, c.Maxswap)
	a.Maxdisk = append(a.Maxdisk, c.Maxdisk)
	a.Diskiops = append(a.Diskiops, c.Diskiops)
	a.Networkiops = append(a.Networkiops, c.Networkiops)
}

func (a *CpuperfArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(a.Ontunetime))
	data = append(data, pq.Array(a.Agentid))
	data = append(data, pq.StringArray(a.Hostname))
	data = append(data, pq.Array(a.User))
	data = append(data, pq.Array(a.Sys))
	data = append(data, pq.Array(a.Wait))
	data = append(data, pq.Array(a.Idle))
	data = append(data, pq.Array(a.Memoryused))
	data = append(data, pq.Array(a.Filecache))
	data = append(data, pq.Array(a.Memorysize))
	data = append(data, pq.Array(a.Avm))
	data = append(data, pq.Array(a.Swapused))
	data = append(data, pq.Array(a.Swapsize))
	data = append(data, pq.Array(a.Diskiorate))
	data = append(data, pq.Array(a.Networkiorate))
	data = append(data, pq.StringArray(a.Topproc))
	data = append(data, pq.StringArray(a.Topuser))
	data = append(data, pq.Array(a.Topproccount))
	data = append(data, pq.Array(a.Topcpu))
	data = append(data, pq.StringArray(a.Topdisk))
	data = append(data, pq.StringArray(a.Topvg))
	data = append(data, pq.Array(a.Topbusy))
	data = append(data, pq.Array(a.Maxcpu))
	data = append(data, pq.Array(a.Maxmem))
	data = append(data, pq.Array(a.Maxswap))
	data = append(data, pq.Array(a.Maxdisk))
	data = append(data, pq.Array(a.Diskiops))
	data = append(data, pq.Array(a.Networkiops))

	return data
}

type LastrealtimeperfArr struct {
	Ontunetime    []int64
	Agentid       []int
	Hostname      []string
	User          []int
	Sys           []int
	Wait          []int
	Idle          []int
	Memoryused    []int
	Filecache     []int
	Memorysize    []int
	Avm           []int
	Swapused      []int
	Swapsize      []int
	Diskiorate    []int
	Networkiorate []int
	Topproc       []string
	Topuser       []string
	Topproccount  []int
	Topcpu        []int
	Topdisk       []string
	Topvg         []string
	Topbusy       []int
	Maxcpu        []int
	Maxmem        []int
	Maxswap       []int
	Maxdisk       []int
	Diskiops      []int
	Networkiops   []int
}

func (a *LastrealtimeperfArr) SetData(l Lastrealtimeperf) {
	a.Ontunetime = append(a.Ontunetime, l.Ontunetime)
	a.Agentid = append(a.Agentid, l.Agentid)
	a.Hostname = append(a.Hostname, l.Hostname)
	a.User = append(a.User, l.User)
	a.Sys = append(a.Sys, l.Sys)
	a.Wait = append(a.Wait, l.Wait)
	a.Idle = append(a.Idle, l.Idle)
	a.Memoryused = append(a.Memoryused, l.Memoryused)
	a.Filecache = append(a.Filecache, l.Filecache)
	a.Memorysize = append(a.Memorysize, l.Memorysize)
	a.Avm = append(a.Avm, l.Avm)
	a.Swapused = append(a.Swapused, l.Swapused)
	a.Swapsize = append(a.Swapsize, l.Swapsize)
	a.Diskiorate = append(a.Diskiorate, l.Diskiorate)
	a.Networkiorate = append(a.Networkiorate, l.Networkiorate)
	a.Topproc = append(a.Topproc, l.Topproc)
	a.Topuser = append(a.Topuser, l.Topuser)
	a.Topproccount = append(a.Topproccount, l.Topproccount)
	a.Topcpu = append(a.Topcpu, l.Topcpu)
	a.Topdisk = append(a.Topdisk, l.Topdisk)
	a.Topvg = append(a.Topvg, l.Topvg)
	a.Topbusy = append(a.Topbusy, l.Topbusy)
	a.Maxcpu = append(a.Maxcpu, l.Maxcpu)
	a.Maxmem = append(a.Maxmem, l.Maxmem)
	a.Maxswap = append(a.Maxswap, l.Maxswap)
	a.Maxdisk = append(a.Maxdisk, l.Maxdisk)
	a.Diskiops = append(a.Diskiops, l.Diskiops)
	a.Networkiops = append(a.Networkiops, l.Networkiops)
}

func (a *LastrealtimeperfArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(a.Ontunetime))
	data = append(data, pq.Array(a.Agentid))
	data = append(data, pq.StringArray(a.Hostname))
	data = append(data, pq.Array(a.User))
	data = append(data, pq.Array(a.Sys))
	data = append(data, pq.Array(a.Wait))
	data = append(data, pq.Array(a.Idle))
	data = append(data, pq.Array(a.Memoryused))
	data = append(data, pq.Array(a.Filecache))
	data = append(data, pq.Array(a.Memorysize))
	data = append(data, pq.Array(a.Avm))
	data = append(data, pq.Array(a.Swapused))
	data = append(data, pq.Array(a.Swapsize))
	data = append(data, pq.Array(a.Diskiorate))
	data = append(data, pq.Array(a.Networkiorate))
	data = append(data, pq.StringArray(a.Topproc))
	data = append(data, pq.StringArray(a.Topuser))
	data = append(data, pq.Array(a.Topproccount))
	data = append(data, pq.Array(a.Topcpu))
	data = append(data, pq.StringArray(a.Topdisk))
	data = append(data, pq.StringArray(a.Topvg))
	data = append(data, pq.Array(a.Topbusy))
	data = append(data, pq.Array(a.Maxcpu))
	data = append(data, pq.Array(a.Maxmem))
	data = append(data, pq.Array(a.Maxswap))
	data = append(data, pq.Array(a.Maxdisk))
	data = append(data, pq.Array(a.Diskiops))
	data = append(data, pq.Array(a.Networkiops))

	return data
}

type LastperfArr struct {
	Ontunetime       []int64
	Hostname         []string
	User             []int
	Sys              []int
	Wait             []int
	Idle             []int
	Avm              []int
	Memoryused       []int
	Filecache        []int
	Swapused         []int
	Diskiorate       []int
	Networkiorate    []int
	Topproc          []string
	Topcpu           []int
	Topdisk          []string
	Topbusy          []int
	Filesystem       []string
	Runqueue         []int
	Blockqueue       []int
	Pagingspacein    []int
	Pagingspaceout   []int
	Filesystemin     []int
	Filesystemout    []int
	Memoryscan       []int
	Memoryfreed      []int
	Swapactive       []int
	Fork             []int
	Exec             []int
	Interupt         []int
	Systemcall       []int
	Contextswitch    []int
	Semaphore        []int
	Msg              []int
	Diskiops         []int
	Networkiops      []int
	Networkreadrate  []int
	Networkwriterate []int
}

func (a *LastperfArr) SetData(l Lastperf) {
	a.Ontunetime = append(a.Ontunetime, l.Ontunetime)
	a.Hostname = append(a.Hostname, l.Hostname)
	a.User = append(a.User, l.User)
	a.Sys = append(a.Sys, l.Sys)
	a.Wait = append(a.Wait, l.Wait)
	a.Idle = append(a.Idle, l.Idle)
	a.Avm = append(a.Avm, l.Avm)
	a.Memoryused = append(a.Memoryused, l.Memoryused)
	a.Filecache = append(a.Filecache, l.Filecache)
	a.Swapused = append(a.Swapused, l.Swapused)
	a.Diskiorate = append(a.Diskiorate, l.Diskiorate)
	a.Networkiorate = append(a.Networkiorate, l.Networkiorate)
	a.Topproc = append(a.Topproc, l.Topproc)
	a.Topcpu = append(a.Topcpu, l.Topcpu)
	a.Topdisk = append(a.Topdisk, l.Topdisk)
	a.Topbusy = append(a.Topbusy, l.Topbusy)
	a.Filesystem = append(a.Filesystem, l.Filesystem)
	a.Runqueue = append(a.Runqueue, l.Runqueue)
	a.Blockqueue = append(a.Blockqueue, l.Blockqueue)
	a.Pagingspacein = append(a.Pagingspacein, l.Pagingspacein)
	a.Pagingspaceout = append(a.Pagingspaceout, l.Pagingspaceout)
	a.Filesystemin = append(a.Filesystemin, l.Filesystemin)
	a.Filesystemout = append(a.Filesystemout, l.Filesystemout)
	a.Memoryscan = append(a.Memoryscan, l.Memoryscan)
	a.Memoryfreed = append(a.Memoryfreed, l.Memoryfreed)
	a.Swapactive = append(a.Swapactive, l.Swapactive)
	a.Fork = append(a.Fork, l.Fork)
	a.Exec = append(a.Exec, l.Exec)
	a.Interupt = append(a.Interupt, l.Interupt)
	a.Systemcall = append(a.Systemcall, l.Systemcall)
	a.Contextswitch = append(a.Contextswitch, l.Contextswitch)
	a.Semaphore = append(a.Semaphore, l.Semaphore)
	a.Msg = append(a.Msg, l.Msg)
	a.Diskiops = append(a.Diskiops, l.Diskiops)
	a.Networkiops = append(a.Networkiops, l.Networkiops)
	a.Networkreadrate = append(a.Networkreadrate, l.Networkreadrate)
	a.Networkwriterate = append(a.Networkwriterate, l.Networkwriterate)
}

func (a *LastperfArr) GetArgs() []interface{} {
	data := make([]interface{}, 0)
	data = append(data, pq.Array(a.Ontunetime))
	data = append(data, pq.StringArray(a.Hostname))
	data = append(data, pq.Array(a.User))
	data = append(data, pq.Array(a.Sys))
	data = append(data, pq.Array(a.Wait))
	data = append(data, pq.Array(a.Idle))
	data = append(data, pq.Array(a.Avm))
	data = append(data, pq.Array(a.Memoryused))
	data = append(data, pq.Array(a.Filecache))
	data = append(data, pq.Array(a.Swapused))
	data = append(data, pq.Array(a.Diskiorate))
	data = append(data, pq.Array(a.Networkiorate))
	data = append(data, pq.StringArray(a.Topproc))
	data = append(data, pq.Array(a.Topcpu))
	data = append(data, pq.StringArray(a.Topdisk))
	data = append(data, pq.Array(a.Topbusy))
	data = append(data, pq.StringArray(a.Filesystem))
	data = append(data, pq.Array(a.Runqueue))
	data = append(data, pq.Array(a.Blockqueue))
	data = append(data, pq.Array(a.Pagingspacein))
	data = append(data, pq.Array(a.Pagingspaceout))
	data = append(data, pq.Array(a.Filesystemin))
	data = append(data, pq.Array(a.Filesystemout))
	data = append(data, pq.Array(a.Memoryscan))
	data = append(data, pq.Array(a.Memoryfreed))
	data = append(data, pq.Array(a.Swapactive))
	data = append(data, pq.Array(a.Fork))
	data = append(data, pq.Array(a.Exec))
	data = append(data, pq.Array(a.Interupt))
	data = append(data, pq.Array(a.Systemcall))
	data = append(data, pq.Array(a.Contextswitch))
	data = append(data, pq.Array(a.Semaphore))
	data = append(data, pq.Array(a.Msg))
	data = append(data, pq.Array(a.Diskiops))
	data = append(data, pq.Array(a.Networkiops))
	data = append(data, pq.Array(a.Networkreadrate))
	data = append(data, pq.Array(a.Networkwriterate))

	return data
}
