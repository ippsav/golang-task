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
			if out := TestValidity(tc.input); out != tc.ok {
				tr.Errorf("validate failed. Expected %v, got: %v", tc.ok, out)
			}
		})
	}
}
