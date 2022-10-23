package test

import (
	"testing"

	"github.com/gardusig/template-go"
)

func TestRelease(t *testing.T) {
	const limit = 100
	for x := -limit; x <= limit; x += 1 {
		for y := -limit; y <= limit; y += 1 {
			if template.Sum(x, y) != x+y {
				t.Fail()
			}
		}
	}
}
