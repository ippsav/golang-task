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
	fmt.Printf("Average number is : %f\n", average)
	story, _ := wholeStory(customString)
	fmt.Printf("Whole story: %s\n", story)
	storyStats(customString)
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

/*
	Estimate: 5min
	Func: whole	// for _, story := range storyParts {
	// }Story
	Param:
		- input: string
	Result: returns the average number from the given string
*/

func wholeStory(input string) (string, error) {
	if !testValidity(input) {
		return "", fmt.Errorf("the given input %s does not meet the given rules", input)
	}
	subStrings := strings.Split(input, "-")
	parsedStrings := []string{}
	for _, subString := range subStrings {
		_, err := strconv.Atoi(subString)
		if err == nil {
			continue
		}
		parsedStrings = append(parsedStrings, subString)
	}
	return strings.Join(parsedStrings, " "), nil
}

/*
	Estimate: 10min
	Func: storyStats
	Param:
		- input: string
	Result:
  		 -the shortest word
  		 -the longest word
  		 -the average word length
  		 -the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
*/

func storyStats(input string) (string, string, int, []string, error) {
	story, err := wholeStory(input)
	if err != nil {
		return "", "", -1, nil, err
	}
	storyParts := strings.Split(story, " ")
	var shortestWord, longestWord string
	var averageWordLength float32
	wordsList := []string{}
	shortestWord, longestWord = storyParts[0], storyParts[0]
	sumWordLength := 0
	for _, word := range storyParts {
		wordLength := len(word)
		sumWordLength += wordLength
		if len(shortestWord) > wordLength {
			shortestWord = story
		}
		if len(longestWord) < wordLength {
			longestWord = story
		}
	}

	// placeholder for return statement
	return "", "", -1, nil, err
}
