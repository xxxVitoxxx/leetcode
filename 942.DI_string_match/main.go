package main

func diStringMatch(s string) []int {
	r := make([]int, len(s)+1)
	for i := range r {
		r[i] = i
	}

	res := make([]int, len(s)+1)
	for i := 0; i < len(res); i++ {
		if i == len(res)-1 {
			res[i] = r[0]
		} else if s[i] == 'I' {
			res[i] = r[0]
			r = r[1:]
		} else {
			res[i] = r[len(r)-1]
			r = r[:len(r)-1]
		}
	}
	return res
}

func diStringMatch2(s string) []int {
	rMin, rMax := 0, len(s)
	res := make([]int, rMax+1)
	for i := 0; i < len(res); i++ {
		if i == len(s) || s[i] == 'I' {
			res[i], rMin = rMin, rMin+1
		} else {
			res[i], rMax = rMax, rMax-1
		}
	}
	return res
}

func diStringMatch3(s string) []int {
	rMin, rMax := 0, len(s)
	res := make([]int, rMax+1)
	for i, r := range s {
		if r == 'I' {
			res[i], rMin = rMin, rMin+1
		} else {
			res[i], rMax = rMax, rMax-1
		}
	}
	res[len(s)] = rMin
	return res
}
