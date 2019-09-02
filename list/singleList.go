package main

import "fmt"

type listNode struct {
	next *listNode
	val  int
}

type linkedList struct {
	head   *listNode
	length uint
}

func main() {
	list := initList()
	list.insertAfter(list.head, 5)
	list.insertAfter(list.head, 6)
	list.insertAfter(list.head, 7)
	list.insertAfter(list.head, 8)

	list.printList()
	list.reserver()
	list.printList()

}

func initList() *linkedList {
	return &linkedList{newListNode(0), 0}
}

func newListNode(val int) *listNode {
	return &listNode{nil, val}
}

func (list *linkedList) insertAfter(node *listNode, val int) bool {
	if node == nil {
		return false
	}
	newNode := newListNode(val)

	if node.next != nil {
		newNode.next = node.next
	}

	node.next = newNode

	return true

}

func (list *linkedList) printList() {
	cur := list.head.next
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

func (list *linkedList) reserver() {
	if nil == list.head || nil == list.head.next || nil == list.head.next.next {
		return
	}

	var pre *listNode = nil //暂存前一个结点

	cur := list.head.next //暂存当前遍历到的结点

	for cur != nil {
		//从后向前摆放
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	list.head.next = pre
}


