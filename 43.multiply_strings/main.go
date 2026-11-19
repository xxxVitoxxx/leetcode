package main

import (
	"strings"
)

// time complexity: O(MN)
// space complexity: O(M+N)
/*
		    456
		*   123
	 ------------
		   1368
		   912
		  456
		-------
		  56088
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n1, n2 := len(num1), len(num2)
	arr := make([]int, n1+n2)

	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			num := int(num1[i]-'0') * int(num2[j]-'0')
			num += arr[i+j+1]

			arr[i+j+1] = num % 10
			arr[i+j] += num / 10
		}
	}

	var str strings.Builder
	for _, n := range arr {
		if str.Len() == 0 && n == 0 {
			continue
		}

		str.WriteByte(byte(n) + '0')
	}

	return str.String()
}
