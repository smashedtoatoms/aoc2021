package main

import (
	"fmt"
	"os"

	aoc "smashedtoatoms.com/aoc2021"
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
		fmt.Println(aoc.RunDay4(path))
	case "5":
		fmt.Println(aoc.RunDay5(path))
	case "6":
		fmt.Println(aoc.RunDay6(path))
	case "7":
		fmt.Println(aoc.RunDay7(path))
	default:
		fmt.Println("No implementation for day specified")
	}
}
