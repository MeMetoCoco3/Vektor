package main

import (
	"fmt"
	"vektor/Vektor"
)

func main() {
	input := []vektor.Scalar{vektor.Scalar(4), vektor.Scalar(-2)}

	fmt.Printf("Sc1: %.02f + Sc2: %.02f = %.02f", input[0], input[1], vektor.ScalarAdition(input[0], input[1]))
}
