package main

import "strings"

/*
t -> f
w -> e
r -> t
e -> r

=> wertf
*/
// Topological Sort Using Kahn's Algorithm(BFS)
//
// time complexity: O(n). n is the total length of all the words in the input list, added together.
// space complexity: O(1)
func alienOrder(words []string) string {
	inDegree := make(map[byte]int)
	for _, word := range words {
		for i := range word {
			inDegree[word[i]] = 0
		}
	}

	adjList := make(map[byte]map[byte]bool)
	for i := 0; i < len(words)-1; i++ {
		if len(words[i]) > len(words[i+1]) && strings.HasPrefix(words[i], words[i+1]) {
			return ""
		}

		for j := 0; j < len(words[i]) && j < len(words[i+1]); j++ {
			if words[i][j] == words[i+1][j] {
				continue
			}

			if adjList[words[i][j]] == nil {
				adjList[words[i][j]] = map[byte]bool{}
			}

			if !adjList[words[i][j]][words[i+1][j]] {
				adjList[words[i][j]][words[i+1][j]] = true
				inDegree[words[i+1][j]]++
			}

			break
		}
	}

	var queue []byte
	for b, d := range inDegree {
		if d == 0 {
			queue = append(queue, b)
		}
	}

	var str strings.Builder
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		str.WriteByte(curr)

		for next := range adjList[curr] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if str.Len() != len(inDegree) {
		return ""
	}

	return str.String()
}

// Topological Sort Using Kahn's Algorithm(DFS)
//
// time complexity: O(n). n is the total length of all the words in the input list, added together.
// space complexity: O(1)
func alienOrder2(words []string) string {
	states := make(map[byte]int)
	for _, word := range words {
		for i := range word {
			states[word[i]] = 0
		}
	}

	adjList := make(map[byte]map[byte]bool)
	for i := 0; i < len(words)-1; i++ {
		if len(words[i]) > len(words[i+1]) && strings.HasPrefix(words[i], words[i+1]) {
			return ""
		}

		for j := 0; j < len(words[i]) && j < len(words[i+1]); j++ {
			if words[i][j] == words[i+1][j] {
				continue
			}

			if adjList[words[i][j]] == nil {
				adjList[words[i][j]] = map[byte]bool{}
			}

			adjList[words[i][j]][words[i+1][j]] = true

			break
		}
	}

	var result []byte
	var hasCycle func(b byte) bool
	hasCycle = func(b byte) bool {
		if states[b] == 1 {
			return true
		}

		if states[b] == 2 {
			return false
		}

		states[b] = 1
		for next := range adjList[b] {
			if hasCycle(next) {
				return true
			}
		}

		states[b] = 2
		result = append(result, b)

		return false
	}

	for b := range states {
		if hasCycle(b) {
			return ""
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
