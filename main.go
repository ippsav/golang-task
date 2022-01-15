package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	customString := "23-ab-48-cab-41-hahaa"
	fmt.Printf("%t\n", TestValidity(customString))
	average, _ := AverageNumber(customString)
	fmt.Printf("Average number is : %f\n", average)
	story, _ := wholeStory(customString)
	fmt.Printf("Whole story: %s\n", story)
	shortestWord, longestWord, averageWordLength, list, _ := StoryStats(customString)
	fmt.Printf("Shortest word: %s\n", shortestWord)
	fmt.Printf("Longest word: %s\n", longestWord)
	fmt.Printf("Average word length: %f\n", averageWordLength)
	fmt.Printf("Word list: %s", list)
}

/*
	Estimate: 5 min
	Func: testValidity
	Param:
		- input: string
	Result: returns boolean if the string follows rules given in the task markdown
*/
func TestValidity(input string) bool {
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

func AverageNumber(input string) (float32, error) {
	if !TestValidity(input) {
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
	if !TestValidity(input) {
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

func StoryStats(input string) (string, string, float64, []string, error) {
	story, err := wholeStory(input)
	if err != nil {
		return "", "", -1, nil, err
	}
	storyParts := strings.Split(story, " ")
	var shortestWord, longestWord string
	var averageWordLength float64
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
			longestWord = word
		}
	}
	averageWordLength = float64(sumWordLength) / float64(len(storyParts))
	for _, word := range storyParts {
		roundedAverage := math.Round(float64(averageWordLength))
		if len(word) == int(roundedAverage) {
			wordsList = append(wordsList, word)
		}
	}

	return shortestWord, longestWord, averageWordLength, wordsList, nil
}
