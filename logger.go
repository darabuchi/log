package log

import (
	"bytes"
	"fmt"
	"github.com/petermattis/goid"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	level Level

	outList []zapcore.WriteSyncer
}

func newLogger() *Logger {
	return &Logger{
		level:   DebugLevel,
		outList: []zapcore.WriteSyncer{os.Stdout},
	}
}

func (p *Logger) SetLevel(level Level) {
	p.level = level
}

func (p *Logger) SetOutput(write zapcore.WriteSyncer) {
	p.outList = []zapcore.WriteSyncer{write}
}

func (p *Logger) AddOutput(write zapcore.WriteSyncer) {
	p.outList = append(p.outList, write)
}

func (p *Logger) Log(level Level, args ...interface{}) {
	p.log(level, fmt.Sprint(args...))
}

func (p *Logger) Logf(level Level, format string, args ...interface{}) {
	p.log(level, fmt.Sprintf(format, args))
}

func (p *Logger) log(level Level, msg string) {
	if !p.levelEnabled(level) {
		return
	}

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("(%d.%d) ", pid, goid.Get()))

	b.WriteString(time.Now().Format("2006-01-02 15:04:05.9999Z07:00"))

	color := getColorByLevel(level)

	b.WriteString(color)
	b.WriteString(" [")
	b.WriteString(level.String()[:4])
	b.WriteString("] ")
	b.WriteString(endColor)

	b.WriteString(strings.TrimSpace(msg))

	var callerName string
	pc, file, callerLine, ok := runtime.Caller(4)
	if ok {
		callerName = runtime.FuncForPC(pc).Name()
	}

	b.WriteString(color)
	b.WriteString(" (")
	b.WriteString(path.Join(getPackageName(callerName), path.Base(file)))
	b.WriteString(":")
	b.WriteString(fmt.Sprintf("%d", callerLine))
	b.WriteString(")")
	b.WriteString(endColor)

	b.WriteByte('\n')

	p.write(level, b.Bytes())
}

func (p *Logger) write(level Level, buf []byte) {
	for _, out := range p.outList {
		_, _ = out.Write(buf)
	}
	if level > ErrorLevel {
		p.Sync()
	}
}

func (p *Logger) levelEnabled(level Level) bool {
	return p.level >= level
}

func (p *Logger) Trace(args ...interface{}) {
	p.Log(TraceLevel, args...)
}

func (p *Logger) Debug(args ...interface{}) {
	p.Log(DebugLevel, args...)
}

func (p *Logger) Print(args ...interface{}) {
	p.Log(DebugLevel, args...)
}

func (p *Logger) Info(args ...interface{}) {
	p.Log(InfoLevel, args...)
}

func (p *Logger) Warn(args ...interface{}) {
	p.Log(WarnLevel, args...)
}

func (p *Logger) Error(args ...interface{}) {
	p.Log(ErrorLevel, args...)
}

func (p *Logger) Panic(args ...interface{}) {
	p.Log(PanicLevel, args...)
}

func (p *Logger) Fatal(args ...interface{}) {
	p.Log(FatalLevel, args...)
}

func (p *Logger) Tracef(format string, args ...interface{}) {
	p.Logf(TraceLevel, format, args...)
}

func (p *Logger) Printf(format string, args ...interface{}) {
	p.Logf(DebugLevel, format, args...)
}

func (p *Logger) Debugf(format string, args ...interface{}) {
	p.Logf(DebugLevel, format, args...)
}

func (p *Logger) Infof(format string, args ...interface{}) {
	p.Logf(InfoLevel, format, args...)
}

func (p *Logger) Warnf(format string, args ...interface{}) {
	p.Logf(WarnLevel, format, args...)
}

func (p *Logger) Errorf(format string, args ...interface{}) {
	p.Logf(ErrorLevel, format, args...)
}

func (p *Logger) Fatalf(format string, args ...interface{}) {
	p.Logf(FatalLevel, format, args...)
}

func (p *Logger) Panicf(format string, args ...interface{}) {
	p.Logf(PanicLevel, format, args...)
}

func (p *Logger) Sync() {
	for _, out := range p.outList {
		_ = out.Sync()
	}
}
