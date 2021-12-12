package day12

import (
	"testing"
)

func Test1(t *testing.T) {
	in := [][]string{
		{"start", "A"},
		{"start", "b"},
		{"A", "c"},
		{"A", "b"},
		{"b", "d"},
		{"A", "end"},
		{"b", "end"},
	}
	out := findAllUndergroundPathsVisitOnce(in)
	expected := 10
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}

func Test2(t *testing.T) {
	in := [][]string{
		{"start", "A"},
		{"start", "b"},
		{"A", "c"},
		{"A", "b"},
		{"b", "d"},
		{"A", "end"},
		{"b", "end"},
	}
	out := findAllUndergroundPathsVisitTwice(in)
	expected := 36
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}
