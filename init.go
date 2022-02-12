package log

import (
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
