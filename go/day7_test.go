package aoc

import (
	"testing"
)

func TestDay7x1(t *testing.T) {
	positions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	cheapestFuelCost := findCheapestFuelCost(positions, false)
	expected := 37
	if cheapestFuelCost != expected {
		t.Errorf("wrong fuel cost: got %d, expected %d", cheapestFuelCost, expected)
		return
	}
}

func TestDay7x2(t *testing.T) {
	positions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	cheapestFuelCost := findCheapestFuelCost(positions, true)
	expected := 168
	if cheapestFuelCost != expected {
		t.Errorf("wrong fuel cost: got %d, expected %d", cheapestFuelCost, expected)
		return
	}
}
