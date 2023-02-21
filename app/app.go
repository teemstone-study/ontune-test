package app

import (
	"log"
	"mgr/data"
	"os"
	"sync"
	"time"
)

const (
	DEBUG_FLAG      = false
	TIME_DEBUG_FLAG = false
	LASTPERF_TIMER  = 60
	PROCESS_TIMER   = 1
	MAX_THREAD      = 8
	DATE_FORMAT     = "06010215"
	CPU_CORE        = 4
)

var (
	MapHostInfo        map[int]string = make(map[int]string)
	GlobalOntunetime   int64          = time.Now().Unix()
	GlobalOntunetimets time.Time      = time.Now()
	AgentDataProcess   *sync.Map      = &sync.Map{}
	DISK_IONAME        map[int]string = map[int]string{
		0: "total",
		1: "C:",
		2: "D:",
	}
	NET_IONAME map[int]string = map[int]string{
		0: "total",
		3: "Intel[R] Ethernet Connection [7] I219-LM",
		4: "Teredo Tunneling Pseudo-Interface",
	}

	DEVICE_IDS    []string = []string{"total", "C:", "D:", "Intel[R] Ethernet Connection [7] I219-LM", "Teredo Tunneling Pseudo-Interface"}
	VOLUME_GROUPS []string = []string{
		"NULL", "", "rootvg", "None", "N/A", "vg_linux63x8664", "caavg_private", "ProLinux-vg", "centos", "vg_centos62x8664",
		"VolGroup00", "vg_centos67", "vg_oraclelinux69", "ubuntu64", "VG_XenStorage-628ef03e-1cf7-cd4", "XSLocalEXT-a4fc36e8-ca45-6f11-c", "vg_centos6", "vg00",
		"prolinux", "gpfs1", "gpfs2", "freedisk", "centos_centos7kvm", "centos_numatest", "debian8-vg", "vg_data",
		"ontubun1604-vg", "2a-ce75126bbade", "rhel", "centos_k8snode1", "centos_master", "zabbix-vg", "cl_centos8", "old_rootvg",
		"DSS0", "dg1", "dg2", "XSLocalEXT-86d07858-845c-9433-3", "prchfpdg1", "prchfpdg2", "cl_justin", "d1home",
		"d2home", "testvg02", "testvg03", "testvg04", "testvg05", "testvg06", "testvg07", "testvg01",
	}
)

type Bitmask uint32

func (value Bitmask) IsSet(key Bitmask) bool {
	return value&key != 0
}

type ChannelStruct struct {
	DemoBasicData      chan []*data.Basicperf
	DemoAvgBasicData   chan []*data.Basicperf
	DemoCpuData        chan []*data.Cpuperf
	DemoAvgCpuData     chan []*data.Cpuperf
	DemoDiskData       chan []*data.Diskperf
	DemoAvgDiskData    chan []*data.Diskperf
	DemoNetData        chan []*data.Netperf
	DemoAvgNetData     chan []*data.Netperf
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
		DemoBasicData:      make(chan []*data.Basicperf),
		DemoAvgBasicData:   make(chan []*data.Basicperf),
		DemoCpuData:        make(chan []*data.Cpuperf),
		DemoAvgCpuData:     make(chan []*data.Cpuperf),
		DemoDiskData:       make(chan []*data.Diskperf),
		DemoAvgDiskData:    make(chan []*data.Diskperf),
		DemoNetData:        make(chan []*data.Netperf),
		DemoAvgNetData:     make(chan []*data.Netperf),
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
	switch log_type {
	case "log":
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
