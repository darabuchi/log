package log

import (
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

func SetOutput(writes ...io.Writer) *Logger {
	return std.SetOutput(writes...)
}

func AddOutput(writes ...io.Writer) *Logger {
	return std.AddOutput(writes...)
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
