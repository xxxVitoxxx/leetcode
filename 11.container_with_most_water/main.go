package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		var area int
		if height[left] <= height[right] {
			area = (right - left) * height[left]
			left++
		} else {
			// height[left] > height[right]
			area = (right - left) * height[right]
			right--
		}

		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func maxArea2(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left != right {
		area := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, area)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
