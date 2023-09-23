package main

func isValid(s string) bool {
	m := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	stack := []rune{}
	for _, r := range s {
		if val, ok := m[r]; ok {
			stack = append(stack, val)
			continue
		}

		l := len(stack)
		if l == 0 {
			return false
		} else if stack[l-1] != r {
			return false
		} else {
			stack = stack[:l-1]
		}
	}
	return len(stack) == 0
}
