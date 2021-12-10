package day11

import (
	"testing"
)

func Test1(t *testing.T) {
	in := [][]int{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}
	out := countFlashingOctopi(in, 100)
	expected := 1656
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}

func Test2(t *testing.T) {
	in := [][]int{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}
	out := findSimultaneousFlashOfOctopi(in)
	expected := 195
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}
