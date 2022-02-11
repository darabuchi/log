package log

import (
	"strings"
)

type Formatter struct {
	module string
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
