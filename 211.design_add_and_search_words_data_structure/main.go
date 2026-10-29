package main

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	for _, r := range word {
		index := r - 'a'
		if this.children[index] == nil {
			this.children[index] = &WordDictionary{}
		}
		this = this.children[index]
	}
	this.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(*WordDictionary, int) bool
	dfs = func(wd *WordDictionary, index int) bool {
		if index == len(word) {
			return wd.isEnd
		}

		if word[index] != '.' {
			i := word[index] - 'a'
			if wd.children[i] == nil {
				return false
			}

			return dfs(wd.children[i], index+1)
		}

		for _, child := range wd.children {
			if child != nil && dfs(child, index+1) {
				return true
			}
		}

		return false
	}

	return dfs(this, 0)
}

type WordDictionary2 struct {
	nodes map[rune]*WordDictionary2
	isEnd bool
}

func Constructor2() WordDictionary2 {
	return WordDictionary2{
		nodes: make(map[rune]*WordDictionary2),
		isEnd: false,
	}
}

func (this *WordDictionary2) AddWord(word string) {
	for _, r := range word {
		if _, found := this.nodes[r]; !found {
			node := Constructor2()
			this.nodes[r] = &node
		}

		this = this.nodes[r]
	}
	this.isEnd = true
}

func (this *WordDictionary2) Search(word string) bool {
	for i, r := range word {
		if _, found := this.nodes[r]; !found {
			if r != '.' {
				return false
			}

			for _, node := range this.nodes {
				if node.Search(word[i+1:]) {
					return true
				}
			}

			return false
		}

		this = this.nodes[r]
	}

	return this.isEnd
}
