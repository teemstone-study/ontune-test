package app

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"mgr/data"
	"os"
	"sync"
	"time"
)

const (
	DEBUG_FLAG         = false
	TIME_DEBUG_FLAG    = false
	LASTPERF_TIMER     = 60
	PROCESS_TIMER      = 1
	MAX_THREAD         = 8
	HOURLY_DATE_FORMAT = "06010215"
	DAILY_DATE_FORMAT  = "06010200"
	CPU_CORE           = 4
	DF_COUNT           = 4
)

var (
	MapHostInfo        map[int]string = make(map[int]string)
	GlobalOntunetime   int64          = time.Now().Unix()
	GlobalOntunetimets time.Time      = time.Now()
	AgentDataProcess   *sync.Map      = &sync.Map{}
	PROCCMD_ARR        []int          = []int{rand.Intn(len(data.PROCCMD_DATA)), rand.Intn(len(data.PROCCMD_DATA)), rand.Intn(len(data.PROCCMD_DATA))}
	PROCUSER_ARR       []int          = []int{rand.Intn(len(data.PROCCMD_DATA))}
	PROCARG_ARR        []int          = []int{rand.Intn(len(data.PROCCMD_DATA)), rand.Intn(len(data.PROCCMD_DATA)), rand.Intn(len(data.PROCCMD_DATA))}
)

type Bitmask uint32

func (value Bitmask) IsSet(key Bitmask) bool {
	return value&key != 0
}

type ChannelStruct struct {
	DemoPerfData       chan []*data.Perf
	DemoAvgPerfData    chan []*data.Perf
	DemoAvgMaxPerfData chan []*data.Perf
	DemoCpuData        chan []*data.Cpu
	DemoAvgCpuData     chan []*data.Cpu
	DemoAvgMaxCpuData  chan []*data.Cpu
	DemoDiskData       chan []*data.Disk
	DemoAvgDiskData    chan []*data.Disk
	DemoAvgMaxDiskData chan []*data.Disk
	DemoNetData        chan []*data.Net
	DemoAvgNetData     chan []*data.Net
	DemoAvgMaxNetData  chan []*data.Net
	DemoProcData       chan []*data.Pid
	DemoAvgProcData    chan []*data.Pid
	DemoAvgMaxProcData chan []*data.Pid
	DemoAvgDfData      chan []*data.Df
	AverageRequest     chan string
	AgentData          chan *sync.Map
	LastPerfData       chan *sync.Map
	AgentCopyDone      chan bool
	LastperfCopyDone   chan bool
	AgentInsertDone    chan bool
	LastperfInsertDone chan bool
}

var (
	GlobalChannel ChannelStruct = ChannelStruct{
		DemoPerfData:       make(chan []*data.Perf),
		DemoAvgPerfData:    make(chan []*data.Perf),
		DemoAvgMaxPerfData: make(chan []*data.Perf),
		DemoCpuData:        make(chan []*data.Cpu),
		DemoAvgCpuData:     make(chan []*data.Cpu),
		DemoAvgMaxCpuData:  make(chan []*data.Cpu),
		DemoDiskData:       make(chan []*data.Disk),
		DemoAvgDiskData:    make(chan []*data.Disk),
		DemoAvgMaxDiskData: make(chan []*data.Disk),
		DemoNetData:        make(chan []*data.Net),
		DemoAvgNetData:     make(chan []*data.Net),
		DemoAvgMaxNetData:  make(chan []*data.Net),
		DemoProcData:       make(chan []*data.Pid),
		DemoAvgProcData:    make(chan []*data.Pid),
		DemoAvgMaxProcData: make(chan []*data.Pid),
		DemoAvgDfData:      make(chan []*data.Df),
		AverageRequest:     make(chan string),
		AgentData:          make(chan *sync.Map),
		LastPerfData:       make(chan *sync.Map),
		AgentCopyDone:      make(chan bool),
		LastperfCopyDone:   make(chan bool),
		AgentInsertDone:    make(chan bool),
		LastperfInsertDone: make(chan bool),
	}
)

func LogWrite(log_type string, data string) {
	var file *os.File
	var err error

	if !DEBUG_FLAG && log_type == "debug" {
		return
	}

	switch log_type {
	case "log":
		file, err = os.OpenFile("ontuneLog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	case "debug":
		file, err = os.OpenFile("ontuneLog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	case "error":
		file, err = os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	if err != nil {
		panic(err.Error())
	}
	logger := *log.New(file, "", 0)
	defer file.Close()

	err = logger.Output(1, data)
	ErrorCheck(err)
}

func getOntuneinfo(dbtype string) []interface{} {
	var parameters []interface{} = make([]interface{}, 0)
	parameters = append(parameters, DEFAULT_MAJORVERSION)
	parameters = append(parameters, DEFAULT_MINORVERSION)
	parameters = append(parameters, DEFAULT_RELEASEVERSION)

	if dbtype == "pg" {
		parameters = append(parameters, time.Now().Unix())
	} else if dbtype == "ts" {
		parameters = append(parameters, time.Now())
	}
	parameters = append(parameters, []byte("8778"))
	parameters = append(parameters, []byte("2580"))
	parameters = append(parameters, DEFAULT_FDATA)
	parameters = append(parameters, DEFAULT_STANDAREBIAS)
	parameters = append(parameters, DEFAULT_DAYLIGHT)
	parameters = append(parameters, DEFAULT_DAYLIGHTBIAS)
	parameters = append(parameters, DEFAULT_TABLEMODE)
	parameters = append(parameters, DEFAULT_BIAS)
	parameters = append(parameters, DEFAULT_VALUE)
	parameters = append(parameters, DEFAULT_VALUE)
	parameters = append(parameters, DEFAULT_VALUE)
	parameters = append(parameters, DEFAULT_VALUE)
	parameters = append(parameters, DEFAULT_VALUE)
	parameters = append(parameters, DEFAULT_VALUE)

	return parameters
}

func contains(key any) bool {
	if _, ok := AgentDataProcess.Load(key); ok {
		return true
	} else {
		AgentDataProcess.Store(key, struct{}{})
		//LogWrite("log", fmt.Sprintf("make %s", str))
		return false
	}
}

func GetMapSize(src *sync.Map) int {
	var size int
	src.Range(func(key, value any) bool {
		size += 1
		return true
	})

	return size
}

func GetMapDataSize(src *sync.Map) int {
	var data_size int
	src.Range(func(key, value any) bool {
		val_map := value.(*sync.Map)
		data_size += GetMapSize(val_map)
		return true
	})

	return data_size
}

func GetMapKeys(src *sync.Map) string {
	var keys string
	src.Range(func(key, value any) bool {
		keys = keys + key.(string) + ","
		return true
	})

	return keys + "\n"
}

func CopyAgentMap(src *sync.Map) *sync.Map {
	tgt := sync.Map{}
	src.Range(func(key, value any) bool {
		tgt.Store(key, CopyMap(value.(*sync.Map)))
		return true
	})

	return &tgt
}

func CopyMap(src *sync.Map) *sync.Map {
	tgt := sync.Map{}
	src.Range(func(key, value any) bool {
		tgt.Store(key, value)
		return true
	})
	return &tgt
}

func MakeProc(pid_data []*data.Pid) []*data.Proc {
	proc_map := &sync.Map{}
	proc_data := make([]*data.Proc, 0)

	for _, p := range pid_data {
		key := fmt.Sprintf("%d_%d_%d", p.Agentid, p.Cmdid, p.Userid)
		if pmap, ok := proc_map.LoadOrStore(key, &data.Proc{}); ok {
			pmap_struct := pmap.(*data.Proc)
			pmap_struct.Ontunetime = int64(math.Max(float64(pmap_struct.Ontunetime), float64(p.Ontunetime)))
			pmap_struct.Agenttime = int64(math.Max(float64(pmap_struct.Agenttime), float64(p.Agenttime)))
			pmap_struct.Agentid = p.Agentid
			pmap_struct.Cmdid = p.Cmdid
			pmap_struct.Userid = p.Userid
			pmap_struct.Usr += p.Usr
			pmap_struct.Sys += p.Sys
			pmap_struct.Usrsys += p.Usrsys
			pmap_struct.Sz += p.Sz
			pmap_struct.Rss += p.Rss
			pmap_struct.Vmem += p.Vmem
			pmap_struct.Chario += p.Chario
			pmap_struct.Processcnt += p.Processcnt
			pmap_struct.Threadcnt += p.Threadcnt
			pmap_struct.Pvbytes += p.Pvbytes
			pmap_struct.Pgpool += p.Pgpool
		}
	}

	proc_map.Range(func(key, value any) bool {
		proc_data = append(proc_data, value.(*data.Proc))
		return true
	})

	return proc_data
}
