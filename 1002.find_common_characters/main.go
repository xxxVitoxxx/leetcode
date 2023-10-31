package main

import "fmt"

// use hash table
func commonChars(words []string) []string {
	resMap := make(map[rune]int)
	for _, word := range words[0] {
		resMap[word]++
	}

	for _, word := range words[1:] {
		wordMap := make(map[rune]int)
		for _, w := range word {
			wordMap[w]++
		}

		for k, v := range resMap {
			if val, ok := wordMap[k]; !ok {
				delete(resMap, k)
			} else {
				resMap[k] = min(v, val)
			}
		}
	}

	var res []string
	for k, v := range resMap {
		for v > 0 {
			res = append(res, string(k))
			v--
		}
	}
	return res
}

func min(n1, n2 int) int {
	if n1 <= n2 {
		return n1
	}
	return n2
}

func commonChars2(words []string) []string {
	arr := make([][26]byte, len(words))
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			arr[i][words[i][j]-'a']++
		}
	}

	var res []string
	for i := 0; i < 26; i++ {
		isValid, count := true, 1<<63-1
		for j := 0; j < len(words); j++ {
			if arr[j][i] <= 0 {
				isValid = false
				break
			} else if count > int(arr[j][i]) {
				count = int(arr[j][i])
			}
		}

		if isValid {
			for count > 0 {
				res = append(res, fmt.Sprintf("%c", i+'a'))
				count--
			}
		}
	}
	return res
}

// 先計算所有 word 字母出現的次數
// 再檢查字母是否存在於每個 word 及出現的最小次數
func commonChars3(words []string) []string {
	arr := make([][26]int, len(words))
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			arr[i][words[i][j]-'a']++
		}
	}

	var res []string
	for i := 0; i < 26; i++ {
		isValid, count := true, 100
		for j := 0; j < len(words); j++ {
			if arr[j][i] == 0 {
				isValid = false
				break
			} else if count > arr[j][i] {
				count = arr[j][i]
			}
		}

		if isValid {
			for count > 0 {
				res = append(res, fmt.Sprintf("%c", i+'a'))
				count--
			}
		}
	}
	return res
}
