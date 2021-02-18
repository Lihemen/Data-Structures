package main

import "fmt"

// AlphabetsSize is the total number of possible characters
const AlphabetsSize = 26

// Node struct
type Node struct {
	children [AlphabetsSize]*Node
	isEnd    bool
}

// Trie struct
type Trie struct {
	root *Node
}

// InitTrie initializes a trie structure
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert method adds a new Node to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		// the index of a is 97. The charIndex basically returns the subtraction of the ASCII value of any letter in the word with a (i.e 'b' - 'a' = 1; because 'b' is 98 and 'a' is 97)
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			// Create a node and assign to index
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search method will take a word and return true/false if the world
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	testTrie := InitTrie()
	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"me",
		"oreo",
		"oregano",
		"meat",
		"met",
	}
	for _, v := range toAdd {
		testTrie.Insert(v)
	}

	fmt.Println(testTrie.Search("oreg"))
	fmt.Println(testTrie.Search("oregon"))
	fmt.Println(testTrie.Search("arg"))
	fmt.Println(testTrie.Search("argon"))
}
