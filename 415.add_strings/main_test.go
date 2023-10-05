package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStrings(t *testing.T) {
	carr := rune(0) + rune('5'-'0')
	fmt.Println(carr + '0')
	fmt.Println('0')

	tests := []struct {
		name   string
		num1   string
		num2   string
		output string
	}{
		{"example1", "11", "123", "134"},
		{"example2", "9", "9", "18"},
		{"example3", "3876620623801494171", "6529364523802684779", "10405985147604178950"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := addStrings(tt.num1, tt.num2)
			assert.Equal(t, tt.output, res)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := addStrings2(tt.num1, tt.num2)
			assert.Equal(t, tt.output, res)
		})
	}
}
