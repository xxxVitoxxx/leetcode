package main

/*
          0         	      1                 6             8         9
      /       \     	  /       \         /       \
     00   ..  09    	 10   ..  19       60   ..  69        ..        ..
    /  \     /  \   	/  \     /  \     /   \    /   \
  000..009 090..099   100..109 190..199  600..609 690..699    ..        ..
*/
// start checking whether number is confusing number or not from 0 tree and below.
// the 0 tree will generate the 1 tree, 6 tree, 8 tree and 9 tree.
func confusingNumberII(n int) int {
	base := []int{0, 1, 6, 8, 9}
	rotated := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}

	var res int
	var dfs func(curr int)
	dfs = func(curr int) {
		for _, b := range base {
			if curr == 0 && b == 0 {
				continue
			}

			next := curr*10 + b
			if next > n {
				return
			}

			num, rotation := next, 0
			for ; num > 0; num /= 10 {
				rotation = rotation*10 + rotated[num%10]
			}
			if next != rotation {
				res++
			}

			dfs(next)
		}
	}

	dfs(0)
	return res
}

/*
          1         	      6                 8                 9
      /       \     	  /       \         /       \         /       \
     10   ..  19    	 60   ..  69       80   ..  89       90   ..  99
    /  \     /  \   	/  \     /  \     /   \    /   \    /   \    /   \
  100..109 190..199   600..609 690..699  800..809 890..899 900..909 990..999
*/
// start checking whether number is a confusing number or not
// beginning with 1, 6, 8, 9 on level 1 and continuing on to the levels below.
//
// number do not repeat across all trees.
func confusingNumberII2(n int) int {
	base := []int{0, 1, 6, 8, 9}
	rotated := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}

	var res int
	var dfs func(v, r, digits int)
	dfs = func(v, r, digits int) {
		if v > n {
			return
		}

		if v != r {
			res++
		}

		for _, b := range base {
			dfs(v*10+b, rotated[b]*digits+r, digits*10)
		}
	}

	for _, b := range base[1:] {
		dfs(b, rotated[b], 10)
	}

	return res
}
