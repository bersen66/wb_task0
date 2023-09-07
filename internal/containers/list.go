package containers

import (
	"github.com/bersen66/wb_task0/pkg/entities"
)

type Node struct {
	Prev  *Node
	Next  *Node
	Value *entities.Order
}

type LinkedList struct {
	head *Node
	tail *Node
	Size int
}

func NewLinkedList() *LinkedList {
	result := new(LinkedList)
	return result
}

func (list *LinkedList) Front() *Node {
	return list.head
}

func (list *LinkedList) Back() *Node {
	return list.tail
}

func (list *LinkedList) InsertBefore(node *Node, value *entities.Order) *Node {
	newNode := new(Node)

	newNode.Value = value

	newNode.Prev = node.Prev
	newNode.Next = node

	if node.Prev != nil {
		node.Prev.Next = newNode
	}

	node.Prev = newNode
	list.Size++
	return newNode
}

func (list *LinkedList) InsertAfter(node *Node, value *entities.Order) *Node {
	newNode := new(Node)

	newNode.Value = value

	newNode.Prev = node
	newNode.Next = node.Next

	if node.Next != nil {
		node.Next.Prev = newNode
	}
	node.Next = newNode

	list.Size++
	return newNode
}

func (list *LinkedList) Empty() bool {
	return list.Size == 0 && list.head == nil && list.tail == nil
}

func (list *LinkedList) insertToEmpty(value *entities.Order) {
	newNode := new(Node)
	newNode.Value = value
	list.head = newNode
	list.tail = newNode
	list.Size++
}

func (list *LinkedList) PushBack(value *entities.Order) *Node {
	if list.Empty() {
		list.insertToEmpty(value)
	} else {
		list.tail = list.InsertAfter(list.tail, value)
	}
	return list.tail
}

func (list *LinkedList) PushFront(value *entities.Order) *Node {
	if list.Empty() {
		list.insertToEmpty(value)
	} else {
		list.head = list.InsertBefore(list.head, value)
	}
	return list.head
}

func (list *LinkedList) MoveToFront(node *Node) {
	if node == list.head {
		return
	}

	node.Prev.Next = node.Next
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node == list.tail {
		list.tail = node.Prev
	}
	node.Prev = nil
	list.head.Prev = node
	node.Next = list.head

	list.head = node
}

func (list *LinkedList) MoveToBack(node *Node) {
	if node == list.tail {
		return
	}

	node.Next.Prev = node.Prev
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node == list.head {
		list.head = node.Next
	}

	node.Next = nil
	list.tail.Next = node
	node.Prev = list.tail
	list.tail = node
}

func (list *LinkedList) Erase(node *Node) {
	switch {
	case list.Size == 1:
		{
			list.head = nil
			list.tail = nil
		}
	case node == list.head:
		{
			list.head = node.Next
			list.head.Prev = nil
		}
	case node == list.tail:
		{
			list.tail = node.Prev
			list.tail.Next = nil
		}
	default:
		{
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
		}

	}
	list.Size--
	node.Next = nil
	node.Prev = nil
}
