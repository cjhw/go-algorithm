package main

type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Val  int
	Pre  *Node
	Next *Node
}

func Constructor() MyLinkedList {
	rear := &Node{
		Val:  -1,
		Pre:  nil,
		Next: nil,
	}

	rear.Next = rear
	rear.Pre = rear
	return MyLinkedList{rear}
}

func (this *MyLinkedList) Get(index int) int {
	head := this.dummy.Next
	for head != this.dummy && index > 0 {
		head = head.Next
		index--
	}

	if index != 0 {
		return -1
	}

	return head.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	dummy := this.dummy
	node := &Node{
		Val:  val,
		Pre:  dummy,
		Next: dummy.Next,
	}

	dummy.Next.Pre = node
	dummy.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {

	dummy := this.dummy
	node := &Node{
		Val:  val,
		Pre:  dummy.Pre,
		Next: dummy,
	}

	dummy.Pre.Next = node
	dummy.Pre = node

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

	head := this.dummy.Next

	for head != this.dummy && index > 0 {
		head = head.Next
		index--
	}

	if index > 0 {
		return
	}

	node := &Node{
		Val:  val,
		Pre:  head.Pre,
		Next: head,
	}

	head.Pre.Next = node
	head.Pre = node

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	head := this.dummy.Next
	if head == this.dummy {
		return
	}

	for head.Next != this.dummy && index > 0 {
		head = head.Next
		index--
	}

	if index == 0 {
		head.Next.Pre = head.Pre
		head.Pre.Next = head.Next
	}
}
