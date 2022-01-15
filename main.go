package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	customString := "23-ab-48-caba-41-haha"
	fmt.Printf("%t\n", testValidity(customString))
	average, _ := averageNumber(customString)
	fmt.Printf("%f\n", average)
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

/*
	Estimate: 5min
	Func: averageNumber
	Param:
		- input: string
	Result: returns the average number from the given string
*/

func averageNumber(input string) (float32, error) {
	if !testValidity(input) {
		return -1, fmt.Errorf("the given input %s does not meet the given rules", input)
	}
	subStrings := strings.Split(input, "-")
	sum := 0
	for _, subString := range subStrings {
		number, err := strconv.Atoi(subString)
		if err != nil {
			continue
		}
		sum += number
	}
	return float32(sum) / float32(len(subStrings)/2), nil
}
