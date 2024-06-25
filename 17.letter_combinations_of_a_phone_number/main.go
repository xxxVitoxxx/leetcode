package main

// backtracking
// time complexity: O(4^n * n), where n is the length of digits.
// note that 4 in this expression is referring to the maximum value length in the hash map,
// and not to the length of digits.
//
// space complexity: O(n), where n is the length of digits.
// not counting space used for the output, the extra space we use relative to
// input size is the space occupied by the recursion call stack.
// it will only go as deep as the number of digits in the input since
// whenever we reach that depth, we backtrack.
//
// as the hash map does not grow as the inputs grows, it occupies O(1) space.
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	letters := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	combinations := []string{}

	var backtrack func(index int, path []rune)
	backtrack = func(index int, path []rune) {
		if len(path) == len(digits) {
			combinations = append(combinations, string(path))
			return
		}

		possibleLetters := letters[rune(digits[index])]
		for _, letter := range possibleLetters {
			backtrack(index+1, append(path, letter))
		}
	}
	backtrack(0, []rune{})

	return combinations
}

func letterCombinations2(digits string) []string {
	if digits == "" {
		return []string{}
	}

	combinations := []string{}
	letters := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	var backtrack func(combination, nextDigits string)
	backtrack = func(combination, nextDigits string) {
		if nextDigits == "" {
			combinations = append(combinations, combination)
			return
		}

		possibleLetter := letters[nextDigits[0]-'2']
		for _, letter := range possibleLetter {
			backtrack(combination+string(letter), nextDigits[1:])
		}
	}
	backtrack("", digits)

	return combinations
}
