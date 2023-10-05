package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	output := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz",
	}

	tests := []struct {
		name   string
		input  int
		output []string
	}{
		{"example1", 3, output[:3]},
		{"example2", 5, output[:5]},
		{"example3", 15, output},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := fizzBuzz(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := fizzBuzz2(tt.input)
			assert.Equal(t, tt.output, res)
		})
	}
}
