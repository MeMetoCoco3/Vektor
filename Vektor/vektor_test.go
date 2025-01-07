package vektor

import "testing"

func TestVektor(t *testing.T) {
	tests := []struct {
		name     string
		input    []Vektor
		function func(x Vektor, y Vektor) Vektor
		expected Vektor
	}{
		{
			name:     "Basic Addition",
			input:    []Vektor{Vektor{4, 10}, Vektor{6, -4}},
			function: VektorAdition,
			expected: Vektor{10, 6},
		},
		{
			name:     "Basic Substraction",
			input:    []Vektor{Vektor{10, 3}, Vektor{6, 8}},
			function: VektorSubstraction,
			expected: Vektor{4, -5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.input[0], tt.input[1])
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}

}
func TestVektorMagnitude(t *testing.T) {
	tests := []struct {
		name     string
		input    Vektor
		expected float64
	}{
		{
			name:     "Magnitude 1",
			input:    Vektor{4, 10},
			expected: 10.770329614269007,
		},
		{
			name:     "Magnitude 2",
			input:    Vektor{-7, 8},
			expected: 10.63014581273465,
		},
		{
			name:     "Magnitude 3",
			input:    Vektor{-6, -8},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := VektorMagnitude(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}

}
func TestVektorDistance(t *testing.T) {
	tests := []struct {
		name     string
		input    []Vektor
		expected float64
	}{
		{
			name:     "Distance 1",
			input:    []Vektor{Vektor{4, 10}, Vektor{6, -4}},
			expected: 14.142135623730951,
		},
		{
			name:     "Distance 2",
			input:    []Vektor{Vektor{10, 3}, Vektor{6, 8}},
			expected: 6.4031242374328485,
		},
		{
			name:     "Distance 3",
			input:    []Vektor{Vektor{-10, 3}, Vektor{-6, -8}},
			expected: 11.704699910719626,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := VektorDistance(tt.input[0], tt.input[1])
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}

}
