package main

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if _, ok := m[t[i]]; !ok {
			return false
		}
		m[t[i]]--
		if m[t[i]] == 0 {
			delete(m, t[i])
		}
	}
	return len(m) == 0
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arrS := strings.Split(s, "")
	sort.Slice(arrS, func(i, j int) bool {
		return arrS[i] < arrS[j]
	})

	arrT := strings.Split(t, "")
	sort.Slice(arrT, func(i, j int) bool {
		return arrT[i] < arrT[j]
	})

	for i := 0; i < len(arrS); i++ {
		if arrS[i] != arrT[i] {
			return false
		}
	}
	return true
}

func isAnagram3(s string, t string) bool {
	arrS := strings.Split(s, "")
	arrT := strings.Split(t, "")
	sort.Strings(arrS)
	sort.Strings(arrT)
	return strings.Join(arrS, "") == strings.Join(arrT, "")
}

func isAnagram4(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
		counter[t[i]-'a']--
	}

	for i := 0; i < len(counter); i++ {
		if counter[i] != 0 {
			return false
		}
	}
	return true
}

func isAnagram5(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		counter[t[i]-'a']--
		if counter[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

func isAnagram6(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make([]int, 26)
	var dist int
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
		if counter[s[i]-'a'] == 1 {
			dist++
		}
	}

	for i := 0; i < len(t); i++ {
		counter[t[i]-'a']--
		if counter[t[i]-'a'] < 0 {
			return false
		} else if counter[t[i]-'a'] == 0 {
			dist--
		}
	}
	return dist == 0
}
