package day10

import (
	"sort"

	"smashedtoatoms.com/aoc2021/internal/shared"
)

// Run to find corrupt and incomplete navigation subsystem routines.
// Then identify and fix them and score appropriately
func Run(path string) (int, int) {
	lines := shared.GetLines(path, "inputs/day10.txt")

	return score1(lines), score2(lines)
}

func score1(lines []string) int {
	chars := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	points := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	score := 0

	for _, line := range lines {
		stack := make([]rune, 0)
	lineLoop:
		for _, char := range line {
			switch char {
			case '(', '[', '{', '<':
				stack = append(stack, chars[char])
			case ')', ']', '}', '>':
				if len(stack) == 0 || stack[len(stack)-1] != char {
					score += points[char]
					break lineLoop
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}

	return score
}

func score2(lines []string) int {
	chars := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	points := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	scores := make([]int, 0)

	for _, line := range lines {
		shouldScore := true
		score := 0
		stack := make([]rune, 0)
	lineLoop:
		for _, char := range line {
			switch char {
			case '(', '[', '{', '<':
				stack = append(stack, chars[char])
			case ')', ']', '}', '>':
				if len(stack) > 0 && stack[len(stack)-1] == char {
					stack = stack[:len(stack)-1]
				} else {
					shouldScore = false
					break lineLoop
				}
			}
		}
		if shouldScore {
			for i := len(stack) - 1; i >= 0; i-- {
				score = score*5 + points[stack[i]]
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}
