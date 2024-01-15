package main

type DoubleLinkedList struct {
	head *DoubleLinkedNode
	len  int
}

func NewDoubleLinkedList(head *DoubleLinkedNode) DoubleLinkedList {
	var l int
	if head != nil {
		l = 1
	}
	return DoubleLinkedList{
		head: head,
		len:  l,
	}
}

// IsEmpty Determines if the linked list is empty
// Takes O(1) time
func (d *DoubleLinkedList) IsEmpty() bool {
	return d.head == nil
}

// Len Returns the length of the linked list
// Takes O(1) time
func (d *DoubleLinkedList) Len() int {
	return d.len
}
