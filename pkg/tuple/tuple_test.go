package tuple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTuple_isPoint(t *testing.T) {
	t1 := Tuple{
		x: 4.3,
		y: -4.2,
		z: 3.1,
		w: 1.0,
	}
	assert.Equal(t, t1.isPoint(), true, "t1 is a point")
	assert.Equal(t, t1.isVector(), false, "t1 is not a vector")

	t1.w = 0.0
	assert.Equal(t, t1.isPoint(), false, "t1 is no longer a point")
	assert.Equal(t, t1.isVector(), true, "t1 is now a vector")
}

func TestTuple_new_Point(t *testing.T) {
	t1 := Point(4, -4, 3)
	assert.True(t, t1.isPoint(), "Tuple.Point creates a point")
}
