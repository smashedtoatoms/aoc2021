package shared

import (
	"testing"
)

func TestAbsPositiveNumber(t *testing.T) {
	in := 1
	out := Abs(in)
	expected := 1
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}

func TestAbsNegativeNumber(t *testing.T) {
	in := -1
	out := Abs(in)
	expected := 1
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}
