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

type FormatFull interface {
	Format
	// ParsingAndEscaping Default should be on
	ParsingAndEscaping(disable bool)

	// DisableCaller Default should be on
	Caller(disable bool)

	Clone() Format
}

type Formatter struct {
	Module string

	DisableParsingAndEscaping bool
	DisableCaller             bool
}

func (p *Formatter) format(entry Entry) []byte {
	var b bytes.Buffer

	if entry.PrefixMsg != "" {
		b.WriteString(entry.PrefixMsg)
		b.WriteString(" ")
	}

	b.WriteString(fmt.Sprintf("(%d.%d) ", entry.Pid, entry.Gid))

	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.999Z07:00"))

	color := getColorByLevel(entry.Level)

	b.WriteString(color)
	b.WriteString(" [")
	b.WriteString(entry.Level.String())
	b.WriteString("] ")
	b.WriteString(colorEnd)

	b.WriteString(strings.TrimSpace(entry.Message))

	if !p.DisableCaller || entry.TraceId != "" {
		b.WriteString(colorCyan)
		b.WriteString(" [ ")

		if !p.DisableCaller {
			b.WriteString(path.Join(entry.CallerDir, path.Base(entry.File)))
			b.WriteString(":")
			b.WriteString(fmt.Sprintf("%d", entry.CallerLine))
			b.WriteString(" ")
			b.WriteString(entry.CallerFunc)
			b.WriteString(" ")
		}

		if entry.TraceId != "" {
			b.WriteString(entry.TraceId)
			b.WriteString(" ")
		}

		b.WriteString("]")
		b.WriteString(colorEnd)
	}

	if entry.SuffixMsg != "" {
		b.WriteString(entry.SuffixMsg)
	}

	b.WriteByte('\n')

	return b.Bytes()
}

func (p *Formatter) Format(entry Entry) []byte {
	if p.DisableParsingAndEscaping {
		return p.format(entry)
	}

	var b bytes.Buffer
	for _, msg := range strings.Split(entry.Message, "\n") {
		entry.Message = msg
		b.Write(p.format(entry))
	}

	return b.Bytes()
}

func (p *Formatter) ParsingAndEscaping(disable bool) {
	p.DisableParsingAndEscaping = disable
}

func (p *Formatter) Caller(disable bool) {
	p.DisableCaller = disable
}

func (p *Formatter) Clone() Format {
	return &Formatter{
		Module:                    p.Module,
		DisableParsingAndEscaping: p.DisableParsingAndEscaping,
		DisableCaller:             p.DisableCaller,
	}
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

	colorEnd = "\u001B[0m"
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

func SplitPackageName(f string) (string, string) {
	slashIndex := strings.LastIndex(f, "/")
	if slashIndex > 0 {
		idx := strings.Index(f[slashIndex:], ".") + slashIndex
		return f[:idx], f[idx:]
	}

	return f, ""
}
