package main

import (
	"container/list"
	"fmt"
)

type gragh struct {
	vertex int
	data   []*list.List
}

func main() {
	myGragh := newGragh(5)
	fmt.Println(myGragh)
}

func newGragh(v int) *gragh {
	myGragh := new(gragh)
	myGragh.vertex = v
	myGragh.data = make([]*list.List, v)
	for i := 0; i < v; i++ {
		myGragh.data[i] = new(list.List)
	}
	return myGragh
}

func (myGragh *gragh)addEdge(s, t int)  {
	myGragh.data[s].PushBack(t)
	myGragh.data[t].PushBack(s)
}

func bfs(s, t int)  {

}




