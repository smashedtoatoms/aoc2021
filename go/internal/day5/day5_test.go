package day5

import (
	"testing"
)

func TestOne(t *testing.T) {
	in := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	out := findVentOverlaps(in, false)
	expected := 5
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}

func TestTwo(t *testing.T) {
	in := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	out := findVentOverlaps(in, true)
	expected := 12
	if out != expected {
		t.Errorf("wrong number of vent overlaps: got %d, expected %d", out, expected)
		return
	}
}
