package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert will push a new node onto the tree if the node doesn't exist
func (n *Node) Insert(key int) {
	if n.Key < key {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else if n.Key > key {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	}
}

//Search will return true if the key is found else false
func (n *Node) Search(key int) bool {
	if n == nil { // This means we have arrived at the end of the tree
		return false
	}
	if n.Key < key {
		// move right
		return n.Right.Search(key)
	} else if n.Key > key {
		// move left
		return n.Left.Search(key)
	}
	return true
}

func main() {
	bst := &Node{Key: 100}
	bst.Insert(500)
	bst.Insert(200)
	bst.Insert(300)
	fmt.Println(bst)
	fmt.Println(bst.Search(200))
}
