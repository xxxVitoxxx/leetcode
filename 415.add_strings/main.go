package main

import (
	"math"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	var result string
	l1, l2, carry := len(num1)-1, len(num2)-1, 0
	for l1 >= 0 || l2 >= 0 || carry > 0 {
		if l1 >= 0 {
			carry += int(num1[l1] - '0')
			l1 -= 1
		}

		if l2 >= 0 {
			carry += int(num2[l2] - '0')
			l2 -= 1
		}

		result = strconv.Itoa(carry%10) + result
		carry /= 10
	}

	return result
}

func addStrings2(num1 string, num2 string) string {
	l1, l2 := len(num1)-1, len(num2)-1
	longest := int(math.Max(float64(len(num1)), float64(len(num2))))
	result, carry := make([]rune, longest+1), 0
	for longest > 0 {
		if l1 >= 0 {
			carry += int(num1[l1] - '0')
			l1--
		}

		if l2 >= 0 {
			carry += int(num2[l2] - '0')
			l2--
		}

		result[longest] = rune(carry%10) + '0'
		carry /= 10
		longest--
		if longest == 0 && carry > 0 {
			result[longest] = '1'
			return string(result)
		}
	}
	return string(result[1:])
}
