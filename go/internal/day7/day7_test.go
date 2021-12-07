package day7

import (
	"testing"
)

func Test1(t *testing.T) {
	in := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	out := findCheapestFuelCost(in, false)
	expected := 37
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}

func Test2(t *testing.T) {
	in := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	out := findCheapestFuelCost(in, true)
	expected := 168
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}
