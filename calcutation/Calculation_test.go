package calcutation_test

import (
	"goroutines-test/calcutation"
	"testing"
)

func TestCalculate(t *testing.T) {
	if calcutation.Calculate(2, 2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		inputA int
		inputB int
		output int
	}{
		{2, 2, 4},
		{1, 1, 2},
		{-4, -2, -6},
		{-2, 6, 4},
		{3, -5, -2},
	}

	for _, test := range tests {
		if outputX := calcutation.Calculate(test.inputA, test.inputB); outputX != test.output {
			t.Error("Test failed: value A {} value B {} -> expected {}", test.inputA, test.inputB, test.output)
		}
	}
}
