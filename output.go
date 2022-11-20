package log

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
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

func GetOutputWriterHourly(filename string, max uint) io.Writer {
	if max <= 0 {
		max = 24
	}

	hook, err := rotatelogs.
			New(filename+"%Y%m%d%H.log",
				rotatelogs.WithLinkName(filename+".log"),
				rotatelogs.WithRotationTime(time.Hour),
				rotatelogs.WithRotationSize(100*1024*1024),
				rotatelogs.WithRotationCount(max),
				// rotatelogs.WithRotationTime(time.Minute*5),
			)
	if err != nil {
		std.Panicf("err:%v", err)
	}

	return hook
}
