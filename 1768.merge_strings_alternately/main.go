package main

func mergeAlternately(word1, word2 string) string {
	var res string
	pnt1, pnt2 := 0, 0
	for pnt1 < len(word1) && pnt2 < len(word2) {
		res += string(word1[pnt1]) + string(word2[pnt2])
		pnt1, pnt2 = pnt1+1, pnt2+1
	}

	for ; pnt1 < len(word1); pnt1++ {
		res += string(word1[pnt1])
	}

	for ; pnt2 < len(word2); pnt2++ {
		res += string(word2[pnt2])
	}
	return res
}

func mergeAlternately2(word1, word2 string) string {
	var res []byte
	minLen, remains := 0, ""
	if len(word1) > len(word2) {
		minLen, remains = len(word2), word1[len(word2):]
	} else {
		minLen, remains = len(word1), word2[len(word1):]
	}

	for i := 0; i < minLen; i++ {
		res = append(res, word1[i], word2[i])
	}
	res = append(res, []byte(remains)...)
	return string(res)
}

func mergeAlternately3(word1, word2 string) string {
	var res []byte
	pnt1, pnt2 := 0, 0
	for pnt1 < len(word1) && pnt2 < len(word2) {
		res = append(res, word1[pnt1], word2[pnt2])
		pnt1, pnt2 = pnt1+1, pnt2+1
	}

	if pnt1 != len(word1) {
		res = append(res, word1[pnt1:]...)
	} else {
		res = append(res, word2[pnt2:]...)
	}
	return string(res)
}

func mergeAlternately4(word1, word2 string) string {
	var res string
	minLen := len(word1)
	if len(word2) < len(word1) {
		minLen = len(word2)
	}

	for i := 0; i < minLen; i++ {
		res += string(word1[i]) + string(word2[i])
	}
	return res + word1[minLen:] + word2[minLen:]
}

func mergeAlternately5(word1, word2 string) string {
	var res string
	for len(word1) > 0 && len(word2) > 0 {
		res += word1[:1] + word2[:1]
		word1, word2 = word1[1:], word2[1:]
	}
	return res + word1[:] + word2[:]
}
