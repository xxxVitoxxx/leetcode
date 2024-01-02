package main

// stack
// time complexity: O(n)
// space complexity: O(n)
func decodeString(s string) string {
	var stack []rune
	for _, r := range s {
		if r != ']' {
			stack = append(stack, r)
		} else {
			var encoded string
			for stack[len(stack)-1] != '[' {
				encoded = string(stack[len(stack)-1]) + encoded
				stack = stack[:len(stack)-1]
			}
			// pop the '[' from the stack
			stack = stack[:len(stack)-1]

			num, base := 0, 1
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				num = int(stack[len(stack)-1]-'0')*base + num
				base *= 10
				stack = stack[:len(stack)-1]
			}

			var str string
			for n := 1; n <= num; n++ {
				str += encoded
			}
			stack = append(stack, []rune(str)...)
		}
	}

	var ans string
	for i := 0; i < len(stack); i++ {
		ans += string(stack[i])
	}
	return ans
}
