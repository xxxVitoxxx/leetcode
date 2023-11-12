package main

func isLongPressedName(name, typed string) bool {
	pnt1, pnt2 := 0, 0
	for pnt1 < len(name) && pnt2 < len(typed) {
		ncount := 1
		for pnt1 < len(name)-1 && name[pnt1] == name[pnt1+1] {
			pnt1, ncount = pnt1+1, ncount+1
		}

		tcount := 1
		for pnt2 < len(typed)-1 && typed[pnt2] == typed[pnt2+1] {
			pnt2, tcount = pnt2+1, tcount+1
		}

		if name[pnt1] != typed[pnt2] || ncount > tcount {
			return false
		}
		pnt1, pnt2 = pnt1+1, pnt2+1
	}
	return pnt1 == len(name) && pnt2 == len(typed)
}

func isLongPressedName2(name, typed string) bool {
	pnt1, pnt2 := 0, 0
	nLen, tLen := len(name), len(typed)
	for pnt1 < nLen && pnt2 < tLen {
		char := name[pnt1]
		if char != typed[pnt2] {
			return false
		}

		var ncount int
		for pnt1 < nLen && name[pnt1] == char {
			ncount, pnt1 = ncount+1, pnt1+1
		}
		var tcount int
		for pnt2 < tLen && typed[pnt2] == char {
			tcount, pnt2 = tcount+1, pnt2+1
		}

		if ncount > tcount {
			return false
		}
	}
	return pnt1 == nLen && pnt2 == tLen
}

func isLongPressedName3(name, typed string) bool {
	pnt, nLen := 0, len(name)
	var prev byte
	for i := 0; i < len(typed); i++ {
		if pnt < nLen && name[pnt] == typed[i] {
			pnt++
			prev = typed[i]
		} else if prev != typed[i] {
			return false
		}
	}
	return pnt == nLen
}
