package main

type queue struct {
	length int
	data   []int
}

func main() {
	myQueue := newQueue(5)

}

func newQueue(length int) *queue {
	myQueue := new(queue)
	data := make([]int, length)
	myQueue.data = data
	myQueue.length = length
	return myQueue
}
func push(val int) {

}

func pop() {

}
