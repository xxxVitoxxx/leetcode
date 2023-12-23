package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	ans1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	ans2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
		},
	}

	tests := []struct {
		name   string
		head   *ListNode
		output *ListNode
	}{
		{"example1", list1, ans1},
		{"example2", list2, ans2},
		{"example3", &ListNode{}, &ListNode{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseList(tt.head)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestReverseList2(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	ans1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	ans2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
		},
	}

	tests := []struct {
		name   string
		head   *ListNode
		output *ListNode
	}{
		{"example1", list1, ans1},
		{"example2", list2, ans2},
		{"example3", &ListNode{}, &ListNode{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseList2(tt.head)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestReverseList3(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	ans1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	ans2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
		},
	}

	tests := []struct {
		name   string
		head   *ListNode
		output *ListNode
	}{
		{"example1", list1, ans1},
		{"example2", list2, ans2},
		{"example3", &ListNode{}, &ListNode{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := reverseList3(tt.head)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}
