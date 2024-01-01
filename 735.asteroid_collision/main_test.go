package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		name      string
		asteroids []int
		output    []int
	}{
		{"example1", []int{5, 10, -5}, []int{5, 10}},
		{"example2", []int{8, -8}, []int{}},
		{"example3", []int{10, 2, -5}, []int{10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := asteroidCollision(tt.asteroids)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := asteroidCollision2(tt.asteroids)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}
