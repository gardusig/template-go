package test

import (
	"testing"

	"github.com/gardusig/template-go"
)

func TestRelease(t *testing.T) {
	if template.Sum(1, 2) != 3 {
		t.Fail()
	}
}
