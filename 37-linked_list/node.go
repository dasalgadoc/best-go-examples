package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func NewNode(data int, nextNode *Node) *Node {
	return &Node{
		data: data,
		next: nextNode,
	}
}

func (n *Node) Data() string {
	return fmt.Sprintf("<Node data: %d>", n.data)
}
