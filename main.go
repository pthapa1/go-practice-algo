package main

import (
	"fmt"
	// algo "github.com/pthapa1/go-practice-algo/algorithms"
	"reflect"
)

type ReturnType[T any] struct {
	Data  interface{}
	Error error
}

type Node[T any] struct {
	data T
	next *Node[T]
}

type SinglyLinkedList[T any] struct {
	head   *Node[T]
	length int
}

// Add new node on the head of linked list. Contant Time O(1)
func (l *SinglyLinkedList[T]) prepend(n Node[T]) {
	oldHead := l.head // save current head as old head
	l.head = &n       // make n our new head.
	n.next = oldHead  // and point it to our old head
	l.length++        // increase length for prepend
}

// Print all data associated with Singly Linked List
func (l SinglyLinkedList[T]) printAllSLLData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%v, --> ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

// delete a node given it's value.
func (l *SinglyLinkedList[T]) deleteNode(value T) ReturnType[T] {

	if l.length == 0 {
		return ReturnType[T]{Data: nil, Error: fmt.Errorf("List is empty, nothing to delete")}
	}
	// if the value is on head
	if reflect.DeepEqual(l.head.data, value) {
		l.head = l.head.next
		// l.head.next = nil // it is not necessary to point the current head to nil
		l.length--
		return ReturnType[T]{Data: l.head.data, Error: nil}
	}

	toDelete := l.head // start on head.
	for !reflect.DeepEqual(toDelete.next.data, value) {
		// what if it does not exists.
		if toDelete.next.next == nil {
			// value does not exists in
			return ReturnType[T]{Data: nil, Error: fmt.Errorf("Cannot delete Node since value does not exist in this SLL")}
		}
		toDelete = toDelete.next
	}

	toDelete.next = toDelete.next.next
	l.length--
	return ReturnType[T]{Data: fmt.Sprintf("Deleted Node with %v", value), Error: nil}
}

func main() {
	node1 := Node[int]{data: 200}
	node2 := Node[int]{data: 20}
	node3 := Node[int]{data: 2}
	node4 := Node[int]{data: 1}
	newLinkedList := SinglyLinkedList[int]{}
	newLinkedList.prepend(node1)
	newLinkedList.prepend(node2)
	newLinkedList.prepend(node3)
	newLinkedList.prepend(node4)
	newLinkedList.printAllSLLData()
}
