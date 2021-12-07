package day6

import (
	"fmt"
	"strings"

	"smashedtoatoms.com/aoc2021/internal/shared"
)

// RunDay6 counts lanternfish as they breed over time.
func RunDay6(path string) (int, int) {
	lines := shared.GetLines(path, "inputs/day6.txt")
	inputData, err := shared.StringsToInts(strings.Split(lines[0], ","))
	if err != nil {
		panic(err)
	}
	return countLanternfish(inputData, 80), countLanternfish(inputData, 256)
}

func countLanternfish(inputData []int, days int) int {
	fishCount := make(map[string]int)
	lifecycles := make(map[int]int)
	count := 0

	for _, fish := range inputData {
		lifecycles[fish]++
	}

	for fish, val := range lifecycles {
		count += (countFish(fish, days, 1, fishCount) * val)
	}

	return count
}

func countFish(fish int, days int, count int, fishCount map[string]int) int {
	fish--
	days--

	if fish < 0 && days >= 0 {
		fish = 6

		key := fmt.Sprintf("%d:%d", days, 8)

		if fishCount[key] != 0 {
			count += fishCount[key]
		} else {
			growth := countFish(8, days, 1, fishCount)
			count += growth
			fishCount[key] = growth
		}

	}

	if days < 0 {
		return count
	}

	return countFish(fish, days, count, fishCount)
}
