package workers

import (
	"sync/atomic"
	"testing"

	"github.com/lpabon/lputils/tests"
)

func TestWorkers(t *testing.T) {
	sum := int64(0)

	Workers(10, func(v interface{}) {
		// consumer
		i := v.(int64)
		atomic.AddInt64(&sum, i)
	}, func(ch chan interface{}) {
		//producer
		for p := int64(0); p < int64(100); p++ {
			ch <- p
		}
	})

	for p := int64(0); p < int64(100); p++ {
		sum = sum - p
	}
	tests.Assert(t, sum == 0)
}
