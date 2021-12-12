package day12

import (
	"unicode"

	"smashedtoatoms.com/aoc2021/internal/shared"
)

// Calculates the paths through the caves.
func Run(path string) (int, int) {
	caves := shared.SplitToStrings(shared.GetLines(path, "inputs/day12.txt"), "-")

	return findAllUndergroundPathsVisitOnce(caves), findAllUndergroundPathsVisitTwice(caves)
}

func findAllUndergroundPathsVisitOnce(caves [][]string) int {
	exits := map[string][]string{}
	visits := map[string]int{}

	for _, cave := range caves {
		exits[string(cave[0])] = append(exits[string(cave[0])], string(cave[1]))
		exits[string(cave[1])] = append(exits[string(cave[1])], string(cave[0]))

	}
	return getAllPossiblePathsVisitOnce(exits, "start", visits)
}

func getAllPossiblePathsVisitOnce(exits map[string][]string, currentCave string, visits map[string]int) int {
	visits[currentCave] += 1
	sum := 0
	for _, e := range exits[currentCave] {
		if e == "end" {
			sum += 1
			continue
		}
		if !unicode.IsLower(rune(e[0])) || visits[e] == 0 {
			nextVisits := visits
			if unicode.IsLower(rune(e[0])) {
				nextVisits = shared.CopyStringIntMap(visits)
			}
			sum += getAllPossiblePathsVisitOnce(exits, e, nextVisits)
		}
	}
	return sum
}

func findAllUndergroundPathsVisitTwice(caves [][]string) int {
	exits := map[string][]string{}
	visits := map[string]int{}

	for _, cave := range caves {
		exits[string(cave[0])] = append(exits[string(cave[0])], string(cave[1]))
		exits[string(cave[1])] = append(exits[string(cave[1])], string(cave[0]))
	}
	return getAllPossiblePathsWithSmallDoubleVisit(exits, "start", visits, false)
}

func getAllPossiblePathsWithSmallDoubleVisit(
	exits map[string][]string,
	currentCave string,
	visits map[string]int,
	hasVisitedTwice bool,
) int {
	if visits[currentCave] > 0 && unicode.IsLower(rune(currentCave[0])) {
		hasVisitedTwice = true
	}
	visits[currentCave] += 1
	sum := 0
	for _, c := range exits[currentCave] {
		if c == "end" {
			sum += 1
			continue
		}
		if c != "start" && (!unicode.IsLower(rune(c[0])) || visits[c] < 1 || !hasVisitedTwice) {
			nextVisits := visits
			if unicode.IsLower(rune(c[0])) {
				nextVisits = shared.CopyStringIntMap(visits)
			}
			sum += getAllPossiblePathsWithSmallDoubleVisit(exits, c, nextVisits, hasVisitedTwice)
		}
	}
	return sum
}
