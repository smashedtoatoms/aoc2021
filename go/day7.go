package aoc

import (
	"strconv"
	"strings"
)

func RunDay7(path string) (int, int) {
	lines := getLines(path, "inputs/day7.txt")
	positions, err := lineToCrabPositions(lines[0])
	if err != nil {
		panic(err)
	}

	return findCheapestFuelCost(positions, false), findCheapestFuelCost(positions, true)
}

func findCheapestFuelCost(positions []int, shouldUseCrabCalculation bool) int {
	fuelCosts := make(map[int]int)
	min := min(positions)
	max := max(positions)
	for i := min; i < max; i++ {
		for _, position := range positions {
			if _, ok := fuelCosts[i]; !ok {
				fuelCosts[i] = 0
			}
			if shouldUseCrabCalculation {
				for j := 1; j <= abs(position-i); j++ {
					fuelCosts[i] += j
				}
			} else {
				fuelCosts[i] += abs(position - i)
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
