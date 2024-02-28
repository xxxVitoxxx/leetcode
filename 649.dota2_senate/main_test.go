package main

import "testing"

func TestPredictPartyVictory(t *testing.T) {
	tests := []struct {
		name   string
		senate string
		output string
	}{
		{"example1", "RD", "Radiant"},
		{"example2", "RDD", "Dire"},
		{"example3", "D", "Dire"},
		{"example4", "RDR", "Radiant"},
		{"example5", "DDRRR", "Dire"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := predictPartyVictory(tt.senate)
			if res != tt.output {
				t.Errorf("got %v, want %v", res, tt.output)
			}
		})
	}
}
