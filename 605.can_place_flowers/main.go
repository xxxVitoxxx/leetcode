package main

// time complexity: O(n)
// space complexity: O(1)
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			emptyLeft := i == 0 || flowerbed[i-1] == 0
			emptyRight := i == len(flowerbed)-1 || flowerbed[i+1] == 0
			if emptyLeft && emptyRight {
				flowerbed[i] = 1
				n--
			}
		}
	}
	return n <= 0
}

// time complexity: O(n)
// space complexity: O(1)
func canPlaceFlowers2(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}
