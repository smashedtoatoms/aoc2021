package aoc

import (
	"testing"
)

func TestDay5x1(t *testing.T) {
	input := []string{
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
	overlaps := findVentOverlaps(input, false)
	expected := 5
	if overlaps != expected {
		t.Errorf("wrong number of vent overlaps: got %d, expected %d", overlaps, expected)
		return
	}
}

func TestDay5x2(t *testing.T) {
	input := []string{
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
	overlaps := findVentOverlaps(input, true)
	expected := 12
	if overlaps != expected {
		t.Errorf("wrong number of vent overlaps: got %d, expected %d", overlaps, expected)
		return
	}
}
