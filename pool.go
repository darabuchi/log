package log

import (
	"bytes"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func GetBuffer() *bytes.Buffer {
	return pool.Get().(*bytes.Buffer)
}

func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	pool.Put(buf)
}
