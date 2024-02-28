package main

import "strings"

// queue
// time complexity: O(n)
// space complexity: O(n)
func predictPartyVictory(senate string) string {
	l := len(senate)
	var rQueue, dQueue []int
	for i := 0; i < l; i++ {
		if senate[i] == 'R' {
			rQueue = append(rQueue, i)
		} else {
			dQueue = append(dQueue, i)
		}
	}

	for len(rQueue) > 0 && len(dQueue) > 0 {
		if rQueue[0] < dQueue[0] {
			rQueue = append(rQueue, rQueue[0]+l)
		} else {
			dQueue = append(dQueue, dQueue[0]+l)
		}
		rQueue = rQueue[1:]
		dQueue = dQueue[1:]
	}

	if len(rQueue) > 0 {
		return "Radiant"
	}
	return "Dire"
}

// one queue
// time complexity: O(n)
// space complexity: O(n)
func predictPartyVictory2(senate string) string {
	queue := []uint8(senate)
	rCount := strings.Count(senate, "R")
	dCount := strings.Count(senate, "D")
	for rCount > 0 && dCount > 0 {
		winParty := queue[0]
		queue = queue[1:]
		for i := 0; i < len(queue); i++ {
			if queue[i] != winParty {
				if queue[i] == 'R' {
					rCount--
				} else {
					dCount--
				}

				queue = append(queue[:i], queue[i+1:]...)
				break
			}
		}

		queue = append(queue, winParty)
	}

	if rCount > 0 {
		return "Radiant"
	}
	return "Dire"
}

// one queue
// time complexity: O(n)
// space complexity: O(n)
func predictPartyVictory3(senate string) string {
	parties := map[uint8]string{
		'R': "Radiant",
		'D': "Dire",
	}

	queue := []uint8(senate)
	for len(queue) > 1 {
		first := queue[0]
		queue = queue[1:]

		var i int
		for ; i < len(queue); i++ {
			if queue[i] != first {
				queue = append(queue[:i], queue[i+1:]...)
				break
			}
		}
		if i == len(queue) {
			return parties[first]
		}

		queue = append(queue, first)
	}

	return parties[queue[0]]
}

// optimize one queue
// time complexity: O(n)
// space complexity: O(n)
func predictPartyVictory4(senate string) string {
	queue := []uint8(senate)
	rCount := strings.Count(senate, "R")
	dCount := strings.Count(senate, "D")
	rBan, dBan := 0, 0
	for rCount > 0 && dCount > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == 'R' {
			if rBan > 0 {
				rBan--
				rCount--
			} else {
				dBan++
				queue = append(queue, curr)
			}
		} else {
			if dBan > 0 {
				dBan--
				dCount--
			} else {
				rBan++
				queue = append(queue, curr)
			}
		}
	}

	if rCount > 0 {
		return "Radiant"
	}
	return "Dire"
}
