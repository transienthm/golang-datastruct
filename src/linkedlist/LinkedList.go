package linkedlist

import "fmt"

type Elem int

type Node struct {
	Date Elem
	Next *Node
}

func New() *Node {
	return &Node{0, nil}
}

func (head *Node) Insert(i int, e Elem) bool {
	p := head
	j := 1

	for p != nil && j < i  {
		p = p.Next
		j++
	}

	if p == nil || j > i {
		fmt.Println("pls check i", i)
		return false
	}

	s := &Node{e, nil}
	s.Next = p.Next
	p.Next = s
	return true
}

func (head *Node) Traverse() {
	p := head.Next
	for nil != p {
		fmt.Println(p.Date)
		p = p.Next
	}
	fmt.Println("-------done--------")
}

func (head *Node) Get(i int) Elem {
	p := head.Next
	for j := 1; j < i; j++ {
		if nil == p {
			return -1
		}
		p = p.Next
	}

	return p.Date
}

func (head *Node) Delete(i int) bool {
	p := head
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}

	if nil == p || j > i {
		fmt.Println("pls check i: ", i)
		return false
	}

	p.Next = p.Next.Next
	return true
}

