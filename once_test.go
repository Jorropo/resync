package resync

import (
	"testing"

	"github.com/cheekybits/is"
)

func TestOnceReset(t *testing.T) {
	is := is.New(t)
	var calls int
	var c Once
	c.Do(func() {
		calls++
	})
	c.Do(func() {
		calls++
	})
	c.Do(func() {
		calls++
	})
	is.Equal(calls, 1)
	c.Reset()
	c.Do(func() {
		calls++
	})
	c.Do(func() {
		calls++
	})
	c.Do(func() {
		calls++
	})
	is.Equal(calls, 2)
}
