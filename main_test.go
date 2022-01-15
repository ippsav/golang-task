package main

import "testing"

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
