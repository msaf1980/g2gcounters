package g2gcounters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter_Add(t *testing.T) {
	c := NewCounter("test")

	c.Add(2)
	assert.Equal(t, int64(2), c.Value())
	assert.Equal(t, int64(0), c.Value())

	c.Incr()
	assert.Equal(t, int64(1), c.Value())
	assert.Equal(t, int64(0), c.Value())

	c.Incr()
	c.Add(3)
	assert.Equal(t, "4", c.String())
	assert.Equal(t, int64(0), c.Value())
}
