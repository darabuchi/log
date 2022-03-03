package log

import "time"

type Entry struct {
	Pid     int
	Gid     int64
	TraceId string

	Time  time.Time
	Level Level
	File  string

	Message    string
	CallerName string
	CallerLine int
	CallerDir  string
	CallerFunc string

	PrefixMsg string
	SuffixMsg string
}
