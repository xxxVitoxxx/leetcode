package main

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	node1_1 := &Node{Val: 1}
	node1_2 := &Node{Val: 2}
	node1_3 := &Node{Val: 3}
	node1_4 := &Node{Val: 4}
	node1_1.Neighbors = []*Node{node1_2, node1_4}
	node1_2.Neighbors = []*Node{node1_1, node1_3}
	node1_3.Neighbors = []*Node{node1_2, node1_4}
	node1_4.Neighbors = []*Node{node1_1, node1_3}

	tests := []struct {
		name   string
		input  *Node
		output *Node
	}{
		{"example1", node1_1, node1_1},
		{"example2", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := cloneGraph(tt.input)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got: %v, want: %v", *res, *tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := cloneGraph2(tt.input)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got: %v, want: %v", *res, *tt.output)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := cloneGraph3(tt.input)
			if !reflect.DeepEqual(res, tt.output) {
				t.Errorf("got: %v, want: %v", *res, *tt.output)
			}
		})
	}
}
