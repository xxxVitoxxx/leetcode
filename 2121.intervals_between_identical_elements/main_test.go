package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDistances(t *testing.T) {
	tests := []struct {
		name   string
		arr    []int
		output []int64
	}{
		{"example1", []int{2, 1, 3, 1, 2, 3, 3}, []int64{4, 2, 7, 2, 4, 4, 5}},
		{"example1", []int{10, 5, 10, 10}, []int64{5, 0, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getDistances(tt.arr)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getDistances2(tt.arr)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getDistances3(tt.arr)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}
