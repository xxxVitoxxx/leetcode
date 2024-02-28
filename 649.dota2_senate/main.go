package main

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
