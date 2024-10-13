package main

import (
	"strconv"
	"testing"
)

func TestReverseBits(t *testing.T) {
	for _, tt := range []struct {
		binaryStr string
		want      string
	}{{
		binaryStr: "00000010100101000001111010011100",
		want:      "00111001011110000010100101000000",
	}} {
		num, _ := strconv.ParseInt(tt.binaryStr, 2, 32)
		want, _ := strconv.ParseInt(tt.want, 2, 32)

		num1 := num
		got := reverseBits(uint32(num1))
		if got != uint32(want) {
			t.Errorf("got: %32b, want: %32b", got, want)
		}

		num2 := num
		got = reverseBits2(uint32(num2))
		if got != uint32(want) {
			t.Errorf("got: %32b, want: %32b", got, want)
		}
	}
}
