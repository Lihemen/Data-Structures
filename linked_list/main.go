package main

import (
	"fmt"
)

// The following implementation is the implementation of a singly linked list

// LinkedList structure
type LinkedList struct {
	head   *Node
	length int
}

// Node structure
type Node struct {
	key  int
	next *Node
}

// insert will take a key and create a new node
func (l *LinkedList) insert(k int) {
	newNode := &Node{key: k}
	newNode.next = l.head
	l.head = newNode
	l.length++
}

// search will traverse through the list and return true/false if the key exists or not
func (l *LinkedList) search(k int) bool {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete will take a key and remove its node from the list by pointing the head to the next node
func (l *LinkedList) delete(k int) {
	if l.head.key == k {
		l.head = l.head.next
		l.length = l.length - 1
		return
	}
	previousNode := l.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete by pointing the node to the next node thereby skipping the node you want to delete
			previousNode.next = previousNode.next.next
			l.length = l.length - 1
		}
		previousNode = previousNode.next
	}
}

func main() {
	list := &LinkedList{}
	fmt.Println(list)
	fmt.Println(list.length)
	data := []int{
		1,
		2,
		3,
		4,
		5,
		6,
	}
	for _, v := range data {
		list.insert(v)
	}
	fmt.Println(list)
	list.delete(3)
	fmt.Println(list)
	fmt.Println(list.search(8))
	fmt.Println(list.search(4))
	fmt.Println(list.search(0))
}
