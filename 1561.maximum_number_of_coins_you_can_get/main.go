package main

import (
	"golang.org/x/exp/slices"
)

func maxCoins(piles []int) int {
	var res int
	slices.Sort(piles)
	for len(piles) > 0 {
		piles = piles[1:]
		piles = piles[:len(piles)-1]
		res += piles[len(piles)-1]
		piles = piles[:len(piles)-1]
	}
	return res
}

func maxCoins2(piles []int) int {
	var res int
	slices.Sort(piles)
	left, right := 0, len(piles)-1
	for left < right {
		res += piles[right-1]
		left, right = left+1, right-2
	}
	return res
}

// 預留前 len(piles)/3 給 bob
func maxCoins3(piles []int) int {
	var res int
	slices.Sort(piles)
	for i := len(piles) / 3; i < len(piles); i += 2 {
		res += piles[i]
	}
	return res
}
