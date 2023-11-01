package main

import "strings"

func reverseStr(s string, k int) string {
	sr := []rune(s)
	for i := 0; i < len(sr); {
		reverse := i + k
		if reverse > len(sr) {
			reverse = len(sr)
		}
		for l, r := i, reverse-1; l < r; l, r = l+1, r-1 {
			sr[l], sr[r] = sr[r], sr[l]
		}
		i += 2 * k
	}
	return string(sr)
}

func reverseStr2(s string, k int) string {
	sr := []rune(s)
	for i, j := 0, k-1; i < len(s); i, j = i+2*k, j+2*k {
		for left, right := i, min(j, len(sr)-1); left < right; left, right = left+1, right-1 {
			sr[left], sr[right] = sr[right], sr[left]
		}
	}
	return string(sr)
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

// note s[0] -> byte s[0:1] -> string
func reverseStr3(s string, k int) string {
	i, res := 0, ""
	for {
		if i+2*k < len(s) {
			tmp := s[i : i+2*k]
			res += reverse(tmp[:k]) + tmp[k:]
			i += 2 * k
		} else {
			if i+k < len(s) {
				tmp := s[i:]
				res += reverse(tmp[:k]) + tmp[k:]
			} else {
				res += reverse(s[i:])
			}
			break
		}
	}
	return res
}

func reverse(s string) string {
	r := []rune(s)
	for left, right := 0, len(r)-1; left < right; left, right = left+1, right-1 {
		r[left], r[right] = r[right], r[left]
	}
	return string(r)
}

func reverseStr4(s string, k int) string {
	var sb strings.Builder
	i, l := 0, len(s)
	for i < l {
		if i+2*k < l {
			j := i + k
			for a := j - 1; a >= i; a-- {
				sb.WriteByte(s[a])
			}
			sb.WriteString(s[j : i+2*k])
			i += 2 * k
		} else {
			if i+k < l {
				j := i + k
				for a := j - 1; a >= i; a-- {
					sb.WriteByte(s[a])
				}
				sb.WriteString(s[j:])
			} else {
				for a := l - 1; a >= i; a-- {
					sb.WriteByte(s[a])
				}
			}
			break
		}
	}
	return sb.String()
}
