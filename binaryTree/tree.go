package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

var tree *Node


func main() {
	insert(5)
	insert(6)
	insert(1)
	insert(3)
	insert(8)
	insert(7)
	insert(12)
	insert(2)
	preOrder(tree)
	
}

func insert(data int) {
	if tree == nil {
		tree = new(Node)
		tree.data = data
		return
	}

	p := tree
	for p != nil {
		if data > p.data {
			if p.right == nil {
				p.right = NewNode(data)
				break
			}


			p = p.right
		} else {
			if p.left == nil {
				p.left = NewNode(data)
				break
			}
			p = p.left
		}
	}

}

func preOrder(root* Node)  {
	if root == nil{
		return
	}

	p := root
	fmt.Println(p.data)
	preOrder(p.left)
	preOrder(p.right)
}
func del(){

}

