package main

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	head *node
	len  int
}

func main() {
	l := newList()
	insertToTail(l, 2)
	insertToTail(l, 3)
	insertToTail(l, 4)
	insertToTail(l, 5)

	l1 := newList()
	insertToTail(l1, 1)
	insertToTail(l1, 5)
	insertToTail(l1, 7)
	insertToTail(l1, 9)

	//deleteBottomN(l, 0)
	ls := mergeSortedList(l,l1)
	printList(ls)

}

func newList() *list {
	newList := new(list)
	newList.head = newNode(0)
	newList.len = 0

	return newList
}

func newNode(data int) *node {
	newNode := new(node)
	newNode.data = data
	return newNode
}

func insertToTail(myList *list, data int) *node {
	p := myList.head
	for p.next != nil {
		p = p.next
	}

	n := newNode(data)
	p.next = n

	return n
}

func insertAfter(listNode *node, data int) {
	if listNode == nil {
		return
	}

	tmp := listNode.next
	n := newNode(data)
	listNode.next = n
	n.next = tmp

}

func isCircle(l *list) bool { //不熟 ***
	if l.head == nil || l.head.next == nil || l.head.next.next == nil {
		return false
	}

	fast := l.head.next
	slow := l.head.next

	for fast.next != nil && fast.next.next != nil {
		if fast == slow {
			return true
		}

		fast = fast.next.next
		slow = slow.next
	}

	return false
}

func deleteNode(l *list, n *node) {
	h := l.head
	for h != nil {
		if h.next == n {
			h.next = n.next
			return
		}
		h = h.next
	}
}
func deleteBottomN(l *list, n int) {
	if l.head == nil {
		return
	}
	if n <= 0 {
		return
	}

	slow := l.head.next
	fast := l.head.next

	for i := 0; i < n; i++ {
		if fast == nil {
			return
		}
		fast = fast.next
	}

	for fast != nil {
		fast = fast.next
		slow = slow.next
	}

	deleteNode(l, slow)
}

func reserver(l *list) {
	if l.head == nil || l.head.next == nil || l.head.next.next == nil {
		return
	}
	var pre *node = nil

	cur := l.head.next

	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	l.head.next = pre

}

func mergeSortedList(l1, l2 *list) *list {
	if l1.head == nil || l1.head.next == nil {
		return l2
	}

	if l2.head == nil || l2.head.next == nil {
		return l1
	}

	p1 := l1.head.next
	p2 := l2.head.next
	l3 := newList()
	p3 := l3.head
	for p1 != nil && p2 != nil {
		if p1.data < p2.data {
			p3.next = p1
			p1 = p1.next
		} else {
			p3.next = p2
			p2 = p2.next
		}

		p3 = p3.next

	}

	if p1 != nil {
		p3.next = p1
	}

	if p2 !=nil {
		p3.next = p2
	}

	return l3
}

func printList(myList *list) {
	if myList.head == nil || myList.head.next == nil {
		return
	}

	p := myList.head
	for p.next != nil {
		fmt.Println(p.next.data)
		p = p.next
	}
}
