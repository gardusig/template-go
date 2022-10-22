package example

import "testing"

func TestSum(t *testing.T) {
	const limit = 100
	for i := -limit; i <= limit; i += 1 {
		for j := -limit; j <= limit; j += 1 {
			validateSum(t, i, j)
		}
	}
}

func validateSum(t *testing.T, x int, y int) {
	if Sum(x, y) != Sum(y, x) {
		t.Errorf("Sum should be commutative")
	}
	if Sum(x, y) != correctSum(y, x) {
		t.Errorf("expected %d but got %d", correctSum(y, x), Sum(x, y))
	}
}

func correctSum(x int, y int) int {
	return x + y
}
