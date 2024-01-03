package main

func removeDuplicates(s string) string {
	bytes := []byte(s)
	stack := make([]byte, 0, len(s))
	for _, b := range bytes {
		l := len(stack)
		if l > 0 && stack[l-1] == b {
			stack = stack[:l-1]
		} else {
			stack = append(stack, b)
		}
	}
	return string(stack)
}

// use recursion
func removeDuplicates2(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return removeDuplicates2(s[:i] + s[i+2:])
		}
	}
	return s
}
