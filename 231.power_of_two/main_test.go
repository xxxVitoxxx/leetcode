package main

import (
	"fmt"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output bool
	}{
		{"example1", 1, true},
		{"example2", 16, true},
		{"example3", 3, false},
		{"example4", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.output != isPowerOfTwo(tt.input) {
				t.Errorf("isPowerOfTwo %s failed, got = %v, want %v", tt.name, isPowerOfTwo(tt.input), tt.output)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if tt.output != isPowerOfTwo2(tt.input) {
				t.Errorf("isPowerOfTwo2 %s failed, got = %v, want %v", tt.name, isPowerOfTwo2(tt.input), tt.output)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if tt.output != isPowerOfTwo3(tt.input) {
				t.Errorf("isPowerOfTwo3 %s failed, got = %v, want %v", tt.name, isPowerOfTwo3(tt.input), tt.output)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if tt.output != isPowerOfTwo4(tt.input) {
				t.Errorf("isPowerOfTwo4 %s failed, got = %v, want %v", tt.name, isPowerOfTwo4(tt.input), tt.output)
			}
		})
	}
}

func TestKd(t *testing.T) {
	n := 8
	// 8 -> 1000
	// 7  -> 00000111
	// -7 -> 11111000
	fmt.Printf("%b\n", -n)
	fmt.Println(n & -n)

	fmt.Printf("%08b\n", uint8(n))
	fmt.Printf("%08b\n", uint8(^n))
	fmt.Printf("%08b\n", uint8(-n))
	fmt.Printf("%08b\n", uint8(-n+1))
}

/*
n := 4
n  -> 00000100
^n -> 11111011
-n -> 11111100
n-1 -> 00000011
&

n := 8
n  -> 00001000
^n -> 11110111
-n -> 11111000
n-1 -> 00000111

n := 6
n  -> 00000110
^n -> 11111001
-n -> 11111010
n-1 -> 00000101

n := 3
n  -> 00000011
^n -> 11111100
-n -> 11111101
n-1 -> 00000010
*/
