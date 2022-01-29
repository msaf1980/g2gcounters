package g2gcounters

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRate(t *testing.T) {
	inputs := []int64{1, 3, 2}

	r := NewRate("test rate")

	prev := r.last

	for _, input := range inputs {
		r.Add(input)
	}

	time.Sleep(100 * time.Millisecond)

	got := r.Value()

	last := r.last

	wantRate := float64(SumInt64(inputs)) * (1000000000.0 / float64(last-prev))

	assert.Equal(t, wantRate, got)

	assert.Equal(t, "0", r.String())
}
