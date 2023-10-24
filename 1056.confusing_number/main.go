package main

import (
	"bytes"
	"strconv"
)

func confusingNumber(n int) bool {
	var rotate int
	for num := n; num > 0; num /= 10 {
		tmp := check(num % 10)
		if tmp < 0 {
			return false
		}
		rotate = rotate*10 + tmp
	}
	return n != rotate
}

func check(n int) int {
	switch n {
	case 0, 1, 8:
		return n
	case 6:
		return 9
	case 9:
		return 6
	default:
		return -1
	}
}

func confusingNumber2(n int) bool {
	str := strconv.Itoa(n)
	m := map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}

	var buf bytes.Buffer
	for i := len(str) - 1; i >= 0; i-- {
		if v, ok := m[str[i]]; ok {
			buf.WriteByte(v)
		} else {
			return false
		}
	}
	return str != buf.String()
}
