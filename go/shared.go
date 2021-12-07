package aoc

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func getLines(path string, fallback string) []string {
	if path == "" {
		path = fallback
	}

	_, b, _, _ := runtime.Caller(0)
	fullpath := filepath.Join(filepath.Dir(b), path)
	data, _ := os.ReadFile(fullpath)
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func filterStrings(s []string, toFilter string) []string {
	var r []string
	for _, str := range s {
		if str != toFilter {
			r = append(r, str)
		}
	}
	return r
}

func stringsToInts(ss []string) ([]int, error) {
	var out []int
	for _, v := range ss {
		n, err := strconv.Atoi(v)
		if err != nil {
			return out, err
		}
		out = append(out, n)
	}
	return out, nil
}

func min(xs []int) int {
	var y int
	for _, x := range xs {
		if x < y || y == 0 {
			y = x
		}
	}
	return y
}

func max(xs []int) int {
	var y int
	for _, x := range xs {
		if x > y {
			y = x
		}
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
