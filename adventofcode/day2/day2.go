package day2

import (
	"adventofcode/input"
	"fmt"
	"regexp"
	"strconv"
)

type Password struct {
	minimumChars int
	maximumChars int
	character    string
	value        string
}

func Run() string {
	passwords := parsePasswords()
	validPasswordCount := 0

	for _, password := range passwords {

		if isValidPassword(password){
			validPasswordCount++
		}
	}

	return fmt.Sprintf("%d valid passwords found", validPasswordCount)
}

func parsePasswords() []Password {
	lines := input.ReadFileToArray("day2/input.txt")
	var allPasswords = make([]Password, 0)

	var ruleAndValue = regexp.MustCompile(`^([0-9]+)-([0-9]+)\s*([a-z]+):\s*([a-z]+\s*)$`)
	for _, item := range lines {
		match := ruleAndValue.FindStringSubmatch(item)

		if len(match) == 5 {
			minChars, _ := strconv.Atoi(match[1])
			maxChars, _ := strconv.Atoi(match[2])
			password := Password {
				minimumChars: minChars,
				maximumChars: maxChars,
				character: match[3],
				value: match[4] }

			allPasswords = append(allPasswords, password)
		}
	}

	return allPasswords
}

func isValidPassword(password Password) bool {
	characterCount := 0
	for i := 0; i < len(password.value); i++ {
		if string(password.value[i]) == password.character {
			characterCount++
		}
	}

	return characterCount <= password.maximumChars &&
		characterCount >= password.minimumChars
}