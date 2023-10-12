package main

func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))
	for i := range s {
		l := len(stack)
		if l > 0 && stack[l-1] == s[i] {
			stack = stack[:l-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

// use recursive
func removeDuplicates2(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return removeDuplicates2(s[:i] + s[i+2:])
		}
	}
	return s
}
