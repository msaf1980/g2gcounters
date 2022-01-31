package g2gcounters

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/msaf1980/g2g/pkg/expvars"
	"github.com/stretchr/testify/assert"
)

func TestTimer(t *testing.T) {
	inputs := []float64{1.0, 3.0, 2.0}
	timer := NewTimer("test timer")

	prev := timer.last

	for _, input := range inputs {
		timer.Add(input)
	}

	time.Sleep(100 * time.Millisecond)

	results := timer.Strings()

	last := timer.last

	wantRate := float64(len(inputs)) * (1000000000.0 / float64(last-prev))

	assert.Equal(t, 9, len(results))

	for _, got := range results {
		switch got.Name {
		case ".min":
			assert.Equal(t, expvars.MValue{Name: ".min", V: "1"}, got)
		case ".max":
			assert.Equal(t, expvars.MValue{Name: ".max", V: "3"}, got)
		case ".median":
			assert.Equal(t, expvars.MValue{Name: ".median", V: "1.5"}, got)
		case ".p90":
			assert.Equal(t, expvars.MValue{Name: ".p90", V: "2.5"}, got)
		case ".p95":
			assert.Equal(t, expvars.MValue{Name: ".p95", V: "2.5"}, got)
		case ".p99":
			assert.Equal(t, expvars.MValue{Name: ".p99", V: "2.5"}, got)
		case ".sum":
			assert.Equal(t, expvars.MValue{Name: ".sum", V: "6"}, got)
		case ".count":
			assert.Equal(t, expvars.MValue{Name: ".count", V: "3"}, got)
		case ".rate":
			assert.Equal(t, expvars.MValue{Name: ".rate", V: expvars.RoundFloat(wantRate)}, got)
		default:
			t.Errorf("unexpected metric: %s", got.Name)
		}
	}
}

func BenchmarkTimer(b *testing.B) {
	concurrencyLevels := []int{5, 10, 20, 50}
	timer := NewTimer("test timer")
	for _, clients := range concurrencyLevels {

		b.Run(fmt.Sprintf("%d_clients", clients), func(b *testing.B) {
			sem := make(chan struct{}, clients)
			wg := sync.WaitGroup{}

			for i := 0; i < clients; i++ {
				wg.Add(1)
				go func() {
					<-sem
					// Test routine
					for n := 0; n < b.N; n++ {
						timer.Add(1)
					}
					// End test routine
					wg.Done()
				}()

			}

			ticker := time.NewTicker(10 * time.Millisecond)
			go func() {
				// Test routine
				for _ = range ticker.C {
					vals := timer.Strings()
					_ = vals
				}
				// End test routine
			}()

			close(sem)
			wg.Wait()
			ticker.Stop()
		})
	}
}
