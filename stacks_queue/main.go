package main

import (
	"fmt"
)

// Queue structure
type Queue struct {
	data []int
}

// Stack structure
type Stack struct {
	data []int
}

// pop will delete the last element added to the stack
func (s *Stack) pop() {
	if len(s.data) == 0 {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

// push adds new elements to the stack
func (s *Stack) push(k int) {
	s.data = append(s.data, k)
}

// enqueue adds new elements to the queue
func (q *Queue) enqueue(k int) {
	q.data = append(q.data, k)
}

// dequeue deletes the first element in the queue
func (q *Queue) dequeue() {
	if len(q.data) == 0 {
		return
	}
	q.data = q.data[1:len(q.data)]
}

func main() {
	testStack := &Stack{}
	testQueue := &Queue{}
	list := []int{1, 2, 3, 4, 5, 6}

	for _, v := range list {
		testStack.push(v)
		testQueue.enqueue(v)
	}
	fmt.Println(testStack)
	fmt.Println(testQueue)
	testStack.pop()
	testQueue.dequeue()
	testQueue.dequeue()
	fmt.Println(testStack)
	fmt.Println(testQueue)

}
