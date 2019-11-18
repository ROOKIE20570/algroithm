package main

import (
	"fmt"
)

type heap struct {
	limit   int //最大容量
	n       int
	content []int
}

func main() {
	h := newHeap(5)
	fmt.Println(h.content)
	h.insert(6)
	h.insert(4)
	h.insert(2)
	h.insert(3)
	h.insert(1)

	//h.insert(6)
	//h.insert(6)
	fmt.Println(h.content)
	h.remove()
	fmt.Println(h.content)
	h.remove()
	fmt.Println(h.content)
	h.remove()
	fmt.Println(h.content)
	h.remove()
	fmt.Println(h.content)
}

func newHeap(limit int) *heap {
	h := new(heap)
	h.limit = limit
	heapArr := make([]int, limit+1)
	h.n = 0
	h.content = heapArr
	return h
}
func (h *heap) insert(data int) {

	if h.n >= h.limit {
		return
	}
	h.n++

	h.content[h.n] = data

	for i := h.n; i/2 > 0 && h.content[i] > h.content[i/2]; {
		h.content[i], h.content[i/2] = h.content[i/2], h.content[i]
		i = i / 2
	}

}

func (h *heap) remove() {
	if (h.n == 0) {
		return
	}

	h.content[1] = h.content[h.n]
	h.content[h.n] = 0
	h.n--
	h.heapify(1)

}

//n当前堆内数据数量 i被堆化的位置
func (h *heap) heapify(i int) {

	for {
		if i*2 <= h.n && h.content[i] < h.content[i*2] {
			h.content[i], h.content[i*2] = h.content[i*2], h.content[i]
			i = i * 2
			continue
		} else if i*2+1 <= h.n && h.content[i] < h.content[i*2+1] {
			h.content[i], h.content[i*2+1] = h.content[i*2+1], h.content[i]
			i = i*2 + 1
			continue
		} else {
			break
		}

	}
}

func factorial(n uint) uint {
	if n == 1 {
		return 1
	}

	return factorial(n-1) * n
	//f(n) = f(n-1)*n
}
