package main

// rotated num and compare it with num
func isStrobogrammatic(num string) bool {
	var rotated []rune
	for _, r := range num {
		switch r {
		case '2', '3', '4', '5', '7':
			return false
		case '6':
			rotated = append([]rune{'9'}, rotated...)
		case '9':
			rotated = append([]rune{'6'}, rotated...)
		default:
			rotated = append([]rune{r}, rotated...)
		}
	}
	return num == string(rotated)
}

// use hash table and two pointers
func isStrobogrammatic2(num string) bool {
	m := map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}
	for left, right := 0, len(num)-1; left <= right; {
		lValue, ok := m[num[left]]
		if !ok {
			return false
		}

		rValue, ok := m[num[right]]
		if !ok {
			return false
		}

		if lValue != num[right] || rValue != num[left] {
			return false
		}

		left, right = left+1, right-1
	}
	return true
}

// use two pointers
func isStrobogrammatic3(num string) bool {
	for left, right := 0, len(num)-1; left <= right; {
		if !check(num[left], num[right]) {
			return false
		}
		left, right = left+1, right-1
	}
	return true
}

func check(b1, b2 byte) bool {
	switch b1 {
	case '0', '1', '8':
		return b1 == b2
	case '6':
		return b2 == '9'
	case '9':
		return b2 == '6'
	default:
		return false
	}
}
