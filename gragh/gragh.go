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

func (myGragh *gragh) addEdge(s, t int) {
	myGragh.data[s].PushBack(t)
	myGragh.data[t].PushBack(s)
}

//思路 visted 已经遍历过的顶点 prev[] 记录路径   queue 存储前一层的结点
func (myGragh *gragh) bfs(s, t int) {
	//一千个伤心的理由
	if s == t {
		return
	}

	visted := make([]bool, myGragh.vertex)
	//队列  记录的
	stash := make([]int, myGragh.vertex)
	prev := make([]int, myGragh.vertex)
	//
	for index := range prev {
		prev[index] = -1
	}

	stash = append(stash, s)

	visted[s] = true
	isFound := false
	for len(stash) > 0 {
		t := stash[0]
		stash = stash[1:]
		cur := myGragh.data[t]
		for e := cur.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visted[k] {
				prev[k] = t
				if k == t {
					isFound = true
					break
				}

				stash = append(stash,k)
				visted[k] = true
			}
		}

	}

	if isFound {
		printPrev(prev,s,t)
	}else {
		fmt.Println("error")
	}
	//visted[s] = true

}

func printPrev(prev []int, s int, t int) {

	if t == s || prev[t] == -1 {
		fmt.Printf("%d ", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d ", t)
	}

}