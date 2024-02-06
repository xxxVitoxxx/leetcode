package main

// bfs
// time complexity: O(v+e) where v is the number of rooms, and e is the number of keys
// space complexity: O(v) to store queue and visited
func canVisitAllRooms(rooms [][]int) bool {
	visited := map[int]bool{0: true}
	var queue []int
	queue = append(queue, rooms[0]...)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if !visited[curr] {
			visited[curr] = true
			queue = append(queue, rooms[curr]...)
		}
	}

	return len(rooms) == len(visited)
}

// dfs
// time complexity: O(v+e)
// space complexity: O(v)
func canVisitAllRooms2(rooms [][]int) bool {
	visited := map[int]bool{0: true}
	var dfs func([]int)
	dfs = func(keys []int) {
		for _, key := range keys {
			if !visited[key] {
				visited[key] = true
				dfs(rooms[key])
			}
		}
	}
	dfs(rooms[0])
	return len(rooms) == len(visited)
}
