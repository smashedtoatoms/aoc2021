package shared

import (
	"os"
	"strconv"
	"strings"
)

func GetLines(path string, fallback string) []string {
	if path == "" {
		path = fallback
	}
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
}

func FilterStrings(s []string, toFilter string) []string {
	var r []string
	for _, str := range s {
		if str != toFilter {
			r = append(r, str)
		}
	}
	return r
}

func GetKeys(xs map[string]string) []string {
	out := make([]string, 0, len(xs))
	for k := range xs {
		out = append(out, k)
	}
	return out
}

func IntsContain(xs []int, x int) bool {
	for _, y := range xs {
		if x == y {
			return true
		}
	}
	return false
}

func StringsContain(xs []string, x string) bool {
	for _, y := range xs {
		if x == y {
			return true
		}
	}
	return false
}

func StringsToInts(ss []string) ([]int, error) {
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

func Min(xs []int) int {
	var y int
	for _, x := range xs {
		if x < y || y == 0 {
			y = x
		}
	}
	return y
}

func Max(xs []int) int {
	var y int
	for _, x := range xs {
		if x > y {
			y = x
		}
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
