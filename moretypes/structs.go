package moretypes

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) Scale(p float64) {
	v.X = v.X * p
	v.Y = v.Y * p
}

