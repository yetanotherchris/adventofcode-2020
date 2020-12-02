package main

import (
	"adventofcode/day1"
	"fmt"
	"os"
)

// Folder structure how to:
// - https://stackoverflow.com/a/61793820
// - create a folder, use go mod init some app
// - use a folder for each package.
// - add a go file to this folder
// - use package "foldername"
// - make all exported functions pascal case

func main() {

	switch os.Args[1] {
	case "day1":
		{
			fmt.Println("Advent of Code - Day 1")
			result := day1.Run()
			fmt.Println(result)
		}
	default:
		{
			fmt.Println("Unrecognised day")
		}
	}
}
