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
)

var (
	MapHostifo         map[int]string = make(map[int]string)
	GlobalOntunetime   int64          = time.Now().Unix()
	GlobalOntunetimets time.Time      = time.Now()
	AgentDataProcess   *sync.Map      = &sync.Map{}
)

type Bitmask uint32

func (value Bitmask) IsSet(key Bitmask) bool {
	return value&key != 0
}

type ChannelStruct struct {
	DemoBasicData      chan []*data.Basicperf
	DemoAvgBasicData   chan []*data.Basicperf
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

func getMapSize(src *sync.Map) int {
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
