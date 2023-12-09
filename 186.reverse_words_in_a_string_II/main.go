package main

func reverseWords(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left, right = left+1, right-1
	}

	pnt1, pnt2 := 0, 0
	for i := 0; i < len(s); i++ {
		if i+1 == len(s) || s[i+1] == ' ' {
			pnt2 = i
			for pnt1 < pnt2 {
				s[pnt1], s[pnt2] = s[pnt2], s[pnt1]
				pnt1, pnt2 = pnt1+1, pnt2-1
			}
			pnt1 = i + 2
		}
	}
}

func reverseWords2(s []byte) {
	reverse2(s, 0, len(s))
	var left int
	for left < len(s) {
		right := left
		for right < len(s) && s[right] != ' ' {
			right++
		}
		reverse2(s, left, right)
		left = right + 1
	}
}

func reverse2(s []byte, left, right int) {
	// s[right] 會是 ' ' 或是 right == len(s)
	// 所以不能直接將 left 與 right 互換
	for i := 0; i < (right-left)/2; i++ {
		tmp := s[right-i-1]
		s[right-i-1] = s[left+i]
		s[left+i] = tmp
	}
}

func reverseWords3(s []byte) {
	left, right := 0, 0
	for left < len(s) {
		right = left
		for right < len(s) && s[right] != ' ' {
			right++
		}
		reverse3(s[left:right])
		left = right + 1
	}
	reverse3(s)
}

func reverse3(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left, right = left+1, right-1
	}
}
