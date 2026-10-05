package main

import "testing"

func TestAlienDictionary(t *testing.T) {
	tests := []struct {
		name   string
		words  []string
		expect string
	}{
		{"example1", []string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
		{"example2", []string{"z", "x"}, "zx"},
		{"example3", []string{"z", "x", "z"}, ""},
		{"example4", []string{"z", "x", "a", "zb", "az"}, ""},
		{"example5", []string{"z", "z"}, "z"},
		{"example6", []string{"abc", "ab"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := alienOrder(tt.words)
			if res != tt.expect {
				t.Errorf("expected: %s, got: %s", tt.expect, res)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := alienOrder2(tt.words)
			if res != tt.expect {
				t.Errorf("expected: %s, got: %s", tt.expect, res)
			}
		})
	}
}
