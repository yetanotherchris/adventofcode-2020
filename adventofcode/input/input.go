package input

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFileToArray(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil
	}

	allText := string(data)
	lines := strings.Split(allText, "\n")

	return lines
}
