package main

import "fmt"

// Maxheap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	h.maxHeapifyDown(0)

	return extracted
}

// maxheapify will heapify from bottom to top (bubble up)
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxheapifydown will heapify from top to bottom (bubble down)
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastindex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastindex {
		if l == lastindex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else {
			childToCompare = r // right child is larger
		}
		// compare and swap
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// Get the left child index
func left(i int) int {
	return 2*i + 1
}

// Get the right child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {

	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
	defer fmt.Println(*m)
}
