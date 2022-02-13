package log

import (
	"github.com/petermattis/goid"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var (
	std = newLogger()

	pid = os.Getpid()
)

func init() {
	//std.SetFormatter(logFmt)
	//std.SetReportCaller(false)
	//std.SetOutput(os.Stdout)
}

func New() *Logger {
	return newLogger()
}

func SetOutput(write io.Writer) {
	std.SetOutput(zapcore.AddSync(write))
}

func SetLevel(level Level) {
	std.SetLevel(level)
}

//func SetModule(module string) {
//	logFmt.SetModule(module)
//}

func AddOutput(write io.Writer) {
	std.AddOutput(zapcore.AddSync(write))
}

func Sync() {
	std.Sync()
}

func Clone() *Logger {
	return std.Clone()
}

func SetCallerDepth(callerDepth int) {
	std.SetCallerDepth(callerDepth)
}

func SetPrefixMsg(prefixMsg string) {
	std.SetPrefixMsg(prefixMsg)
}

func SetSuffixMsg(suffixMsg string) {
	std.SetSuffixMsg(suffixMsg)
}

func GetTrace() string {
	return std.GetTrace(goid.Get())
}

func SetTrace(traceId string) {
	std.SetTrace(goid.Get(), traceId)
}

func DelTrace() {
	std.DelTrace(goid.Get())
}
