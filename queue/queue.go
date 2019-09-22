package main

import "fmt"

type queue struct {
	cap  int
	data []int
	head int
	tail int
}

func main() {
	myQueue := newQueue(5)

}

func newQueue(cap int) *queue {
	myQueue := new(queue)
	data := make([]int, cap)
	myQueue.data = data
	myQueue.cap = cap
	return myQueue
}
func (myQueue *queue) enQueue(val int) bool {
	if myQueue.tail == myQueue.cap {
		return false
	}

	myQueue.data[myQueue.tail] = val
	myQueue.tail++
	return true
}

func (myQueue *queue) deQueue() int {
	if myQueue.head == myQueue.tail {
		return 0
	}

	v := myQueue.data[myQueue.head]
	myQueue.head++
	return v
}

func (myQueue *queue) print() {
	if myQueue.head == myQueue.tail {
		fmt.Println( "empty queue")
	}
	result := "head"
	for i := myQueue.head; i <= myQueue.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", myQueue.q[i])
	}
	result += "<-tail"
	fmt.Println(result)
}
