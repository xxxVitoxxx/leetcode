package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		name     string
		chars    []byte
		output   int
		newChars []byte
	}{
		{"example1", []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6, []byte{'a', '2', 'b', '2', 'c', '3'}},
		{"example2", []byte{'a'}, 1, []byte{'a'}},
		{"exampel3", []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4, []byte{'a', 'b', '1', '2'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chars := make([]byte, len(tt.chars))
			copy(chars, tt.chars)
			res := compress(chars)
			assert.Equal(t, tt.output, res)
			assert.True(t, reflect.DeepEqual(tt.newChars, chars[:tt.output]))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chars := make([]byte, len(tt.chars))
			copy(chars, tt.chars)
			res := compress2(chars)
			assert.Equal(t, tt.output, res)
			assert.True(t, reflect.DeepEqual(tt.newChars, chars[:tt.output]))
		})
	}
}
