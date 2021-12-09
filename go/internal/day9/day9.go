package day9

import (
	"sort"

	"smashedtoatoms.com/aoc2021/internal/shared"
)

type basin map[coordinate]bool

type coordinate struct {
	X int
	Y int
}

type lowPoint struct {
	Value      int
	Coordinate coordinate
}

// Run to find risk level and large low-lying places to avoid.
func Run(path string) (int, int) {
	elevations, err := shared.SplitToInts(shared.GetLines(path, "inputs/day9.txt"))
	if err != nil {
		panic(err)
	}

	return findRiskLevel(elevations), findLargestPlacesToAvoid(elevations)
}

func findRiskLevel(elevations [][]int) int {
	riskLevel := 0
	for _, lp := range findLowPoints(elevations) {
		riskLevel += lp.Value + 1
	}
	return riskLevel
}

func findLargestPlacesToAvoid(elevations [][]int) int {
	basins := findBasins(elevations)
	basinSizes := make([]int, 0)
	for _, basin := range basins {
		basinSizes = append(basinSizes, len(basin))
	}
	sort.Ints(basinSizes)
	len := len(basinSizes)
	return basinSizes[len-1] * basinSizes[len-2] * basinSizes[len-3]
}

func findBasin(lp lowPoint, elevations [][]int, b basin) {
	if b[lp.Coordinate] || getHeight(elevations, lp.Coordinate) == 9 {
		return
	}
	b[lp.Coordinate] = true
	for _, neighbor := range findNeighbors(elevations, lp.Coordinate) {
		if lp.Value <= getHeight(elevations, neighbor) {
			findBasin(lowPoint{lp.Value, neighbor}, elevations, b)
		}
	}
}

func findBasins(elevations [][]int) []basin {
	basins := make([]basin, 0)
	for _, lp := range findLowPoints(elevations) {
		b := getBasinForPoint(elevations, lp)
		basins = append(basins, b)
	}
	return basins
}

func findLowPoints(elevations [][]int) []lowPoint {
	y := len(elevations)
	lowPoints := make([]lowPoint, 0)
	for i := 0; i < len(elevations); i++ {
		x := len(elevations[i])
		for j := 0; j < len(elevations[i]); j++ {
			point := elevations[i][j]
			previous := 0
			next := 0
			above := 0
			below := 0
			if j == 0 {
				previous = elevations[i][j] + 1
			} else {
				previous = elevations[i][j-1]
			}
			if j == x-1 {
				next = elevations[i][j] + 1
			} else {
				next = elevations[i][j+1]
			}
			if i == 0 {
				above = elevations[i][j] + 1
			} else {
				above = elevations[i-1][j]
			}
			if i == y-1 {
				below = elevations[i][j] + 1
			} else {
				below = elevations[i+1][j]
			}
			if point < previous && point < next && point < above && point < below {
				lowPoints = append(lowPoints, lowPoint{point, coordinate{i, j}})
			}
		}
	}
	return lowPoints
}

func findNeighbors(elevations [][]int, c coordinate) []coordinate {
	var offsets = [4]coordinate{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	var neighbors = []coordinate{}
	for _, offset := range offsets {
		pendingCoordinate := coordinate{c.X + offset.X, c.Y + offset.Y}
		if isValidPoint(elevations, pendingCoordinate) {
			neighbors = append(neighbors, pendingCoordinate)
		}
	}
	return neighbors
}

func getBasinForPoint(elevations [][]int, lp lowPoint) basin {
	b := make(basin)
	findBasin(lp, elevations, b)
	return b
}

func getHeight(e [][]int, c coordinate) int {
	return e[c.X][c.Y]
}

func isValidPoint(m [][]int, c coordinate) bool {
	goodX := 0 <= c.X && c.X < len(m)
	goodY := 0 <= c.Y && c.Y < len(m[0])
	return goodX && goodY
}
