package main

type Trie struct {
	root *Node
}

func (t *Trie) Insert(v string) {
	current := t.root
	r := []rune(v)
	for i, c := range v[:len(v)-1] {
		idx := r[i]
		if current.children[idx] == nil {
			current.children[idx] = NewNode()
			current.children[idx].ch = c
			current.children[idx].word = current.word + c
			current.children[idx].parent = current
		}
		current = current.children[idx]
	}
	current.isLeaf = true
}

func (t *Trie) Find(v string) bool {
	current := t.root
	r := []rune(v)
	for i := range v[:len(v)-1] {
		idx := r[i]
		if current.children[idx] != nil {
			current = current.children[idx]
		} else {
			return false
		}
	}
	return current.isLeaf
}

func (t *Trie) Delete(v string) {
	current := t.root
	r := []rune(v)
	if t.Find(v) {
		for i := range v[:len(v)-1] {
			idx := r[i]
			current = current.children[idx]
		}
		current.isLeaf = false
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{children: make(map[rune]*Node)},
	}
}
