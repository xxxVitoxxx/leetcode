package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteMiddle(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	}
	ans1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	ans2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
		},
	}
	ans3 := &ListNode{
		Val: 2,
	}

	list4 := &ListNode{}

	tests := []struct {
		name   string
		head   *ListNode
		output *ListNode
	}{
		{"example1", list1, ans1},
		{"example2", list2, ans2},
		{"example3", list3, ans3},
		{"example4", list4, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := deleteMiddle(tt.head)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestDeleteMiddle2(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	}
	ans1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	ans2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
		},
	}
	ans3 := &ListNode{
		Val: 2,
	}

	list4 := &ListNode{}

	tests := []struct {
		name   string
		head   *ListNode
		output *ListNode
	}{
		{"example1", list1, ans1},
		{"example2", list2, ans2},
		{"example3", list3, ans3},
		{"example4", list4, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := deleteMiddle2(tt.head)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestDeleteMiddle3(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	}
	ans1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	ans2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
		},
	}
	ans3 := &ListNode{
		Val: 2,
	}

	list4 := &ListNode{}

	tests := []struct {
		name   string
		head   *ListNode
		output *ListNode
	}{
		{"example1", list1, ans1},
		{"example2", list2, ans2},
		{"example3", list3, ans3},
		{"example4", list4, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := deleteMiddle3(tt.head)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}

func TestDeleteMiddle4(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	}
	ans1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	ans2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
		},
	}
	ans3 := &ListNode{
		Val: 2,
	}

	list4 := &ListNode{}

	tests := []struct {
		name   string
		head   *ListNode
		output *ListNode
	}{
		{"example1", list1, ans1},
		{"example2", list2, ans2},
		{"example3", list3, ans3},
		{"example4", list4, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := deleteMiddle4(tt.head)
			assert.True(t, reflect.DeepEqual(tt.output, res))
		})
	}
}
