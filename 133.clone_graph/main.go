package main

// Definition for a Node
type Node struct {
	Val       int
	Neighbors []*Node
}

// bfs iteration
// time complexity: O(v+e)
// space complexity: O(v)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	visited := make(map[int]*Node)
	visited[node.Val] = &Node{Val: node.Val}
	queue := []*Node{node}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for i := 0; i < len(curr.Neighbors); i++ {
			if _, ok := visited[curr.Neighbors[i].Val]; !ok {
				queue = append(queue, curr.Neighbors[i])
				visited[curr.Neighbors[i].Val] = &Node{Val: curr.Neighbors[i].Val}
			}
			visited[curr.Val].Neighbors = append(visited[curr.Val].Neighbors, visited[curr.Neighbors[i].Val])
		}
	}

	return visited[node.Val]
}

// dfs recursion
// time complexity: O(v+e)
// space complexity: O(v)
func cloneGraph2(node *Node) *Node {
	if node == nil {
		return node
	}

	visited := make(map[int]*Node)
	var dfs func(*Node)
	dfs = func(n *Node) {
		visited[n.Val] = &Node{Val: n.Val}
		for i := 0; i < len(n.Neighbors); i++ {
			if _, ok := visited[n.Neighbors[i].Val]; !ok {
				dfs(n.Neighbors[i])
			}
			visited[n.Val].Neighbors = append(visited[n.Val].Neighbors, visited[n.Neighbors[i].Val])
		}
	}

	dfs(node)
	return visited[node.Val]
}

// dfs recursion
// time complexity: O(v+e)
// space complexity: O(v)
func cloneGraph3(node *Node) *Node {
	if node == nil {
		return node
	}

	visited := make(map[int]*Node)
	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if cloneNode, ok := visited[n.Val]; ok {
			return cloneNode
		}

		cloneNode := &Node{
			Val:       n.Val,
			Neighbors: make([]*Node, 0, len(n.Neighbors)),
		}
		visited[n.Val] = cloneNode

		for i := 0; i < len(n.Neighbors); i++ {
			cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(n.Neighbors[i]))
		}

		return cloneNode
	}

	return dfs(node)
}
