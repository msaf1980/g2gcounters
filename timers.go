package g2gcounters

import (
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/msaf1980/g2g"
)

// Int is a 64-bit integer variable that satisfies the Var interface.
type Timer struct {
	vals []float64

	last int64

	lock sync.Mutex
}

func NewTimer(name string) *Timer {
	t := new(Timer)
	t.last = time.Now().UnixNano()

	g2g.MPublish(name, t)

	return t
}

func (t *Timer) Add(v float64) {
	t.lock.Lock()

	t.vals = append(t.vals, v)

	t.lock.Unlock()
}

func (t *Timer) Strings() []g2g.MValue {
	var vals []float64

	t.lock.Lock()

	now := time.Now().UnixNano()
	prev := t.last

	t.last = now

	n := len(t.vals)
	if n > 0 {
		vals = t.vals
		t.vals = make([]float64, 0, cap(vals))
	}

	t.lock.Unlock()

	if n == 0 {
		return []g2g.MValue{
			{Name: "count", V: "0"},
			{Name: "min", V: "0"},
			{Name: "max", V: "0"},
			{Name: "median", V: "0"},
			{Name: "p90", V: "0"},
			{Name: "p95", V: "0"},
			{Name: "p99", V: "0"},
			{Name: "sum", V: "0"},
			{Name: "rate", V: "0"},
		}
	} else {
		sort.Float64s(vals)

		count := int64(len(vals))

		durations := now - prev
		rate := float64(count) * (1000000000.0 / float64(durations))

		sum := Sum(vals)
		median, _ := Percentile(vals, 0.5)
		p90, _ := Percentile(vals, 0.9)
		p95, _ := Percentile(vals, 0.95)
		p99, _ := Percentile(vals, 0.99)

		return []g2g.MValue{
			{Name: "count", V: strconv.FormatInt(count, 10)},
			{Name: "min", V: g2g.RoundFloat(vals[0])},
			{Name: "max", V: g2g.RoundFloat(vals[len(vals)-1])},
			{Name: "median", V: g2g.RoundFloat(median)},
			{Name: "p90", V: g2g.RoundFloat(p90)},
			{Name: "p95", V: g2g.RoundFloat(p95)},
			{Name: "p99", V: g2g.RoundFloat(p99)},
			{Name: "sum", V: g2g.RoundFloat(sum)},
			{Name: "rate", V: g2g.RoundFloat(rate)},
		}
	}
}
