package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDietPlanPerformance(t *testing.T) {
	tests := []struct {
		name     string
		calories []int
		k        int
		lower    int
		upper    int
		output   int
	}{
		{"example1", []int{1, 2, 3, 4, 5}, 1, 3, 3, 0},
		{"example2", []int{3, 2}, 2, 0, 1, 1},
		{"example3", []int{6, 5, 0, 0}, 2, 1, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := dietPlanPerformance(tt.calories, tt.k, tt.lower, tt.upper)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := dietPlanPerformance2(tt.calories, tt.k, tt.lower, tt.upper)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := dietPlanPerformance3(tt.calories, tt.k, tt.lower, tt.upper)
			assert.Equal(t, tt.output, res)
		})
	}
}
