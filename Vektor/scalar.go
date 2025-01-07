package vektor

type Scalar float64

func ScalarAdition(x, y Scalar) Scalar {
	return x + y
}
func ScalarSubstraction(x, y Scalar) Scalar {
	return x - y
}
func ScalarDivision(x, y Scalar) Scalar {
	if y == 0 {
		panic("Division by 0 not allowed")
	}
	return x * (1 / y)
}
func ScalarProduct(x, y Scalar) Scalar {
	return x * y
}
func ScalarDistance(x, y Scalar) Scalar {
	return Scalar((x - y).Len())
}

func (s Scalar) Sign() int {
	switch {
	case s > 0:
		return 1
	case s < 0:
		return -1
	default:
		return 0
	}
}

func (s Scalar) Len() float64 {
	switch {
	case s < 0:
		return float64(-s)
	default:
		return float64(s)
	}
}
