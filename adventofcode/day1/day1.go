package day1

import (
	"fmt"
	"strconv"
	"strings"
	"adventofcode/input"
)

// Functions must start with an uppercase letter to be exported
func Run() string {

	lines := input.ReadFileToArray("day1/input.txt")
	slice := convertToIntArray(lines)

	for i := 0; i < len(slice); i++ {
		currentItem := slice[i]

		var hasMatch, value1, value2, multipliedSum = checkBeforeCurrentItem(i, currentItem, slice)
		if hasMatch {
			return fmt.Sprintf("It's %d * %d = %d", value1, value2, multipliedSum)
		}

		hasMatch, value1, value2, multipliedSum = checkAfterCurrentItem(i, currentItem, slice)
		if hasMatch {
			return fmt.Sprintf("It's %d * %d = %d", value1, value2, multipliedSum)

		}
	}

	return "Couldn't find the answer"
}

func convertToIntArray(lines []string) []int {
	slice := make([]int, 0)
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