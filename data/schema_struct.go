package data

type Agentinfo struct {
	Agentid           int
	Agentname         string
	Enabled           int
	Connected         int
	Updated           int
	Shorttermbasic    int
	Shorttermproc     int
	Shorttermio       int
	Shorttermcpu      int
	Longtermbasic     int
	Longtermproc      int
	Longtermio        int
	Longtermcpu       int
	Model             string
	Serial            string
	Group             string
	Ipaddress         string
	Pscommand         string
	Logevent          string
	Processevent      string
	Timecheck         int
	Disconnectedtime  int64
	Skipdatatypes     int
	Virbasicperf      int
	Hypervisor        int
	Serviceevent      string
	Installdate       int64
	Lastconnectedtime int64
}

type Hostinfo struct {
	Agentid        int
	Hostname       string
	Hostnameext    string
	Os             string
	Fw             string
	Agentversion   string
	Model          string
	Serial         string
	Processorcount int
	Processorclock int
	Memorysize     int
	Swapsize       int
	Poolid         int
	Replication    int
	Smt            int
	Micropar       int
	Capped         int
	Ec             int
	Virtualcpu     int
	Weight         int
	Cpupool        int
	Ams            int
	Allip          string
	Numanodecount  int
}

type Basicperf struct {
	Ontunetime       int64
	Agenttime        int64
	Agentid          int
	User             int
	Sys              int
	Wait             int
	Idle             int
	Processorcount   int
	Runqueue         int
	Blockqueue       int
	Waitqueue        int
	Pqueue           int
	Pcrateuser       int
	Pcratesys        int
	Memorysize       int
	Memoryused       int
	Memorypinned     int
	Memorysys        int
	Memoryuser       int
	Memorycache      int
	Avm              int
	Pagingspacein    int
	Pagingspaceout   int
	Filesystemin     int
	Filesystemout    int
	Memoryscan       int
	Memoryfreed      int
	Swapsize         int
	Swapused         int
	Swapactive       int
	Fork             int
	Exec             int
	Interupt         int
	Systemcall       int
	Contextswitch    int
	Semaphore        int
	Msg              int
	Diskreadwrite    int
	Diskiops         int
	Networkreadwrite int
	Networkiops      int
	Topcommandid     int
	Topcommandcount  int
	Topuserid        int
	Topcpu           int
	Topdiskid        int
	Topvgid          int
	Topbusy          int
	Maxpid           int
	Threadcount      int
	Pidcount         int
	Linuxbuffer      int
	Linuxcached      int
	Linuxsrec        int
	Memused_mb       int
	Irq              int
	Softirq          int
	Swapused_mb      int
	Dusm             int
}

type Procperf struct {
	Ontunetime    int64
	Agentid       int
	Hostname      string
	User          int
	Sys           int
	Wait          int
	Idle          int
	Memoryused    int
	Filecache     int
	Memorysize    int
	Avm           int
	Swapused      int
	Swapsize      int
	Diskiorate    int
	Networkiorate int
	Topproc       string
	Topuser       string
	Topproccount  int
	Topcpu        int
	Topdisk       string
	Topvg         string
	Topbusy       int
	Maxcpu        int
	Maxmem        int
	Maxswap       int
	Maxdisk       int
	Diskiops      int
	Networkiops   int
}

type Diskperf struct {
	Ontunetime   int64
	Agenttime    int64
	Agentid      int
	Ionameid     int
	Readrate     int
	Writerate    int
	Iops         int
	Busy         int
	Descid       int
	Readsvctime  int
	Writesvctime int
}

type Netperf struct {
	Ontunetime int64
	Agenttime  int64
	Agentid    int
	Ionameid   int
	Readrate   int
	Writerate  int
	Readiops   int
	Writeiops  int
	Errorps    int
	Collision  int
}

type Cpuperf struct {
	Ontunetime    int64
	Agenttime     int64
	Agentid       int
	Index         int
	User          int
	Sys           int
	Wait          int
	Idle          int
	Runqueue      int
	Fork          int
	Exec          int
	Interupt      int
	Systemcall    int
	Contextswitch int
}

type Lastrealtimeperf struct {
	Ontunetime    int64
	Agentid       int
	Hostname      string
	User          int
	Sys           int
	Wait          int
	Idle          int
	Memoryused    int
	Filecache     int
	Memorysize    int
	Avm           int
	Swapused      int
	Swapsize      int
	Diskiorate    int
	Networkiorate int
	Topproc       string
	Topuser       string
	Topproccount  int
	Topcpu        int
	Topdisk       string
	Topvg         string
	Topbusy       int
	Maxcpu        int
	Maxmem        int
	Maxswap       int
	Maxdisk       int
	Diskiops      int
	Networkiops   int
}

type Lastperf struct {
	Ontunetime       int64
	Hostname         string
	User             int
	Sys              int
	Wait             int
	Idle             int
	Avm              int
	Memoryused       int
	Filecache        int
	Swapused         int
	Diskiorate       int
	Networkiorate    int
	Topproc          string
	Topcpu           int
	Topdisk          string
	Topbusy          int
	Filesystem       string
	Runqueue         int
	Blockqueue       int
	Pagingspacein    int
	Pagingspaceout   int
	Filesystemin     int
	Filesystemout    int
	Memoryscan       int
	Memoryfreed      int
	Swapactive       int
	Fork             int
	Exec             int
	Interupt         int
	Systemcall       int
	Contextswitch    int
	Semaphore        int
	Msg              int
	Diskiops         int
	Networkiops      int
	Networkreadrate  int
	Networkwriterate int
}
