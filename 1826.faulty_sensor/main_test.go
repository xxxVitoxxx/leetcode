package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasSensor(t *testing.T) {
	tests := []struct {
		name    string
		sensor1 []int
		sensor2 []int
		output  int
	}{
		{"example1", []int{2, 3, 4, 5}, []int{2, 1, 3, 4}, 1},
		{"example2", []int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 5}, -1},
		{"example3", []int{2, 3, 2, 2, 3, 2}, []int{2, 3, 2, 3, 2, 7}, 2},
		{"example4", []int{8, 2, 2, 6, 3, 8, 7, 2, 5, 3}, []int{2, 8, 2, 2, 6, 3, 8, 7, 2, 5}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := badSensor(tt.sensor1, tt.sensor2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := badSensor2(tt.sensor1, tt.sensor2)
			assert.Equal(t, tt.output, res)
		})
	}
}
