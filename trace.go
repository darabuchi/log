package log

import (
	"github.com/aofei/sandid"
	"github.com/petermattis/goid"
	"sync"
)

var traceMap sync.Map

func getTrace(gid int64) string {
	tid, ok := traceMap.Load(gid)
	if !ok {
		return ""
	}
	return tid.(string)
}

func setTrace(gid int64, traceId string) {
	if traceId == "" {
		traceId = sandid.New().String()
	}
	traceMap.Store(gid, traceId)
}

func delTrace(gid int64) {
	traceMap.Delete(gid)
}

func GetTrace() string {
	return getTrace(goid.Get())
}

func SetTrace(traceId string) {
	setTrace(goid.Get(), traceId)
}

func DelTrace() {
	delTrace(goid.Get())
}