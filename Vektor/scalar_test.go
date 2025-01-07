package vektor

import (
	"fmt"
	"testing"
)

func TestScalar(t *testing.T) {
	tests := []struct {
		name     string
		input    []Scalar
		function func(x Scalar, y Scalar) Scalar
		expected Scalar
	}{
		{
			name:     "Basic Addition",
			input:    []Scalar{Scalar(4), Scalar(6)},
			function: ScalarAdition,
			expected: Scalar(10),
		},
		{
			name:     "Basic Substraction",
			input:    []Scalar{Scalar(10), Scalar(6)},
			function: ScalarSubstraction,
			expected: Scalar(4),
		},
		{
			name:     "Basic Division",
			input:    []Scalar{Scalar(4), Scalar(2)},
			function: ScalarDivision,
			expected: Scalar(2),
		},
		{
			name:     "Basic Product",
			input:    []Scalar{Scalar(4), Scalar(6)},
			function: ScalarProduct,
			expected: Scalar(24),
		}, {
			name:     "Distance between 2 Scalars, first",
			input:    []Scalar{Scalar(10), Scalar(4)},
			function: ScalarDistance,
			expected: Scalar(6),
		}, {
			name:     "Distance between 2 Scalars, second",
			input:    []Scalar{Scalar(4), Scalar(-2)},
			function: ScalarDistance,
			expected: Scalar(6),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.input[0], tt.input[1])
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			} else {
				fmt.Printf("Expected %v, got %v\n", tt.expected, result)
			}
		})
	}

	divZeroTest := struct {
		name     string
		input    []Scalar
		function func(x Scalar, y Scalar) Scalar
	}{
		name:     "Division By 0",
		input:    []Scalar{Scalar(10), Scalar(0)},
		function: ScalarDivision,
	}

	t.Run(divZeroTest.name, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic")
			} else {
				if r != "Division by 0 not allowed" {
					t.Errorf("Unexpected panic message: %v", r)
				}
			}
		}()

		ScalarDivision(divZeroTest.input[0], divZeroTest.input[1])
	})
}

func TestScalarSign(t *testing.T) {
	tests := []struct {
		name     string
		input    Scalar
		expected int
	}{
		{
			name:     "No sign",
			input:    Scalar(0),
			expected: 0,
		},
		{
			name:     "Positive sign",
			input:    Scalar(10),
			expected: 1,
		},
		{
			name:     "Negative sign",
			input:    Scalar(-10),
			expected: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Sign()
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestScalarLen(t *testing.T) {
	tests := []struct {
		name     string
		input    Scalar
		expected float64
	}{
		{
			name:     "Positive Scalar",
			input:    Scalar(6),
			expected: 6,
		},
		{
			name:     "Scalar with 0 length",
			input:    Scalar(0),
			expected: 0,
		},
		{
			name:     "Negative Scalar",
			input:    Scalar(-1),
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Len()
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
