package main

// This is the robot's control interface.
// You should not implement it, or speculate about its implementation
type Robot struct{}

// Returns true if the cell in front is open and robot moves into the cell.
// Returns false if the cell in front is blocked and robot stays in the current cell.
func (robot *Robot) Move() bool { return true }

// Robot will stay in the same cell after calling TurnLeft/TurnRight.
// Each turn will be 90 degrees.
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}

// Clean the current cell.
func (robot *Robot) Clean() {}

// time complexity: O(n - m)
// where n is a number of cells in the room and m is a number of obstacles.
//   - we visited each non-obstacle cell once and only once.
//   - at each visit, we will check 4 directions around the cell. therefore,
//     the total number of operations would be 4 * (n - m).

// space complexity: O(n - m)
//   - we employed a hash table to keep track of whether a non-obstacle cell is visited or not.

func cleanRoom(robot *Robot) {
	// going clockwise: up, right, down, left
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make(map[[2]int]bool)
	backtrack(0, 0, 0, visited, directions, robot)
}

func goback(robot *Robot) {
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}

// the d means the direction that the robot is currently facing.
// as mentioned in the solution, 0 means the robot is facing up,
// 1 means right, 2 means down, and 3 means left.
// i keeps track of how many right turns we've made so far.
// so, each time you make a right turn, you will increase i by 1.
// assume the robot is currently facing left, so d is 3, and
// you want it to try all the directions in the following clockwise sequence: 3(left), 0(up), 1(right), 2(down).
// as you can see, since the starting direction is 3(facing left), when you add 1(turn right once),
// the answer becomes 4. however, the correct answer should be 0(up), since
// now the robot is facing up. that's why we introduce the modulo operation here,
// it is used to make sure that the direction always falls into the range of 0, 1, 2, 3.
func backtrack(row, col, d int, visited map[[2]int]bool, directions [][]int, robot *Robot) {
	visited[[2]int{row, col}] = true
	robot.Clean()
	// going clockwise: 0: up, 1: right, 2: down, 3: left
	for i := 0; i < 4; i++ {
		// going clockwise: 0: up, 1: right, 2: down, 3: left
		newDirection := (d + i) % 4
		r := row + directions[newDirection][0]
		c := col + directions[newDirection][1]
		if !visited[[2]int{r, c}] && robot.Move() {
			backtrack(r, c, newDirection, visited, directions, robot)
			goback(robot)
		}
		robot.TurnRight()
	}
}
