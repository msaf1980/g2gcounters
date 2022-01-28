package g2gcounters

import (
	"testing"
	"time"

	"github.com/msaf1980/g2g"
	"github.com/stretchr/testify/assert"
)

func TestTimer(t *testing.T) {
	inputs := []float64{1.0, 3.0, 2.0}
	timer := NewTimer("test")

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
		case "min":
			assert.Equal(t, g2g.MValue{Name: "min", V: "1"}, got)

		case "max":
			assert.Equal(t, g2g.MValue{Name: "max", V: "3"}, got)
		case "median":
			assert.Equal(t, g2g.MValue{Name: "median", V: "1.5"}, got)
		case "p90":
			assert.Equal(t, g2g.MValue{Name: "p90", V: "2.5"}, got)
		case "p95":
			assert.Equal(t, g2g.MValue{Name: "p95", V: "2.5"}, got)
		case "p99":
			assert.Equal(t, g2g.MValue{Name: "p99", V: "2.5"}, got)
		case "sum":
			assert.Equal(t, g2g.MValue{Name: "sum", V: "6"}, got)
		case "count":
			assert.Equal(t, g2g.MValue{Name: "count", V: "3"}, got)
		case "rate":
			assert.Equal(t, g2g.MValue{Name: "rate", V: g2g.RoundFloat(wantRate)}, got)
		default:
			t.Errorf("unexpected metric: %s", got.Name)
		}
	}
}
