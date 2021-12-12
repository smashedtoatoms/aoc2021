package main

import (
	"fmt"
	"os"

	"smashedtoatoms.com/aoc2021/internal/day10"
	"smashedtoatoms.com/aoc2021/internal/day11"
	"smashedtoatoms.com/aoc2021/internal/day12"
	"smashedtoatoms.com/aoc2021/internal/day4"
	"smashedtoatoms.com/aoc2021/internal/day5"
	"smashedtoatoms.com/aoc2021/internal/day6"
	"smashedtoatoms.com/aoc2021/internal/day7"
	"smashedtoatoms.com/aoc2021/internal/day8"
	"smashedtoatoms.com/aoc2021/internal/day9"
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
		fmt.Println(day4.Run(path))
	case "5":
		fmt.Println(day5.Run(path))
	case "6":
		fmt.Println(day6.Run(path))
	case "7":
		fmt.Println(day7.Run(path))
	case "8":
		fmt.Println(day8.Run(path))
	case "9":
		fmt.Println(day9.Run(path))
	case "10":
		fmt.Println(day10.Run(path))
	case "11":
		fmt.Println(day11.Run(path))
	case "12":
		fmt.Println(day12.Run(path))
	default:
		fmt.Println("No implementation for day specified")
	}
}
