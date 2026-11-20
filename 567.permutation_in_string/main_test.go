package main

import "testing"

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{{
		name:     "example1",
		s1:       "ab",
		s2:       "eidbaooo",
		expected: true,
	}, {
		name:     "example2",
		s1:       "ab",
		s2:       "eidboaoo",
		expected: false,
	}, {
		name:     "example3",
		s1:       "adc",
		s2:       "dcda",
		expected: true,
	}, {
		name:     "example4",
		s1:       "horse",
		s2:       "ros",
		expected: false,
	}, {
		name:     "example5",
		s1:       "ab",
		s2:       "a",
		expected: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInclusion(tt.s1, tt.s2)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInclusion2(tt.s1, tt.s2)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
