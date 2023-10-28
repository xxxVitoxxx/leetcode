package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		res   []byte
	}{
		{"example1", []byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{"example2", []byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		{"example3", []byte{'t', 'a', 'i', 'w', 'a', 'n'}, []byte{'n', 'a', 'w', 'i', 'a', 't'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.input)
			assert.True(t, reflect.DeepEqual(tt.input, tt.res))
		})
	}
}

func TestReverseString2(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		res   []byte
	}{
		{"example1", []byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{"example2", []byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		{"example3", []byte{'t', 'a', 'i', 'w', 'a', 'n'}, []byte{'n', 'a', 'w', 'i', 'a', 't'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString2(tt.input)
			assert.True(t, reflect.DeepEqual(tt.input, tt.res))
		})
	}
}

func TestReverseString3(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		res   []byte
	}{
		{"example1", []byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{"example2", []byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
		{"example3", []byte{'t', 'a', 'i', 'w', 'a', 'n'}, []byte{'n', 'a', 'w', 'i', 'a', 't'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString3(tt.input)
			assert.True(t, reflect.DeepEqual(tt.input, tt.res))
		})
	}
}
