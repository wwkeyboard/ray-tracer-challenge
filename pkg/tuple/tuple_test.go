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

func TestTuple_Equals(t *testing.T) {
	t1 := Tuple{
		x: 1,
		y: 2,
		z: 3,
		w: 0,
	}
	t2 := Tuple{
		x: 0.999999999,
		y: 2.000000001,
		z: 3.000000000,
		w: 0,
	}
	assert.True(t, Equal(t1, t2), "Tuple equal handles weird floats")

	t2.x = t2.x - FLOAT_PRECISION
	assert.False(t, Equal(t1, t2), "Tuple not equal handles weird floats")
}

func TestTuple_new_Point(t *testing.T) {
	t1 := Point(4, -4, 3)
	assert.True(t, t1.isPoint(), "Point creates a point")
}

func TestTuple_new_Vector(t *testing.T) {
	t1 := Vector(4, -4, 3)
	assert.True(t, t1.isVector(), "Vector creates a vector")
}

func TestTuple_Add(t *testing.T) {
	t1 := Vector(3, -2, 5)
	t2 := Point(-2, 3, 1)
	result := Point(1, 1, 6)

	assert.Equal(t, t1.Add(t2), result, "adding tuples")
}
