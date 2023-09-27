package main

func backspaceCompare(s, t string) bool {
	rs := getRune(s)
	rt := getRune(t)
	return string(rs) == string(rt)
}

func getRune(str string) []rune {
	var r []rune
	for _, s := range str {
		if s == '#' {
			if len(r) > 0 {
				r = r[:len(r)-1]
			}
		} else {
			r = append(r, s)
		}
	}
	return r
}
