package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 1")

	lines := readFileToArray("input.txt")
	slice := convertToIntArray(lines)

	for i := 0; i < len(slice); i++ {
		currentItem := slice[i]

		var hasMatch, value1, value2, multipliedSum = checkBeforeCurrentItem(i, currentItem, slice)
		if hasMatch {
			fmt.Printf("It's %d * %d = %d\n", value1, value2, multipliedSum)
			return
		}

		hasMatch, value1, value2, multipliedSum = checkAfterCurrentItem(i, currentItem, slice)
		if hasMatch {
			fmt.Printf("It's %d * %d = %d\n", value1, value2, multipliedSum)
			return
		}
	}
}

func readFileToArray(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil
	}

	allText := string(data)
	lines := strings.Split(allText, "\n")

	return lines
}

func convertToIntArray(lines []string) []int {
	slice := make([]int, len(lines))
	for _, item := range lines {
		cleanedLine := strings.Replace(item, "\r", "", 1)
		value, _ := strconv.Atoi(cleanedLine)
		slice = append(slice, value)
	}

	return slice
}

func checkBeforeCurrentItem(currentIndex int, currentItem int, slice []int) (bool, int, int, int){
	for before := 0; before < currentIndex; before++ {
		var sum = slice[before] + currentItem

		if sum == 2020 {
			multiplied := slice[before] * currentItem
			return true, slice[before], currentItem, multiplied
		}
	}

	return false,0,0,0
}

func checkAfterCurrentItem(currentIndex int, currentItem int, slice []int) (bool, int, int, int){
	var sum int
	for after := currentIndex + 1; after < len(slice); after++ {
		sum = slice[after] + currentItem

		if sum == 2020 {
			multiplied := slice[after] * currentItem
			return true, slice[after], currentItem, multiplied
		}
	}

	return false,0,0,0
}