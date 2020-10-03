package main

type Node struct {
	children map[rune]*Node
	ch       rune
	word     rune
	parent   *Node

	isLeaf bool
}

func NewNode() *Node {
	return &Node{children: make(map[rune]*Node)}
}
