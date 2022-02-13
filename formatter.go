package log

import (
	"bytes"
	"fmt"
	"path"
	"strings"
)

type Format interface {
	Format(entry Entry) []byte
}

type Formatter struct {
	Module string
}

func (p *Formatter) Format(entry Entry) []byte {
	var b bytes.Buffer

	if entry.PrefixMsg != "" {
		b.WriteString(entry.PrefixMsg)
		b.WriteString(" ")
	}

	b.WriteString(fmt.Sprintf("(%d.%d) ", entry.Pid, entry.Gid))

	if entry.TraceId != "" {
		b.WriteString("<")
		b.WriteString(entry.TraceId)
		b.WriteString("> ")
	}

	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.9999Z07:00"))

	color := getColorByLevel(Level(entry.Level))

	b.WriteString(color)
	b.WriteString(" [")
	b.WriteString(entry.Level.String()[:4])
	b.WriteString("] ")
	b.WriteString(endColor)

	b.WriteString(strings.TrimSpace(entry.Message))

	b.WriteString(color)
	b.WriteString(" (")
	b.WriteString(path.Join(getPackageName(entry.CallerName), path.Base(entry.File)))
	b.WriteString(":")
	b.WriteString(fmt.Sprintf("%d", entry.CallerLine))
	b.WriteString(")")
	b.WriteString(endColor)

	if entry.SuffixMsg != "" {
		b.WriteString(" ")
		b.WriteString(entry.SuffixMsg)
	}

	b.WriteByte('\n')

	return b.Bytes()
}

const (
	colorBlack   = "\u001B[30m"
	colorRed     = "\u001B[31m"
	colorGreen   = "\u001B[32m"
	colorYellow  = "\u001B[33m"
	colorBlue    = "\u001B[34m"
	colorMagenta = "\u001B[35m"
	colorCyan    = "\u001B[36m"
	colorGray    = "\u001B[37m"
	colorWhite   = "\u001B[38m"
)

const (
	endColor = "\u001B[0m"
)

func getColorByLevel(level Level) string {
	switch level {
	case DebugLevel, TraceLevel:
		return colorGreen

	case WarnLevel:
		return colorYellow

	case ErrorLevel, FatalLevel, PanicLevel:
		return colorRed

	default:
		return colorGreen
	}
}

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}
