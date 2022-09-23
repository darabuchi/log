package log

import (
	"sync"

	"github.com/nats-io/nuid"
	"github.com/petermattis/goid"
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
		traceId = GenTraceId()
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

var nid = nuid.New()

func GenTraceId() string {
	return nid.Next()[16:]
}
