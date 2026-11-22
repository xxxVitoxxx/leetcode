package main

// time complexity: O(M+N)
// space complexity: O(M+N)
func addBinary(a string, b string) string {
	result := make([]byte, max(len(a), len(b))+1)
	n1, n2 := len(a)-1, len(b)-1
	var carry int
	for i := len(result) - 1; i >= 0; i-- {
		var num int
		if n1 >= 0 {
			num += int(a[n1] - '0')
			n1--
		}

		if n2 >= 0 {
			num += int(b[n2] - '0')
			n2--
		}

		num += carry
		result[i] = byte(num%2) + '0'
		carry = num / 2
	}

	if result[0] == '0' {
		return string(result[1:])
	}
	return string(result)
}
