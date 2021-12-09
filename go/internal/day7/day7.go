package day7

import (
	"strconv"
	"strings"

	"smashedtoatoms.com/aoc2021/internal/shared"
)

// Run finds the cheapest position to launch a spacecraft to based on
// internal calculations and crab calculations.
func Run(path string) (int, int) {
	lines := shared.GetLines(path, "inputs/day7.txt")
	positions, err := lineToCrabPositions(lines[0])
	if err != nil {
		panic(err)
	}

	return findCheapestFuelCost(positions, false), findCheapestFuelCost(positions, true)
}

func findCheapestFuelCost(positions []int, shouldUseCrabCalculation bool) int {
	fuelCosts := make(map[int]int)
	min := shared.Min(positions)
	max := shared.Max(positions)
	for i := min; i < max; i++ {
		for _, position := range positions {
			if _, ok := fuelCosts[i]; !ok {
				fuelCosts[i] = 0
			}
			if shouldUseCrabCalculation {
				for j := 1; j <= shared.Abs(position-i); j++ {
					fuelCosts[i] += j
				}
			} else {
				fuelCosts[i] += shared.Abs(position - i)
			}
		}
	}
	return getCheapestPosition(fuelCosts)
}

func getCheapestPosition(positions map[int]int) int {
	var min int
	for _, v := range positions {
		if min == 0 || v < min {
			min = v
		}
	}
	return min
}

func lineToCrabPositions(line string) ([]int, error) {
	var positions []int
	for _, i := range strings.Split(line, ",") {
		position, err := strconv.Atoi(i)
		if err != nil {
			return []int{}, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}
