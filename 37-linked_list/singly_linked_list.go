package main

import "fmt"

type SinglyLinkedList struct {
	head *Node
	len  int
}

func NewSinglyLinkedList(head *Node) SinglyLinkedList {
	var l int
	if head != nil {
		l = 1
	}
	return SinglyLinkedList{
		head: head,
		len:  l,
	}
}

// IsEmpty Determines if the linked list is empty
// Takes O(1) time
func (s *SinglyLinkedList) IsEmpty() bool {
	return s.head == nil
}

// Len Returns the length of the linked list
// Takes O(1) time
func (s *SinglyLinkedList) Len() int {
	return s.len
}

// Add Adds new Node containing data to head of the list
// Also called prepend
// Takes O(1) time
func (s *SinglyLinkedList) Add(data int) {
	s.head = NewNode(data, s.head)
	s.len++
}

// Search for the first node containing data that matches the key
// Returns the node or `nil` if not found
// Takes O(n) time
func (s *SinglyLinkedList) Search(data int) *Node {
	current := s.head
	for current != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}

	return nil
}

// Insert a new Node containing data at index position
// Insertion takes O(1) time but finding node at insertion point takes
// O(n) time.
// Takes overall O(n) time.
func (s *SinglyLinkedList) Insert(data, index int) {
	if index < 0 || index > s.len {
		panic("Index out of range")
	}

	if index == 0 {
		s.Add(data)
		return
	}

	current := s.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	newNode := NewNode(data, current.next)
	current.next = newNode
	s.len++
}

// NodeAt Returns the Node at specified index
// Takes O(n) time
func (s *SinglyLinkedList) NodeAt(index int) *Node {
	if index < 0 || index > s.len {
		panic("Index out of range")
	}

	current := s.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current
}

// Remove Node containing data that matches the key
// Returns the node or `nil` if key doesn't exist
// Takes O(n) time
func (s *SinglyLinkedList) Remove(key int) *Node {
	if s.IsEmpty() {
		return nil
	}

	if s.head.data == key {
		s.head = s.head.next
		s.len--
		return nil
	}

	current := s.head
	for current.next != nil {
		if current.next.data == key {
			current.next = current.next.next
			s.len--
			return current
		}
		current = current.next
	}

	return nil
}

// RemoveAt Removes Node at specified index
// Takes O(n) time
func (s *SinglyLinkedList) RemoveAt(index int) *Node {
	if index < 0 || index > s.len {
		panic("Index out of range")
	}

	if index == 0 {
		s.head = s.head.next
		s.len--
		return nil
	}

	current := s.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	current.next = current.next.next
	s.len--

	return current
}

// Print the linked list
// Takes O(n) time
func (s *SinglyLinkedList) Print() {
	current := s.head
	for current != nil {
		if current == s.head {
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
