package log

import (
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

func SetLevel(level Level) *Logger {
	return std.SetLevel(level)
}

func Sync() {
	std.Sync()
}

func Clone() *Logger {
	return std.Clone()
}

func SetCallerDepth(callerDepth int) *Logger {
	return std.SetCallerDepth(callerDepth)
}

func SetPrefixMsg(prefixMsg string) *Logger {
	return std.SetPrefixMsg(prefixMsg)
}

func SetSuffixMsg(suffixMsg string) *Logger {
	return std.SetSuffixMsg(suffixMsg)
}
