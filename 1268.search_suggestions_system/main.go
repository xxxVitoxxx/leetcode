package main

import (
	"golang.org/x/exp/slices"
)

// time complexity: O(m * n) where n is the length of the products and m is the length of the search word.
// space complexity: O(n)
func suggestedProducts(products []string, searchWord string) [][]string {
	slices.Sort(products)
	result := make([][]string, len(searchWord))
	for i := range searchWord {
		for _, product := range products {
			if len(product) >= i && len(result[i]) < 3 && searchWord[:i+1] == product[:i+1] {
				result[i] = append(result[i], product)
			}
		}
	}

	return result
}

// binary search
// time complexity: O(m * logn * p) where n is the length of the products, m is the length of the search word,
// and p is the length of the prefix.
// the pd sort in the worst case is n * logn and loop search word use binary search and loop len(prefix) is m * logn * p
//
// space complexity: O(m)
func suggestedProducts2(products []string, searchWord string) [][]string {
	slices.Sort(products)
	result, prefix := make([][]string, len(searchWord)), ""
	for i, r := range searchWord {
		prefix += string(r)
		index, _ := slices.BinarySearch(products, prefix)
		for j := index; j < len(products) && j < index+3; j++ {
			if !startWith(products[j], prefix) {
				break
			}
			result[i] = append(result[i], products[j])
		}
	}

	return result
}

func startWith(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}

	for i := range prefix {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

// trie and dfs
// time complexity: O(n) where n is total number of characters in products.
// for each prefix we find its representative node in O(len(prefix)) and
// dfs to find at most 3 words which is an O(1) operation.
// thus the overall complexity is dominated by the time required to build trie.
//
// space complexity: O(n) where n is the number of nodes in the trie.
// O(26n) = O(n). 26 is the alphabet size.
func suggestedProducts3(products []string, searchWord string) [][]string {
	trie := newTrie()
	for _, product := range products {
		trie.insert(product)
	}

	results := make([][]string, len(searchWord))
	prefix := []byte{}
	for i := range searchWord {
		prefix = append(prefix, searchWord[i])
		results[i] = trie.startWith(prefix)
	}

	return results
}

type Node struct {
	isWord   bool
	children [26]*Node
}

func newNode() *Node {
	return &Node{}
}

type Trie struct {
	root *Node
}

func newTrie() *Trie {
	return &Trie{newNode()}
}

func (t *Trie) insert(word string) {
	curr := t.root
	for _, r := range word {
		index := r - 'a'
		if curr.children[index] == nil {
			curr.children[index] = newNode()
		}
		curr = curr.children[index]
	}
	curr.isWord = true
}

func (t *Trie) startWith(prefix []byte) []string {
	curr := t.root
	for i := range prefix {
		index := prefix[i] - 'a'
		if curr.children[index] == nil {
			return []string{}
		}
		curr = curr.children[index]
	}

	var result []string
	dfs(curr, prefix, &result)

	return result
}

func dfs(curr *Node, prefix []byte, result *[]string) {
	if len(*result) == 3 {
		return
	}

	if curr.isWord {
		*result = append(*result, string(prefix))
	}

	for i := 0; i < len(curr.children); i++ {
		if curr.children[i] == nil {
			continue
		}
		dfs(curr.children[i], append(prefix, byte(i+'a')), result)
	}
}
