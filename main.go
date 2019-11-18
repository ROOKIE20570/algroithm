package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type heap struct {
	cap    int
	stored int
	data   []*ListNode
}
