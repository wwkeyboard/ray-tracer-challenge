package tuple

// Tuple represents a point or a vector
type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func (t Tuple) isVector() bool {
	return t.w == 0
}

func (t Tuple) isPoint() bool {
	return t.w == 1
}
