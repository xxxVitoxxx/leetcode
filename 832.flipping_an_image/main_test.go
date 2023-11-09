package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlipAndInvertImage(t *testing.T) {
	tests := []struct {
		name   string
		image  [][]int
		output [][]int
	}{
		{"example1", [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{"example2", [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
		{"example3", [][]int{{1}}, [][]int{{0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := flipAndInvertImage(tt.image)
			for i := 0; i < len(tt.output); i++ {
				for j := 0; j < len(tt.output[i]); j++ {
					assert.Equal(t, tt.output[i][j], res[i][j])
				}
			}
		})
	}
}
func TestFlipAndInvertImage2(t *testing.T) {
	tests := []struct {
		name   string
		image  [][]int
		output [][]int
	}{
		{"example1", [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{"example2", [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
		{"example3", [][]int{{1}}, [][]int{{0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := flipAndInvertImage2(tt.image)
			for i := 0; i < len(tt.output); i++ {
				for j := 0; j < len(tt.output[i]); j++ {
					assert.Equal(t, tt.output[i][j], res[i][j])
				}
			}
		})
	}
}

func TestFlipAndInvertImage3(t *testing.T) {
	tests := []struct {
		name   string
		image  [][]int
		output [][]int
	}{
		{"example1", [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{"example2", [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
		{"example3", [][]int{{1}}, [][]int{{0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := flipAndInvertImage3(tt.image)
			for i := 0; i < len(tt.output); i++ {
				for j := 0; j < len(tt.output[i]); j++ {
					assert.Equal(t, tt.output[i][j], res[i][j])
				}
			}
		})
	}
}

func TestFlipAndInvertImage4(t *testing.T) {
	tests := []struct {
		name   string
		image  [][]int
		output [][]int
	}{
		{"example1", [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{"example2", [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
		{"example3", [][]int{{1}}, [][]int{{0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := flipAndInvertImage4(tt.image)
			for i := 0; i < len(tt.output); i++ {
				for j := 0; j < len(tt.output[i]); j++ {
					assert.Equal(t, tt.output[i][j], res[i][j])
				}
			}
		})
	}
}
