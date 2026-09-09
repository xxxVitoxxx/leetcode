package main

func addDigits(num int) int {
	var root int
	for num > 0 {
		root += num % 10
		num /= 10

		if num == 0 && root > 9 {
			num, root = root, 0
		}
	}
	return root
}

/*
1 -> 1
2 -> 2
3 -> 3
4 -> 4
5 -> 5
6 -> 6
7 -> 7
8 -> 8
9 -> 9
10 -> 1
11 -> 2
12 -> 3
13 -> 4
14 -> 5
15 -> 6
16 -> 7
17 -> 8
18 -> 9

we can observe that the result of adding the units digit appears in a cycle,
which is from 1 to 9 repeating constantly.
therefore, we can simply take the given number module 9 and
the result will be the sum of the digits. if the result is 0 then we return 9.

time complexity: O(1)
space complexity: O(1)
*/
func addDigits2(num int) int {
	if num == 0 {
		return 0
	}

	if num%9 == 0 {
		return 9
	}

	return num % 9
}
