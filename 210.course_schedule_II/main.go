package main

// Topological Sort Using Kahn's Algorithm(BFS)
//
// n is the number of courses and m is the size of prerequisites
// time complexity: O(n+m)
// space complexity: O(n+m)
func findOrder(numCourses int, prerequisites [][]int) []int {
	adjList := make(map[int][]int)
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		adjList[pre[1]] = append(adjList[pre[1]], pre[0])
		inDegree[pre[0]]++
	}

	var queue []int
	for i, d := range inDegree {
		if d == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]int, 0, numCourses)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		result = append(result, curr)

		for _, course := range adjList[curr] {
			inDegree[course]--
			if inDegree[course] == 0 {
				queue = append(queue, course)
			}
		}
	}

	if len(result) != numCourses {
		return []int{}
	}

	return result
}

// Topological Sort Using Kahn's Algorithm(DFS)
//
// n is the number of courses and m is the size of prerequisites
// time complexity: O(n+m)
// space complexity: O(n+m)
func findOrder2(numCourses int, prerequisites [][]int) []int {
	adjList := make(map[int][]int)
	for _, pre := range prerequisites {
		adjList[pre[1]] = append(adjList[pre[1]], pre[0])
	}

	// value description
	// 0 -> not visited yet
	// 1 -> recursive call in progress
	// 2 -> recursive call finished and no cycle found
	states := make([]int, numCourses)
	result := make([]int, 0, numCourses)

	var hasCycle func(node int) bool
	hasCycle = func(node int) bool {
		if states[node] == 1 {
			return true
		}

		if states[node] == 2 {
			return false
		}

		states[node] = 1
		for _, course := range adjList[node] {
			if hasCycle(course) {
				return true
			}
		}
		// if slices.ContainsFunc(adjList[node], hasCycle)  {
		// 	return true
		// }

		states[node] = 2
		result = append(result, node)

		return false
	}

	for course := range numCourses {
		if hasCycle(course) {
			return []int{}
		}
	}

	if len(result) != numCourses {
		return []int{}
	}

	// reverse
	// 是遞迴到最後再開始把 course 加進 result ，所以最後需要將順序反轉
	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}

	return result
}
