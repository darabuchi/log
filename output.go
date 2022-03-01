package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
)

func SetOutput(writes ...io.Writer) *Logger {
	return std.SetOutput(writes...)
}

func AddOutput(writes ...io.Writer) *Logger {
	return std.AddOutput(writes...)
}

func GetOutputWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(filename)
	if err != nil {
		std.Panicf("err:%v", err)
	}
	return hook
}
