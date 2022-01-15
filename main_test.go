package main

import (
	"reflect"
	"testing"
)

func TestTestValidity(t *testing.T) {
	tt := []struct {
		name  string
		input string
		ok    bool
	}{
		{
			name:  "Should pass",
			input: "23-ab-48-cab-41-hahaa",
			ok:    true,
		},
		{
			name:  "Should not pass",
			input: "23-ab-48-41-hahaa",
			ok:    false,
		},
		{
			name:  "Should not pass",
			input: "random string",
			ok:    false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(tr *testing.T) {
			if out := testValidity(tc.input); out != tc.ok {
				tr.Errorf("validate failed. Expected %v, got: %v", tc.ok, out)
			}
		})
	}
}

func TestAverageNumber(t *testing.T) {
	tt := []struct {
		name          string
		input         string
		averageNumber float32
	}{
		{
			name:          "Should pass",
			input:         "23-ab-48-cab-41-hahaa",
			averageNumber: 37.333332,
		},
		{
			name:          "Should not pass",
			input:         "23-41-hahaa",
			averageNumber: -1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(tr *testing.T) {
			if out, _ := averageNumber(tc.input); out != tc.averageNumber {
				tr.Errorf("validate failed. Expected %v, got: %v", tc.averageNumber, out)
			}
		})
	}
}

func TestWholeStory(t *testing.T) {
	tt := []struct {
		name  string
		input string
		story string
	}{
		{
			name:  "Should pass",
			input: "23-ab-48-cab-41-hahaa",
			story: "ab cab hahaa",
		},
		{
			name:  "Should pass",
			input: "11-hello-33-world",
			story: "hello world",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(tr *testing.T) {
			if out, _ := wholeStory(tc.input); out != tc.story {
				tr.Errorf("validate failed. Expected %v, got: %v", tc.story, out)
			}
		})
	}
}

func TestStoryStats(t *testing.T) {
	tt := []struct {
		name       string
		input      string
		storyStats struct {
			shortestWord      string
			longestWord       string
			averageWordLength float64
			wordList          []string
		}
	}{
		{
			name:  "Should pass",
			input: "23-ab-48-cab-41-hahaa",
			storyStats: struct {
				shortestWord      string
				longestWord       string
				averageWordLength float64
				wordList          []string
			}{
				shortestWord:      "ab",
				longestWord:       "hahaa",
				averageWordLength: 3.3333333333333335,
				wordList:          []string{"cab"},
			},
		},
		{
			name:  "Should pass",
			input: "23-ab-48-caba-41-haha",
			storyStats: struct {
				shortestWord      string
				longestWord       string
				averageWordLength float64
				wordList          []string
			}{
				shortestWord:      "ab",
				longestWord:       "caba",
				averageWordLength: 3.3333333333333335,
				wordList:          []string{},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(tr *testing.T) {
			shortestWord, longestWord, averageWordLength, wordList, _ := storyStats(tc.input)
			storyStats := struct {
				shortestWord      string
				longestWord       string
				averageWordLength float64
				wordList          []string
			}{
				shortestWord,
				longestWord,
				averageWordLength,
				wordList,
			}
			if !reflect.DeepEqual(storyStats, tc.storyStats) {
				tr.Errorf("validate failed. Expected %v, got %v", storyStats, tc.storyStats)
			}
		})
	}
}
