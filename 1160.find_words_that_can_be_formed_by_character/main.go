package main

// use hash
func countCharacters(words []string, chars string) int {
	charsMap := make(map[rune]int)
	for _, r := range chars {
		charsMap[r]++
	}

	var ans int
	for _, word := range words {
		wordMap := make(map[rune]int)
		for _, r := range word {
			wordMap[r]++
		}

		good := true
		for k, v := range wordMap {
			if charsMap[k] < v {
				good = false
				break
			}
		}
		if good {
			ans += len(word)
		}
	}
	return ans
}

// use 26 sized array
// 因為都是小寫英文，所以可以用字母的 ASCII 值減去 a 的 ASCII ，並將其值映射到 array 的 index
// e.g.
// 'a' - 'a' = 0
// 'b' - 'a' = 1
// ...
// 'z' - 'a' = 25
func countCharacters2(words []string, chars string) int {
	charsCount := make([]int, 26)
	for _, r := range chars {
		charsCount[int(r-'a')]++
	}

	var ans int
	for _, word := range words {
		wordCount := make([]int, 26)
		for _, r := range word {
			wordCount[int(r-'a')]++
		}

		good := true
		for i := 0; i < len(wordCount); i++ {
			if charsCount[i] < wordCount[i] {
				good = false
				break
			}
		}
		if good {
			ans += len(word)
		}
	}
	return ans
}

// optimize countCharacters2
func countCharacters3(words []string, chars string) int {
	charsCount := make([]int, 26)
	for _, r := range chars {
		charsCount[r-'a']++
	}

	var ans int
	for _, word := range words {
		currentCount := make([]int, 26)
		copy(currentCount, charsCount)

		good := true
		for _, r := range word {
			currentCount[r-'a']--
			if currentCount[r-'a'] < 0 {
				good = false
				break
			}
		}
		if good {
			ans += len(word)
		}
	}
	return ans
}
