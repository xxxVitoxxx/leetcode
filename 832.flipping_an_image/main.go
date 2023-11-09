package main

func flipAndInvertImage(image [][]int) [][]int {
	for i := range image {
		left, right := 0, len(image[i])-1
		for left < right {
			image[i][left], image[i][right] = image[i][right], image[i][left]
			left, right = left+1, right-1
		}
	}

	for i := range image {
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] == 0 {
				image[i][j] = 1
			} else {
				image[i][j] = 0
			}
		}
	}
	return image
}

/*
image = [[1, 1, 0, 0], [1, 0, 0, 1], [0, 1, 1, 1], [1, 0, 1, 0]]
return  [[invert(image[0][3]), invert(image[0][2]), invert(image[0][1]), invert(image[0][0])],...]
*/
func flipAndInvertImage2(image [][]int) [][]int {
	for i := range image {
		left, right := 0, len(image[i])-1
		for left < right {
			image[i][left], image[i][right] = invert(image[i][right]), invert(image[i][left])
			left, right = left+1, right-1
		}
		if left == right {
			image[i][left] = invert(image[i][left])
		}
	}
	return image
}

func invert(n int) int {
	if n == 0 {
		return 1
	}
	return 0
}

func flipAndInvertImage3(image [][]int) [][]int {
	for i := range image {
		left, right := 0, len(image[i])-1
		for left < right {
			image[i][left], image[i][right] = image[i][right]^1, image[i][left]^1
			left, right = left+1, right-1
		}
		if left == right {
			image[i][left] ^= 1
		}
	}
	return image
}

func flipAndInvertImage4(image [][]int) [][]int {
	for i := range image {
		for left, right := 0, len(image[i])-1; left <= right; left, right = left+1, right-1 {
			image[i][left], image[i][right] = image[i][right]^1, image[i][left]^1
		}
	}
	return image
}
