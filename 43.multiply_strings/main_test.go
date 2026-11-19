package main

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		num1     string
		num2     string
		expected string
	}{{
		name:     "example1",
		num1:     "0",
		num2:     "0",
		expected: "0",
	}, {
		name:     "example2",
		num1:     "123",
		num2:     "456",
		expected: "56088",
	}, {
		name:     "example3",
		num1:     "2",
		num2:     "3",
		expected: "6",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := multiply(tt.num1, tt.num2)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
