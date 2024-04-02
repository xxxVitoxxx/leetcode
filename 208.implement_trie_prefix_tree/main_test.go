package main

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Error("should be search apple")
	}
	if trie.Search("app") {
		t.Error("should not be search app")
	}
	if !trie.StartsWith("app") {
		t.Error("should be search prefix")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Error("should be search app")
	}
}

func TestTrie2(t *testing.T) {
	trie := Constructor2()
	trie.Insert2("apple")
	if !trie.Search2("apple") {
		t.Error("should be search apple")
	}
	if trie.Search2("app") {
		t.Error("should not be search app")
	}
	if !trie.StartsWith2("app") {
		t.Error("should be search prefix")
	}
	trie.Insert2("app")
	if !trie.Search2("app") {
		t.Error("should be search app")
	}
}
