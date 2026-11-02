package main

type node struct {
	key, value int
	prev, next *node
}

func newNode(key, value int) *node {
	return &node{key: key, value: value}
}

type LRUCache struct {
	capacity   int
	m          map[int]*node
	head, tail *node
}

func Constructor(capacity int) LRUCache {
	head, tail := newNode(-1, -1), newNode(-1, -1)
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) remove(node *node) {
	delete(this.m, node.key)

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node.prev, node.next = nil, nil
}

func (this *LRUCache) insert(node *node) {
	this.m[node.key] = node

	next := this.head.next
	next.prev = node

	node.prev = this.head
	node.next = next

	this.head.next = node
}

// time complexity: O(1)
func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if !ok {
		return -1
	}

	this.remove(node)
	this.insert(node)

	return node.value
}

// time complexity: O(1)
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.value = value
		this.remove(node)
		this.insert(node)
		return
	}

	n := &node{
		key:   key,
		value: value,
	}

	if len(this.m) == this.capacity {
		this.remove(this.tail.prev)
	}

	this.insert(n)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
