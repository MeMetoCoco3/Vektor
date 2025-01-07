package vektor

import "math"

type Vektor struct {
	X, Y float64
}

func VektorAdition(a, b Vektor) Vektor {
	return Vektor{X: a.X + b.X, Y: a.Y + b.Y}
}
func VektorSubstraction(a, b Vektor) Vektor {
	return Vektor{X: a.X - b.X, Y: a.Y - b.Y}
}
func VektorMagnitude(a Vektor) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}
func VektorDistance(a, b Vektor) float64 {
	dx := b.X - a.X
	dy := b.Y - a.Y
	return math.Sqrt(dx*dx + dy*dy)
}
