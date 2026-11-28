package main

type DetectSquares struct {
	count map[[2]int]int
	list  [][2]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		count: make(map[[2]int]int),
		list:  [][2]int{},
	}
}

func (this *DetectSquares) Add(point []int) {
	key := [2]int{point[0], point[1]}
	this.count[key]++
	this.list = append(this.list, key)
}

func (this *DetectSquares) Count(point []int) int {
	var result int
	x, y := point[0], point[1]
	for _, p := range this.list {
		dx, dy := x-p[0], y-p[1]
		if dx == 0 || (dx != dy && dx != -dy) {
			continue
		}

		result += this.count[[2]int{x, p[1]}] * this.count[[2]int{p[0], y}]
	}
	return result
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
