package day5

import (
	"strconv"
	"strings"

	"smashedtoatoms.com/aoc2021/internal/shared"
)

type coordinate struct {
	X int
	Y int
}

// RunDay5 finds vent overlaps based on the input provided.
func RunDay5(path string) (int, int) {
	lines := shared.FilterStrings(shared.GetLines(path, "inputs/day5.txt"), "")
	return findVentOverlaps(lines, true), findVentOverlaps(lines, false)
}

func findVentOverlaps(lines []string, includeDiagonals bool) int {
	points := make(map[coordinate]int)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		sparts := strings.Split(parts[0], ",")
		fparts := strings.Split(parts[1], ",")
		x1, _ := strconv.Atoi(sparts[0])
		y1, _ := strconv.Atoi(sparts[1])
		x2, _ := strconv.Atoi(fparts[0])
		y2, _ := strconv.Atoi(fparts[1])

		if x1 == x2 || y1 == y2 {
			for _, x := range rangeInclusive(x1, x2) {
				for _, y := range rangeInclusive(y1, y2) {
					points[coordinate{x, y}]++
				}
			}
		} else if includeDiagonals {
			xs := rangeInclusive(x1, x2)
			ys := rangeInclusive(y1, y2)
			for i, x := range xs {
				points[coordinate{x, ys[i]}]++
			}
		}
	}

	result := 0
	for _, value := range points {
		if value > 1 {
			result++
		}
	}

	return result
}

func rangeInclusive(a int, b int) []int {
	var result []int
	direction := 1
	if b < a {
		direction = -1
	}

	for i := a; i != b; i += direction {
		result = append(result, i)
	}

	result = append(result, b)

	return result
}
