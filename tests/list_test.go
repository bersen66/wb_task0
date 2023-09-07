package main

import (
	"fmt"
	"github.com/bersen66/wb_task0/internal/containers"
	"github.com/bersen66/wb_task0/pkg/entities"
	"testing"
)

func TestCreation(t *testing.T) {
	list := containers.NewLinkedList()
	if !list.Empty() {
		t.Errorf("list is not empty!")
	}
}

func TestPushFront(t *testing.T) {
	expected := make([]*entities.Order, 0, 10)
	list := containers.NewLinkedList()
	for i := 0; i < 10; i++ {
		order := entities.RandomOrder()
		list.PushFront(&order)
		expected = append(expected, &order)
	}
	var i = 9
	for it := list.Front(); it != nil; it = it.Next {
		t.Logf("it: %v exp: %v", it.Value.Uid, expected[i].Uid)
		if it.Value != expected[i] {
			t.Errorf("Value not inserted")
		}
		i--
	}

	if len(expected) != list.Size {
		t.Errorf("Sizes not equal")
	}
}

func TestPushBack(t *testing.T) {
	expected := make([]*entities.Order, 0, 10)
	list := containers.NewLinkedList()
	for i := 0; i < 10; i++ {
		order := entities.RandomOrder()
		list.PushBack(&order)
		expected = append(expected, &order)
	}
	var i = 0
	for it := list.Front(); it != nil; it = it.Next {
		t.Logf("it: %v exp: %v", it.Value.Uid, expected[i].Uid)
		if it.Value != expected[i] {
			t.Errorf("Value not inserted")
		}
		i++
	}

	if len(expected) != list.Size {
		t.Errorf("Sizes not equal")
	}
}

func TestIterationBackwards(t *testing.T) {
	expected := make([]*entities.Order, 0, 10)
	list := containers.NewLinkedList()
	for i := 0; i < 10; i++ {
		order := entities.RandomOrder()
		list.PushBack(&order)
		expected = append(expected, &order)
	}
	var i = 9
	for it := list.Back(); it != nil; it = it.Prev {
		t.Logf("it: %v exp: %v", it.Value.Uid, expected[i].Uid)
		if it.Value != expected[i] {
			t.Errorf("Value not inserted")
		}
		i--
	}
}

func TestSingle(t *testing.T) {
	list := containers.NewLinkedList()
	order := entities.RandomOrder()
	list.PushFront(&order)
	if list.Front().Value != &order && list.Back().Value != &order {
		t.Errorf("Single element not inserted")
	}
	if list.Size != 1 {
		t.Errorf("Size is not 1")
	}
}

func TestErase(t *testing.T) {
	list := containers.NewLinkedList()
	expected := make([]*entities.Order, 0, 10)
	for i := 0; i < 10; i++ {
		order := entities.RandomOrder()
		list.PushBack(&order)
		list.PushBack(&order)
		expected = append(expected, &order)
	}

	var i = 0
	for it := list.Front(); it != nil; {
		next := it.Next
		if i%2 == 0 {
			list.Erase(it)
		}
		i++
		it = next
	}

	i = 0
	for it := list.Front(); it != nil; it = it.Next {
		t.Logf("it: %v exp: %v", it.Value.Uid, expected[i].Uid)
		if it.Value != expected[i] {
			t.Errorf("Value not inserted")
		}
		i++
	}

}

func printList(list *containers.LinkedList) {
	fmt.Print("List: ")
	for it := list.Front(); it != nil; it = it.Next {
		fmt.Printf("%v, ", it.Value.Uid)
	}
	fmt.Println()
}

func TestMoveToFront(t *testing.T) {
	list := containers.NewLinkedList()

	for i := 0; i < 10; i++ {
		order := entities.RandomOrder()
		list.PushFront(&order)
	}
	fmt.Println("Has ")
	printList(list)

	for i := 0; i < 10; i++ {
		list.MoveToFront(list.Back())
	}

	printList(list)
}

func TestMoveToBack(t *testing.T) {
	list := containers.NewLinkedList()

	for i := 0; i < 10; i++ {
		order := entities.RandomOrder()
		list.PushBack(&order)
	}
	fmt.Println("Has ")
	printList(list)

	for i := 0; i < 10; i++ {
		list.MoveToBack(list.Front())
	}

	list.MoveToBack(list.Front().Next)

	printList(list)
}
