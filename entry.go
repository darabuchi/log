package log

import "time"

type Entry struct {
	Time  time.Time
	Level Level
	File  string

	Message    string
	CallerName string
	CallerLine int
}
