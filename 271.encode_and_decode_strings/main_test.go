package main

import (
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{{
		name: "example1",
		strs: []string{"Hello", "World"},
	}, {
		name: "example2",
		strs: []string{""},
	}, {
		name: "example3",
		strs: []string{"63/Rc", "h", "BmI3FS~J9#vmk", "7uBZ?7*/", "24h+X", "O "},
	}}

	codec := &Codec{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodeResult := codec.Encode(tt.strs)
			decodeResult := codec.Decode(encodeResult)

			if len(decodeResult) != len(tt.strs) {
				t.Error("length not match")
			}

			var noAllEqual bool
			for i := range len(decodeResult) {
				if decodeResult[i] != tt.strs[i] {
					noAllEqual = true
				}
			}
			if noAllEqual {
				t.Errorf("got: %v, want: %v", decodeResult, tt.strs)
			}
		})
	}
}
