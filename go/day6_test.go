package aoc

import (
	"testing"
)

func TestDay6x1(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	days := 80
	lanternfish := countLanternfish(input, days)
	expected := 5934
	if lanternfish != expected {
		t.Errorf("wrong number of lanternfish: got %d, expected %d", lanternfish, expected)
		return
	}
}

func TestDay6x2(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	days := 256
	lanternfish := countLanternfish(input, days)
	expected := 26984457539
	if lanternfish != expected {
		t.Errorf("wrong number of lanternfish: got %d, expected %d", lanternfish, expected)
		return
	}
}
