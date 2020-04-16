package main

import (
	"testing"
)

func TestTableAddition(t *testing.T) {
	var tests = []struct {
		input_a  int
		input_b  int
		expected int
	}{
		{2, 4, 6},
		{-1, -1, -2},
		{0, 0, 0},
		{2, -4, -2},
	}

	for _, test := range tests {
		if output := Addition(test.input_a, test.input_b); output != test.expected {
			t.Errorf("Test failed: %d input_a, %d input_b, was expecting %d but received %d", test.input_a, test.input_b, test.expected, output)
		}
	}

}
