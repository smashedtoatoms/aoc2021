package aoc

import (
	"testing"
)

func TestDay7(t *testing.T) {
	input := []string{""}
	derps := doIt(input)
	expected := 0
	if derps != expected {
		t.Errorf("wrong number: got %d, expected %d", derps, expected)
		return
	}
}
