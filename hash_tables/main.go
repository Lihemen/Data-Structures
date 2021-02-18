package main

// ArraySize is the size of the hash table
const ArraySize = 7

// HashTable structure designed using the seperate chaining method
type HashTable struct {
	array [ArraySize]*Bucket
}

// Bucket structure
type Bucket struct {
	head *bucketNode
}

// bucketNode is a linked list that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// HashTable Functions

// Insert will take a key and add it to the hash table
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Delete will take a key and remove it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Search will take a key and search for it in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Bucket
// Insert
func (b *Bucket) insert(k string) {
	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
}

// delete will take a key and delete the node from the bucket
func (b *Bucket) delete(k string) {
	// If the head is the key we are looking for, we simply set the head to the next node
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// search will tke a key and return a boolean
func (b *Bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// hash takes a key and returns a hash code
func hash(key string) int {
	// Get the sum of each ascii character
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will initialize the hash table and create a bucket in each slot
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

func main() {
	hashTable := Init()

	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	for _, v := range list {
		// To really see the power of the hash table, put a breakpoint in the line below and step over into the line continuously
		hashTable.Insert(v)
	}
}
