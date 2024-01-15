package main

import "fmt"

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

// Add new Node containing data to head of the list
// Also called prepend
// Takes O(1) time
func (d *DoubleLinkedList) Add(data int) {
	if d.IsEmpty() {
		d.head = NewDoubleLinkedNode(data, nil, nil)
	} else {
		newNode := NewDoubleLinkedNode(data, nil, d.head)
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
}

// Search for the first node containing data that matches the key
// Returns the node or `nil` if not found
// Takes O(n) time
func (d *DoubleLinkedList) Search(data int) *DoubleLinkedNode {
	current := d.head
	for current != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}

	return nil
}

// NodeAt Returns the Node at specified index
// Takes O(n) time
func (d *DoubleLinkedList) NodeAt(index int) *DoubleLinkedNode {
	if index >= d.len {
		return nil
	}

	current := d.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current
}

// Insert a new Node containing data at index position
// Insertion takes O(1) time but finding node at insertion point takes
// O(n) time.
// Takes overall O(n) time.
func (d *DoubleLinkedList) Insert(data, index int) {
	if index < 0 || index > d.len {
		panic("Index out of range")
	}

	if index == 0 {
		d.Add(data)
		return
	}

	current := d.NodeAt(index)
	previous := current.prev
	newNode := NewDoubleLinkedNode(data, previous, current)
	current.prev = newNode
	previous.next = newNode
	d.len++
}

// Remove Node containing data that matches the key
// Returns the node or `nil` if key doesn't exist
// Takes O(n) time
func (d *DoubleLinkedList) Remove(data int) {
	if d.IsEmpty() {
		panic("Cannot remove from an empty list")
	}

	if d.head.data == data {
		d.head = d.head.next
		d.len--
		return
	}

	current := d.head
	for current.next != nil {
		if current.data == data {
			current.prev.next = current.next
			current.next.prev = current.prev
			d.len--
			return
		}
		current = current.next
	}
}

// RemoveAt removes Node at specified index
// Takes O(n) time
func (d *DoubleLinkedList) RemoveAt(index int) {
	if index < 0 || index > d.len {
		panic("Index out of range")
	}

	if index == 0 {
		d.head = d.head.next
		d.len--
		return
	}

	current := d.NodeAt(index)
	previous := current.prev
	previous.next = current.next
	current.next.prev = previous
	d.len--
}

// Print a string representation of the list.
// Takes O(n) time.
func (d *DoubleLinkedList) Print() {
	current := d.head
	for current != nil {
		if current == d.head {
			fmt.Printf("[Head: %d]", current.data)
		} else if current.next == nil {
			fmt.Printf(" -> [Tail: %d]", current.data)
		} else {
			fmt.Printf(" -> [%d]", current.data)
		}
		current = current.next
	}
	fmt.Println()
}
