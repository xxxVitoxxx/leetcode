package main

import "strings"

/*
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// array
type Trie struct {
	children [26]*Trie
	exist    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, r := range word {
		index := r - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &Trie{}
		}
		curr = curr.children[index]
	}
	curr.exist = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, r := range word {
		index := r - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}

	return curr.exist
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, r := range prefix {
		index := r - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}

	return true
}

// binary search tree
type Node struct {
	word  string
	left  *Node
	right *Node
}

type Trie2 struct {
	root *Node
}

func Constructor2() Trie2 {
	return Trie2{}
}

func insert2(node *Node, word string) {
	if node == nil {
		node = &Node{word: word}
		return
	}

	if node.word > word {
		if node.left == nil {
			node.left = &Node{word: word}
		} else {
			insert2(node.left, word)
		}
	} else {
		if node.right == nil {
			node.right = &Node{word: word}
		} else {
			insert2(node.right, word)
		}
	}
}

func (this *Trie2) Insert2(word string) {
	if this.root == nil {
		this.root = &Node{word: word}
		return
	}
	insert2(this.root, word)
}

func exist(node *Node, word string) bool {
	if node == nil {
		return false
	} else if node.word == word {
		return true
	} else if node.word > word {
		return exist(node.left, word)
	}
	return exist(node.right, word)
}

func (this *Trie2) Search2(word string) bool {
	return exist(this.root, word)
}

func hasPrefix(node *Node, prefix string) bool {
	if node == nil {
		return false
	} else if strings.HasPrefix(node.word, prefix) {
		return true
	} else if node.word > prefix {
		return hasPrefix(node.left, prefix)
	}
	return hasPrefix(node.right, prefix)
}

func (this *Trie2) StartsWith2(prefix string) bool {
	return hasPrefix(this.root, prefix)
}
