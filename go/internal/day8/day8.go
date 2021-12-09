package day8

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"smashedtoatoms.com/aoc2021/internal/shared"
)

type samplesAndDisplays struct {
	Samples  []string
	Displays []string
}

// Day8 takes one list of sample strings to determine the values of display
// numbers, and then calculates the display number and sums them all.
func RunDay8(path string) (int, int) {
	lines := shared.GetLines(path, "inputs/day8.txt")
	samplesAndData := getSamplesAndDisplays(lines)
	knownCount := countKnownDisplays(samplesAndData)
	allNumberSum, err := sumAllNumbers(samplesAndData)
	if err != nil {
		panic(err)
	}

	return knownCount, allNumberSum
}

func countKnownDisplays(sds []samplesAndDisplays) int {
	count := 0
	for _, sd := range sds {
		for _, ds := range sd.Displays {
			if shared.IntsContain([]int{2, 3, 4, 7}, len(ds)) {
				count += 1
			}
		}
	}
	return count
}

func sumAllNumbers(sds []samplesAndDisplays) (int, error) {
	numbers := 0
	for _, sds := range sds {
		mappings, remaining := getKnownSampleMappings(sds)
		getSixes(&mappings, &remaining)
		for _, sample := range sds.Samples {
			if len(sample) == 5 {
				if isSubset(mappings[7], sample) {
					mappings[3] = sample
				} else if isSubset(sample, mappings[6]) {
					mappings[5] = sample
				} else {
					mappings[2] = sample
				}
			}
		}

		finalDisplay := ""
		for _, display := range sds.Displays {
			value, err := bySortedValue(mappings, display)
			if err != nil {
				return -1, errors.Wrap(err, "failed to calculate display")
			}
			finalDisplay = fmt.Sprintf("%s%d", finalDisplay, value)
		}
		final, err := strconv.Atoi(finalDisplay)
		if err != nil {
			return -1, errors.Wrap(err, "failed to convert display to number")
		}
		numbers += final
	}
	return numbers, nil
}

func getKnownSampleMappings(sds samplesAndDisplays) (map[int]string, []string) {
	mappings := make(map[int]string, 0)
	remaining := make([]string, 0)

	for _, sample := range sds.Samples {
		if len(sample) == 2 {
			mappings[1] = sample
		} else if len(sample) == 3 {
			mappings[7] = sample
		} else if len(sample) == 4 {
			mappings[4] = sample
		} else if len(sample) == 7 {
			mappings[8] = sample
		} else {
			remaining = append(remaining, sample)
		}
	}

	return mappings, remaining
}

func getSixes(mappings *map[int]string, remaining *[]string) {
	newRemaining := make([]string, 0)
	for _, s := range *remaining {
		if len(s) == 6 {
			if !isSubset((*mappings)[1], s) {
				(*mappings)[6] = s
			} else if isSubset((*mappings)[4], s) {
				(*mappings)[9] = s
			} else {
				(*mappings)[0] = s
			}
		} else {
			newRemaining = append(newRemaining, s)
		}
	}
	*remaining = newRemaining
}

func isSubset(s1 string, s2 string) bool {
	if s1 == "" || s2 == "" {
		return false
	}

	for _, c := range s1 {
		if !strings.Contains(s2, string(c)) {
			return false
		}
	}

	return true
}

func getSamplesAndDisplays(lines []string) []samplesAndDisplays {
	sds := make([]samplesAndDisplays, 0)

	for _, line := range lines {
		parts := strings.Split(line, " | ")
		if len(parts) > 1 {
			sds = append(sds, samplesAndDisplays{
				Samples:  strings.Split(parts[0], " "),
				Displays: strings.Split(parts[1], " "),
			})
		}
	}

	return sds
}

func bySortedValue(m map[int]string, value string) (int, error) {
	for k, v := range m {
		if sortString(v) == sortString(value) {
			return k, nil
		}
	}
	return -1, errors.New("value not found")
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
