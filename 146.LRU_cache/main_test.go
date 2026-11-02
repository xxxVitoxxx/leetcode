package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	obj := Constructor(2)

	obj.Put(1, 1)
	obj.Put(2, 2)
	if value := obj.Get(1); value != 1 {
		t.Errorf("got: %v, want: 1", value)
	}

	obj.Put(3, 3)
	if value := obj.Get(2); value != -1 {
		t.Errorf("got: %v, want: -1", value)
	}

	obj.Put(4, 4)
	if value := obj.Get(1); value != -1 {
		t.Errorf("got: %v, want: -1", value)
	}
	if value := obj.Get(3); value != 3 {
		t.Errorf("got: %v, want: 3", value)
	}
	if value := obj.Get(4); value != 4 {
		t.Errorf("got: %v, want: 4", value)
	}
}
