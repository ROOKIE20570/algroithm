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
}

func newHeap(limit int) *heap {
	h := new(heap)
	h.limit = limit
	heapArr := []int{}
	h.content = heapArr
	return h
}
func (h *heap) insert(data int) {
	if h.n >= h.limit {
		return
	}

	h.n++
	h.content[h.n] = data
	i := h.n
	for ; i/2 > 0 && h.content[i] > h.content[i/2]; {
		h.content[i], h.content[i/2] = h.content[i/2], h.content[i]
		i = i / 2
	}

}

func (h *heap) remove() {
	if (h.n == 0) {
		return
	}

	h.content[1] = h.content[h.n]
	h.n--
	heapify(h.content,h.n,1)
}

func heapify(arr []int, n, i int) {
	for {
		maxPos := i
		if i*2 <= n && arr[maxPos] < arr[i*2] {
			maxPos = i * 2
		}

		if i*2+1 <= n && arr[maxPos] < arr[i*2+1] {
			maxPos = i*2 + 1
		}

		if maxPos == i {
			break
		}

		arr[i],arr[maxPos] = arr[maxPos],arr[i]

		i = maxPos
	}
}
