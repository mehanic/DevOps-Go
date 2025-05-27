package main

import "fmt"

type Queue struct {
	enqueueStack []int
	dequeueStack []int
}

func NewQueue() *Queue {
	return &Queue{
		enqueueStack: []int{},
		dequeueStack: []int{},
	}
}

func (q *Queue) Enqueue(x int) {
	q.enqueueStack = append(q.enqueueStack, x)
}

func (q *Queue) transferEnqueueToDequeue() {
	if len(q.dequeueStack) == 0 {
		for len(q.enqueueStack) > 0 {
			// Возьмём последний элемент из enqueueStack
			n := len(q.enqueueStack) - 1
			elem := q.enqueueStack[n]
			q.enqueueStack = q.enqueueStack[:n]

			// Поместим его в dequeueStack
			q.dequeueStack = append(q.dequeueStack, elem)
		}
	}
}

func (q *Queue) Dequeue() (int, bool) {
	q.transferEnqueueToDequeue()
	if len(q.dequeueStack) == 0 {
		return 0, false // очередь пуста
	}
	n := len(q.dequeueStack) - 1
	elem := q.dequeueStack[n]
	q.dequeueStack = q.dequeueStack[:n]
	return elem, true
}

func (q *Queue) Peek() (int, bool) {
	q.transferEnqueueToDequeue()
	if len(q.dequeueStack) == 0 {
		return 0, false
	}
	return q.dequeueStack[len(q.dequeueStack)-1], true
}

func main() {
	q := NewQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	val, ok := q.Peek()
	if ok {
		fmt.Println("Peek:", val) // Peek: 10
	}

	val, ok = q.Dequeue()
	if ok {
		fmt.Println("Dequeue:", val) // Dequeue: 10
	}

	val, ok = q.Peek()
	if ok {
		fmt.Println("Peek:", val) // Peek: 20
	}
}
