package main

// DFS
// time complexity: O(m * n)
// space complexity: O(m * n)

// we can use a graph traversal algorithm like depth-first search or
// breadth-first search to check if a path exists from the source node
// to the destination node. in this approach, we will solve the problem using DFS.

// in DFS, we use a recursive function to explore node as far as possible along each branch.
// upon reaching the end of a branch, we backtrack to the previous node and
// continue exploring the next branches.

// once we encounter an unvisited node, we will take one of its neighbor nodes(if exists)
// as the next node on this branch. recursively call the function to take the next node
// as the starting node and solve the subproblem.

// we begin a DFS traversal from the start cell, treating it as a node.
// we determine which cells the ball will roll into until it strikes a wall
// in each of the four directions, and then we run the DFS recursively from these cells.
// we continue traversing in this fashion recursively until we reach the destination cell
// or there are no more unvisited cells.

// to figure out which cell the ball will roll out of, we can keep traveling in that directions
// using a for loop until we reach a cell with a wall. the ball will roll to the preceding cell.
func hasPath(maze [][]int, start []int, destination []int) bool {
	m, n := len(maze), len(maze[0])
	visited := make(map[[2]int]bool)
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var dfs func([]int) bool
	dfs = func(curr []int) bool {
		if visited[[2]int{curr[0], curr[1]}] {
			return false
		}

		visited[[2]int{curr[0], curr[1]}] = true
		if curr[0] == destination[0] && curr[1] == destination[1] {
			return true
		}

		for i := 0; i < len(directions); i++ {
			r, c := curr[0], curr[1]
			for r >= 0 && r < m && c >= 0 && c < n && maze[r][c] == 0 {
				r += directions[i][0]
				c += directions[i][1]
			}

			r -= directions[i][0]
			c -= directions[i][1]
			if dfs([]int{r, c}) {
				return true
			}
		}

		return false
	}

	return dfs(start)
}

// BFS
// time complexity: O(m * n)
// space complexity: O(m * n)

// as our task is to check if there exists a path from the start cell to the destination cell,
// we can use a breadth-first search algorithm as well.

// if traverses in a level-wise manner, i.e. all the nodes at the present level are
// explored before moving on to the nodes at the next level, where a level is
// the distance from a starting node. BFS is implemented with a queue.

// similar to a DFS traversal, we begin a BFS from the start cell,
// treating it as a node. we push start into the BFS queue.
// we determine with cells the ball with roll into until it strikes a wall
// in each of the four directions, and then push these cells into the BFS queue.
// we continue traversing in this fashion until we reach the destination cell
// or there are no more cells that can visited.
func hasPath2(maze [][]int, start []int, destination []int) bool {
	m, n := len(maze), len(maze[0])
	visited := make(map[[2]int]bool)
	// up, right, down, left
	directions := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	queue := [][]int{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr[0] == destination[0] && curr[1] == destination[1] {
			return true
		}

		for i := 0; i < len(directions); i++ {
			r := curr[0]
			c := curr[1]
			// move the ball in the chosen direction until it can.
			for r >= 0 && r < m && c >= 0 && c < n && maze[r][c] == 0 {
				r += directions[i][0]
				c += directions[i][1]
			}

			// revert the last move to get the cell to which the ball rolls.
			r -= directions[i][0]
			c -= directions[i][1]
			if !visited[[2]int{r, c}] {
				visited[[2]int{r, c}] = true
				queue = append(queue, []int{r, c})
			}
		}
	}

	return false
}
