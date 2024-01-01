package main

// stack
// time complexity: O(n)
// space complexity: O(n)
func removeStars(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}
