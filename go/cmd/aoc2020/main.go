package main

import (
	"fmt"
	"os"

	"smashedtoatoms.com/aoc2021/internal/day4"
	"smashedtoatoms.com/aoc2021/internal/day5"
	"smashedtoatoms.com/aoc2021/internal/day6"
	"smashedtoatoms.com/aoc2021/internal/day7"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Provide number of day to run")
		os.Exit(1)
	}

	day := os.Args[1]

	var path string
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	switch day {
	case "4":
		fmt.Println(day4.RunDay4(path))
	case "5":
		fmt.Println(day5.RunDay5(path))
	case "6":
		fmt.Println(day6.RunDay6(path))
	case "7":
		fmt.Println(day7.RunDay7(path))
	default:
		fmt.Println("No implementation for day specified")
	}
}
