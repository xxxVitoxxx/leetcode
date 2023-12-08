package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name   string
		s      []byte
		answer []byte
	}{
		{"example1", []byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'}, []byte{'b', 'l', 'u', 'e', ' ', 'i', 's', ' ', 's', 'k', 'y', ' ', 't', 'h', 'e'}},
		{"example2", []byte{'a'}, []byte{'a'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseWords(tt.s)
			assert.True(t, reflect.DeepEqual(tt.answer, tt.s))
		})
	}
}

func TestReverseWords2(t *testing.T) {
	tests := []struct {
		name   string
		s      []byte
		answer []byte
	}{
		{"example1", []byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'}, []byte{'b', 'l', 'u', 'e', ' ', 'i', 's', ' ', 's', 'k', 'y', ' ', 't', 'h', 'e'}},
		{"example2", []byte{'a'}, []byte{'a'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseWords2(tt.s)
			assert.True(t, reflect.DeepEqual(tt.answer, tt.s))
		})
	}
}

func TestReverseWords23(t *testing.T) {
	tests := []struct {
		name   string
		s      []byte
		answer []byte
	}{
		{"example1", []byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'}, []byte{'b', 'l', 'u', 'e', ' ', 'i', 's', ' ', 's', 'k', 'y', ' ', 't', 'h', 'e'}},
		{"example2", []byte{'a'}, []byte{'a'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseWords3(tt.s)
			assert.True(t, reflect.DeepEqual(tt.answer, tt.s))
		})
	}
}
