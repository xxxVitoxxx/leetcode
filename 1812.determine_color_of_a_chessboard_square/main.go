package main

import (
	"slices"
	"strconv"
)

// time complexity: O(1)
// space complexity: O(1)
// nolint:unused
func squareIsWhite(coordinates string) bool {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8"}

	cooBytes := []byte(coordinates)
	letter := string(cooBytes[0])
	number := string(cooBytes[1])

	letterIndex := slices.Index(letters, letter)
	numberIndex := slices.Index(numbers, number)
	if letterIndex%2 == 0 || letterIndex == 0 {
		if numberIndex%2 == 0 || numberIndex == 0 {
			return false
		} else {
			return true
		}
	} else {
		if numberIndex%2 == 0 || numberIndex == 0 {
			return true
		} else {
			return false
		}
	}
}

// time complexity: O(1)
// space complexity: O(1)
// nolint:unused
func squareIsWhite2(coordinates string) bool {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	letter := coordinates[:1]
	number := coordinates[1:]

	col := slices.Index(letters, letter)
	row, err := strconv.Atoi(number)
	if err != nil {
		return false
	}

	return (col+row)%2 == 0
}

// observation: 以 a1 當錨點，往水平或是垂直方向移動一步，顏色就會翻轉
// 所以重點是從 a1 開始，計算座標距離，如果是奇數則為白色，偶數則為黑色
//
// time complexity: O(1)
// space complexity: O(1)
// nolint:unused
func squareIsWhite3(coordinates string) bool {
	return (coordinates[0]+coordinates[1])%2 == 1
}
