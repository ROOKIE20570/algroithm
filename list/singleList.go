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

	list1 := initList()
	list1.insertAfter(list1.head, 3)
	list1.insertAfter(list1.head, 6)
	list1.insertAfter(list1.head, 11)
	list1.insertAfter(list1.head, 12)

	MergeSortedList(list1,list).printList()


}

func initList() *linkedList {
	return &linkedList{newListNode(0), 0}
}

func newListNode(val int) *listNode {
	return &listNode{nil, val}
}

func (link *linkedList)newEndNode(val int) *listNode {
	return &listNode{link.head, val}
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
		//首先把下一个要操作的结点记下来
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	list.head.next = pre
}

func (list *linkedList) HasCycle() bool {
	if nil != list.head {
		fast, slow := list.head, list.head

		for fast != nil && fast.next != nil {
			fast = fast.next.next
			slow = slow.next

			if slow == fast {
				return true
			}
		}
	}

	return false
}

func (this *linkedList) deleteBottomN(n int) {
	if n <= 0 || nil == this.head || nil == this.head.next {
		return
	}

	fast := this.head
	for i := 0; fast != nil && i < n; i++ {
		fast = fast.next
	}

	if nil == fast{
		return
	}

	slow := this.head
	for fast.next != nil  {
		fast = fast.next
		slow = slow.next
	}

	slow.next = slow.next.next	

}

func MergeSortedList(l1, l2 *linkedList) *linkedList {
	if nil == l1 || nil == l1.head || nil == l1.head.next {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}

	l := &linkedList{head: &listNode{}}
	cur := l.head
	curl1 := l1.head.next
	curl2 := l2.head.next
	for nil != curl1 && nil != curl2 {
		if curl1.val < curl2.val {
			cur.next = curl2
			curl2 = curl2.next
		} else {
			cur.next = curl1
			curl1 = curl1.next
		}
		cur = cur.next
	}

	if nil != curl1 {
		cur.next = curl1
	} else if nil != curl2 {
		cur.next = curl2
	}



	return l
}

