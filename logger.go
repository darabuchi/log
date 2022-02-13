package log

import (
	"fmt"
	"github.com/aofei/sandid"
	"github.com/petermattis/goid"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"sync"
	"time"
)

type Logger struct {
	level Level

	outList []zapcore.WriteSyncer

	Format Format

	callerDepth int

	PrefixMsg string
	SuffixMsg string

	one sync.Once

	traceMap sync.Map
}

func newLogger() *Logger {
	return &Logger{
		level:       DebugLevel,
		outList:     []zapcore.WriteSyncer{os.Stdout},
		Format:      &Formatter{},
		callerDepth: 4,
	}
}

func (p *Logger) SetCallerDepth(callerDepth int) {
	p.callerDepth = callerDepth
}

func (p *Logger) SetPrefixMsg(prefixMsg string) {
	p.PrefixMsg = prefixMsg
}

func (p *Logger) SetSuffixMsg(suffixMsg string) {
	p.SuffixMsg = suffixMsg
}

func (p *Logger) Clone() *Logger {
	return &Logger{
		level:   p.level,
		outList: p.outList,
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

	entry := Entry{
		Pid:       pid,
		Gid:       goid.Get(),
		Time:      time.Now(),
		Level:     level,
		Message:   msg,
		SuffixMsg: p.SuffixMsg,
		PrefixMsg: p.PrefixMsg,
	}

	entry.TraceId = p.GetTrace(entry.Gid)

	var pc uintptr
	var ok bool
	pc, entry.File, entry.CallerLine, ok = runtime.Caller(p.callerDepth)
	if ok {
		entry.CallerName = runtime.FuncForPC(pc).Name()
	}

	p.write(level, p.Format.Format(entry))
}

func (p *Logger) GetTrace(gid int64) string {
	tid, ok := p.traceMap.Load(gid)
	if !ok {
		return ""
	}
	return tid.(string)
}

func (p *Logger) SetTrace(gid int64, traceId string) {
	if traceId == "" {
		traceId = sandid.New().String()
	}
	p.traceMap.Store(gid, traceId)
}

func (p *Logger) DelTrace(gid int64) {
	p.traceMap.Delete(gid)
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
