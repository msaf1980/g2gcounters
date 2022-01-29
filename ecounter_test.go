package g2gcounters

import (
	"testing"

	"github.com/msaf1980/g2g/pkg/expvars"
	"github.com/stretchr/testify/assert"
)

func TestECounter_Add(t *testing.T) {
	c := NewECounter("test ecounter")

	c.Add(2)
	assert.Equal(t, int64(2), c.Value())
	assert.Equal(t, int64(0), c.Value())

	c.Incr()
	assert.Equal(t, int64(1), c.Value())
	assert.Equal(t, int64(0), c.Value())

	c.Incr()
	c.Add(3)
	assert.Equal(t, []expvars.MValue{{Name: "", V: "4"}}, c.Strings())
	assert.Equal(t, int64(0), c.Value())
	assert.Equal(t, []expvars.MValue{}, c.Strings())
}
