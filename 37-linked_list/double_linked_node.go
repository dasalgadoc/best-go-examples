package main

import "fmt"

type DoubleLinkedNode struct {
	data int
	prev *DoubleLinkedNode
	next *DoubleLinkedNode
}

func NewDoubleLinkedNode(data int, prevNode, nextNode *DoubleLinkedNode) *DoubleLinkedNode {
	return &DoubleLinkedNode{
		data: data,
		prev: prevNode,
		next: nextNode,
	}
}

func (n *DoubleLinkedNode) Data() string {
	return fmt.Sprintf("<Node data: %d>", n.data)
}
