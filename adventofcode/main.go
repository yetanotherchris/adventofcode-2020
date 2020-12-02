package main

import (
	"adventofcode/day1"
	"adventofcode/day2"
	"fmt"
	"os"
)

// Folder structure how to:
// - https://stackoverflow.com/a/61793820
// - create a folder, use go mod init someapp
// - use a folder for each package.
// - add a go file to this folder
// - use package "foldername"
// - make all exported functions pascal case
// - use "go build ." to build everything, ditto for run

func main() {

	switch os.Args[1] {
	case "day1":
		{
			fmt.Println("Advent of Code - Day 1")
			result := day1.Run()
			fmt.Println(result)
		}
	case "day2":
		{
			fmt.Println("Advent of Code - Day 2")
			result := day2.Run()
			fmt.Println(result)
		}
	default:
		{
			fmt.Println("Unrecognised day")
		}
	}
}
