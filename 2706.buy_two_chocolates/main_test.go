package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyChoco(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		money  int
		output int
	}{
		{"example1", []int{1, 2, 2}, 3, 0},
		{"example2", []int{3, 2, 3}, 3, 3},
		{"example1", []int{98, 54, 6, 34, 66, 63, 52, 39}, 62, 22},
		{"example1", []int{4, 9, 8, 19, 4, 5}, 10, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := buyChoco(tt.prices, tt.money)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := buyChoco2(tt.prices, tt.money)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := buyChoco3(tt.prices, tt.money)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := buyChoco5(tt.prices, tt.money)
			assert.Equal(t, tt.output, res)
		})
	}
}

func TestBuyChoco4(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		money  int
		output int
	}{
		{"example1", []int{1, 2, 2}, 3, 0},
		{"example2", []int{3, 2, 3}, 3, 3},
		{"example1", []int{98, 54, 6, 34, 66, 63, 52, 39}, 62, 22},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := buyChoco4(tt.prices, tt.money)
			assert.Equal(t, tt.output, res)
		})
	}
}
