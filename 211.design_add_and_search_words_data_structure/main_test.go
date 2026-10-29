package main

import "testing"

func TestWordDictionary(t *testing.T) {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	if obj.Search("pad") {
		t.Errorf("'pad' not exist in dictionary")
	}
	if !obj.Search("bad") {
		t.Errorf("'bad' is exist in dictionary")
	}
	if !obj.Search(".ad") {
		t.Errorf("'.ad' is exist in dictionary")
	}
	if !obj.Search("b..") {
		t.Errorf("'b..' is exist in dictionary")
	}
}

func TestWordDictionary2(t *testing.T) {
	obj := Constructor2()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	if obj.Search("pad") {
		t.Errorf("'pad' not exist in dictionary")
	}
	if !obj.Search("bad") {
		t.Errorf("'bad' is exist in dictionary")
	}
	if !obj.Search(".ad") {
		t.Errorf("'.ad' is exist in dictionary")
	}
	if !obj.Search("b..") {
		t.Errorf("'b..' is exist in dictionary")
	}
}
