package day11

import (
	"smashedtoatoms.com/aoc2021/internal/shared"
)

type coordinate struct {
	X int
	Y int
}

// Run counts the octopi flashes in the given matrix and also calculates
// when they all simultaneously flash.
func Run(path string) (int, int) {
	matrix, err := shared.SplitToInts(shared.GetLines(path, "inputs/day11.txt"))
	if err != nil {
		panic(err)
	}

	return countFlashingOctopi(matrix, 100), findSimultaneousFlashOfOctopi(matrix)
}

func countFlashingOctopi(matrix [][]int, count int) int {
	mtx := make([][]int, len(matrix))
	copy(mtx, matrix)
	acc := 0
	for i := 0; i < count; i++ {
		acc += countFlashingOctopiOnce(mtx)
	}
	return acc
}

func findSimultaneousFlashOfOctopi(matrix [][]int) int {
	acc := 0
	for {
		acc += 1
		if countFlashingOctopiOnce(matrix) == len(matrix)*len(matrix[0]) {
			break
		}
	}
	return acc
}

func countFlashingOctopiOnce(mtx [][]int) int {
	acc := 0
	for y, row := range mtx {
		for x := range row {
			if mtx[y][x] == 9 {
				mtx[y][x] = -1
			} else {
				mtx[y][x] = mtx[y][x] + 1
			}
		}
	}
	for y, row := range mtx {
		for x := range row {
			if mtx[y][x] == -1 {
				flashNeighbors(mtx, coordinate{x, y}, &acc)
			}
		}
	}
	return acc
}

func flashNeighbors(mtx [][]int, c coordinate, acc *int) {
	if acc != nil {
		*acc += 1
	}
	var offsets = [8]coordinate{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}
	var neighbors = []coordinate{}
	for _, offset := range offsets {
		pendingCoordinate := coordinate{c.X + offset.X, c.Y + offset.Y}
		if isValidPoint(mtx, pendingCoordinate) {
			neighbors = append(neighbors, pendingCoordinate)
		}
	}
	for _, neighbor := range neighbors {
		if mtx[neighbor.Y][neighbor.X] == 9 {
			mtx[neighbor.Y][neighbor.X] = -1
			flashNeighbors(mtx, neighbor, acc)
		} else if mtx[neighbor.Y][neighbor.X] > 0 {
			mtx[neighbor.Y][neighbor.X] = mtx[neighbor.Y][neighbor.X] + 1
		}
	}
	mtx[c.Y][c.X] = 0
}

func isValidPoint(m [][]int, c coordinate) bool {
	goodX := 0 <= c.X && c.X < len(m)
	goodY := 0 <= c.Y && c.Y < len(m[0])
	return goodX && goodY
}
