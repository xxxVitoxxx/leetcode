package main

func makeGood(s string) string {
	if len(s) < 2 {
		return s
	}

	var stack []rune
	for _, r := range s {
		if len(stack) == 0 {
			stack = append(stack, r)
			continue
		}

		if stack[len(stack)-1] == r+32 || stack[len(stack)-1] == r-32 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}

func makeGood2(s string) string {
	if len(s) < 2 {
		return s
	}

	var bytes []byte
	for i := range s {
		l := len(bytes)
		if l > 0 && (bytes[l-1] == s[i]+32 || bytes[l-1] == s[i]-32) {
			bytes = bytes[:l-1]
		} else {
			bytes = append(bytes, s[i])
		}
	}
	return string(bytes)
}

// use recursion
func makeGood3(s string) string {
	if len(s) < 2 {
		return s
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1]+32 || s[i] == s[i+1]-32 {
			return makeGood3(s[:i] + s[i+2:])
		}
	}
	return s
}
