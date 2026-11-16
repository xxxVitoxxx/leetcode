package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name     string
		root     *Node
		expected []string
	}{{
		name: "example1",
		root: &Node{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 4,
				},
				Right: &Node{
					Val: 5,
				},
			},
			Right: &Node{
				Val: 3,
				Right: &Node{
					Val: 7,
				},
			},
		},
		expected: []string{"1", "#", "2", "3", "#", "4", "5", "7", "#"},
	}, {
		name:     "example2",
		root:     nil,
		expected: []string{},
	}, {
		name: "example3",
		root: &Node{
			Val: 1,
			Right: &Node{
				Val: 3,
			},
		},
		expected: []string{"1", "#", "3", "#"},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := connect(cloneNode(tt.root))
			gotOrder := order(got)

			if !reflect.DeepEqual(gotOrder, tt.expected) {
				t.Errorf("got %v, want %v", gotOrder, tt.expected)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := connect2(cloneNode(tt.root))
			gotOrder := order(got)

			if !reflect.DeepEqual(gotOrder, tt.expected) {
				t.Errorf("got %v, want %v", gotOrder, tt.expected)
			}
		})
	}
}

func cloneNode(root *Node) *Node {
	if root == nil {
		return nil
	}

	return &Node{
		Val:   root.Val,
		Left:  cloneNode(root.Left),
		Right: cloneNode(root.Right),
	}
}

func order(root *Node) []string {
	if root == nil {
		return []string{}
	}

	queue := []*Node{root}
	var result []string

	for len(queue) > 0 {
		size := len(queue)
		for i := range size {
			curr := queue[i]

			result = append(result, strconv.Itoa(curr.Val))
			if curr.Next == nil {
				result = append(result, "#")
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		queue = queue[size:]
	}

	return result
}
