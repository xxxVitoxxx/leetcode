package main

import (
	"golang.org/x/exp/slices"
)

// brute force
func buyChoco(prices []int, money int) int {
	cost := 101
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			total := prices[i] + prices[j]
			if total <= money && total < cost {
				cost = total
			}
		}
	}
	if money >= cost {
		return money - cost
	}
	return money
}

// greedy
func buyChoco2(prices []int, money int) int {
	slices.Sort(prices)
	cost := prices[0] + prices[1]
	if money >= cost {
		return money - cost
	}
	return money
}

// counting sort
func buyChoco3(prices []int, money int) int {
	freq := make([]int, 101)
	for _, price := range prices {
		freq[price]++
	}

	minmum, secondMinmum := 0, 0
	for i := 1; i < len(freq); i++ {
		if freq[i] > 1 {
			minmum, secondMinmum = i, i
			break
		} else if freq[i] == 1 {
			minmum = i
			break
		}
	}

	if secondMinmum == 0 {
		for i := minmum + 1; i < len(freq); i++ {
			if freq[i] > 0 {
				secondMinmum = i
				break
			}
		}
	}

	cost := minmum + secondMinmum
	if money >= cost {
		return money - cost
	}
	return money
}

// two passes
func buyChoco4(prices []int, money int) int {
	minIndex := getMinIndex(prices)
	minCost := prices[minIndex]
	prices = append(prices[:minIndex], prices[minIndex+1:]...)
	secondMinIndex := getMinIndex(prices)
	cost := minCost + prices[secondMinIndex]
	if money >= cost {
		return money - cost
	}
	return money
}

func getMinIndex(prices []int) int {
	var minIndex int
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

// one pass
func buyChoco5(prices []int, money int) int {
	minmum, secondMinmum := 101, 101
	for i := 0; i < len(prices); i++ {
		if prices[i] < minmum {
			secondMinmum = minmum
			minmum = prices[i]
		} else if prices[i] < secondMinmum {
			secondMinmum = prices[i]
		}
	}

	cost := minmum + secondMinmum
	if money >= cost {
		return money - cost
	}
	return money
}
