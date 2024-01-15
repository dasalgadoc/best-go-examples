package main

import "fmt"

func main() {
	singlyLinkedList()
	doublyLinkedList()
}

func singlyLinkedList() {
	fmt.Printf("Singly Linked List\n")
	// Create a new linked list
	ll := NewSinglyLinkedList(nil)

	// Add some values to the linked list
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)

	// Print the linked list
	ll.Print()

	// Remove some values from the linked list
	ll.Remove(1)
	ll.Remove(3)
	ll.Remove(5)

	// Print the linked list
	ll.Print()
}

func doublyLinkedList() {
	fmt.Printf("Doubly Linked List\n")
	// Create a new linked list
	ll := NewDoubleLinkedList(nil)

	// Add some values to the linked list
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Add(5)

	// Print the linked list
	ll.Print()

	// Remove some values from the linked list
	ll.Remove(1)
	ll.Remove(3)
	ll.Remove(5)

	// Print the linked list
	ll.Print()
}
