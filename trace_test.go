package log

import (
	"testing"
)

func TestGenTraceId(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(GenTraceId())
	}
}
