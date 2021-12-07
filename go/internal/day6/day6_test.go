package day6

import (
	"testing"
)

func Test1(t *testing.T) {
	in := []int{3, 4, 3, 1, 2}
	days := 80
	out := countLanternfish(in, days)
	expected := 5934
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}

func Test2(t *testing.T) {
	in := []int{3, 4, 3, 1, 2}
	days := 256
	out := countLanternfish(in, days)
	expected := 26984457539
	if out != expected {
		t.Errorf("wrong result: got %d, expected %d", out, expected)
		return
	}
}
