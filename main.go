package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	customString := "23-ab-48-caba-41-haha"
	fmt.Printf("%t", testValidity(customString))
}

/*
	Estimate: 5 min
	Func: testValidity
	Param:
		- input: string
	Result: returns boolean if the string follows rules given in the task markdown
*/
func testValidity(input string) bool {
	if len(input) == 0 {
		return false
	}
	subStrings := strings.Split(input, "-")
	isStringAlphabetic := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	parseSwitch := false
	for _, subString := range subStrings {
		_, err := strconv.Atoi(subString)
		if !parseSwitch && err != nil {
			return false
		} else if parseSwitch && err != nil {
			if !isStringAlphabetic(subString) {
				return false
			}
		}
		parseSwitch = !parseSwitch
	}
	return true
}
