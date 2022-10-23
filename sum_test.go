package template

import "testing"

func TestSum(t *testing.T) {
	const limit = 100
	for x := -limit; x <= limit; x += 1 {
		for y := -limit; y <= limit; y += 1 {
			validateSum(t, x, y)
		}
	}
}

func validateSum(t *testing.T, x int, y int) {
	if Sum(x, y) != Sum(y, x) {
		t.Errorf("Sum should be commutative")
	}
	expectedSum := x + y
	if Sum(x, y) != expectedSum {
		t.Errorf("Expected %d but got %d", expectedSum, Sum(x, y))
	}
}
