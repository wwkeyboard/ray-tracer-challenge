package tuple

import "math"

const FLOAT_PRECISION = 0.00001

// Tuple represents a point or a vector
type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

// Point returns a Tuple representing a point at x, y, z
func Point(x, y, z float64) Tuple {
	return Tuple{
		x: x,
		y: y,
		z: z,
		w: 1,
	}
}

// Vector returns a Tuple representing a vector at x, y, z
func Vector(x, y, z float64) Tuple {
	return Tuple{
		x: x,
		y: y,
		z: z,
		w: 0,
	}
}

// Equal test that the tuples are equal to within a good enough precision.
// This accounts for some weirdness in floating point math.
func Equal(t1, t2 Tuple) bool {
	if math.Abs(t1.x-t2.x) < FLOAT_PRECISION &&
		math.Abs(t1.y-t2.y) < FLOAT_PRECISION &&
		math.Abs(t1.z-t2.z) < FLOAT_PRECISION &&
		t1.w == t2.w {
		return true
	}
	return false
}

func (t Tuple) isVector() bool {
	return t.w == 0
}

func (t Tuple) isPoint() bool {
	return t.w == 1
}
