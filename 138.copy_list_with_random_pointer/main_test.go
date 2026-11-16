package main

import "testing"

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name string
		head *Node
	}{{
		name: "example1",
		head: getTests1Head(),
	}, {
		name: "example2",
		head: getTests2Head(),
	}, {
		name: "example3",
		head: getTests3Head(),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copyRandomList(tt.head)

			origNodes := listNodes(tt.head)
			gotNodes := listNodes(got)

			if len(gotNodes) != len(origNodes) {
				t.Fatalf("want %d nodes, got %d", len(origNodes), len(gotNodes))
			}

			for i, o := range origNodes {
				g := gotNodes[i]
				if g == o {
					t.Errorf("node %d: copy reused the same pointer as original", i)
				}
				if g.Val != o.Val {
					t.Errorf("node %d: want val %d, got %d", i, o.Val, g.Val)
				}
				if (o.Random == nil) != (g.Random == nil) {
					t.Errorf("node %d: random nil mismatch", i)
					continue
				}
				if o.Random != nil && indexOf(origNodes, o.Random) != indexOf(gotNodes, g.Random) {
					t.Errorf("node %d: random pointer points to wrong node", i)
				}
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copyRandomList2(tt.head)

			origNodes := listNodes(tt.head)
			gotNodes := listNodes(got)

			if len(gotNodes) != len(origNodes) {
				t.Fatalf("want %d nodes, got %d", len(origNodes), len(gotNodes))
			}

			for i, o := range origNodes {
				g := gotNodes[i]
				if g == o {
					t.Errorf("node %d: copy reused the same pointer as original", i)
				}
				if g.Val != o.Val {
					t.Errorf("node %d: want val %d, got %d", i, o.Val, g.Val)
				}
				if (o.Random == nil) != (g.Random == nil) {
					t.Errorf("node %d: random nil mismatch", i)
					continue
				}
				if o.Random != nil && indexOf(origNodes, o.Random) != indexOf(gotNodes, g.Random) {
					t.Errorf("node %d: random pointer points to wrong node", i)
				}
			}
		})
	}
}

func listNodes(head *Node) []*Node {
	var nodes []*Node
	for n := head; n != nil; n = n.Next {
		nodes = append(nodes, n)
	}
	return nodes
}

func indexOf(nodes []*Node, target *Node) int {
	for i, n := range nodes {
		if n == target {
			return i
		}
	}
	return -1
}

// [[7,null],[13,0],[11,4],[10,2],[1,0]]
func getTests1Head() *Node {
	zerothNode := &Node{Val: 7}
	firstNode := &Node{Val: 13}
	secondNode := &Node{Val: 11}
	thirdNode := &Node{Val: 10}
	fourthNode := &Node{Val: 1}

	zerothNode.Next = firstNode
	firstNode.Next = secondNode
	firstNode.Random = zerothNode
	secondNode.Next = thirdNode
	secondNode.Random = thirdNode
	thirdNode.Next = fourthNode
	thirdNode.Random = secondNode
	fourthNode.Random = zerothNode

	return zerothNode
}

// [[1,1],[2,1]]
func getTests2Head() *Node {
	zerothNode := &Node{Val: 1}
	firstNode := &Node{Val: 2}

	zerothNode.Next = firstNode
	zerothNode.Random = firstNode
	firstNode.Random = firstNode

	return zerothNode
}

// [[3,null],[3,0],[3,null]]
func getTests3Head() *Node {
	zerothNode := &Node{Val: 3}
	firstNode := &Node{Val: 3}
	secondNode := &Node{Val: 3}

	zerothNode.Next = firstNode
	firstNode.Next = secondNode
	firstNode.Random = zerothNode

	return zerothNode
}
