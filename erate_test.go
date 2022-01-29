package g2gcounters

import (
	"testing"
	"time"

	"github.com/msaf1980/g2g/pkg/expvars"

	"github.com/stretchr/testify/assert"
)

func TestERate(t *testing.T) {
	inputs := []int64{1, 3, 2}

	r := NewERate("test erate")

	prev := r.last

	for _, input := range inputs {
		r.Add(input)
	}

	time.Sleep(100 * time.Millisecond)

	got := r.Strings()

	last := r.last

	wantRate := expvars.RoundFloat(float64(SumInt64(inputs)) * (1000000000.0 / float64(last-prev)))

	assert.Equal(t, []expvars.MValue{{Name: "rate", V: wantRate}}, got)

	assert.Equal(t, []expvars.MValue{}, r.Strings())
}
