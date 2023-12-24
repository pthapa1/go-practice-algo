package algo

import (
	"fmt"
	"reflect"
)

type ReturnType[T any] struct {
	Data  interface{}
	Error error
}

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	length int
}

// Add new node on the head of linked list. Contant Time O(1)
func (l *LinkedList[T]) Prepend(n Node[T]) {
	oldHead := l.head // save current head as old head
	l.head = &n       // make n our new head.
	n.Next = oldHead  // and point it to our old head
	l.length++        // increase length for prepend
}

// Print all data associated with Singly Linked List
func (l *LinkedList[T]) PrintAllSLLData() []T {
	toPrint := l.head
	nodeValueInlist := []T{}
	for toPrint != nil {
		data := toPrint.Data
		nodeValueInlist = append(nodeValueInlist, data)
		toPrint = toPrint.Next
	}
	fmt.Println()
	// fmt.Println(nodeValueInlist)
	return nodeValueInlist
}

// Delete a node given it's value.
func (l *LinkedList[T]) DeleteNode(value T) ReturnType[T] {
	if l.length == 0 {
		return ReturnType[T]{Data: nil, Error: fmt.Errorf("list is empty, nothing to delete")}
	}
	// if the value is on head
	if reflect.DeepEqual(l.head.Data, value) {
		l.head = l.head.Next
		// l.head.next = nil // it is not necessary to point the current head to nil
		l.length--
		return ReturnType[T]{Data: l.head.Data, Error: nil}
	}

	toDelete := l.head // start on head.
	for !reflect.DeepEqual(toDelete.Next.Data, value) {
		// what if it does not exists.
		if toDelete.Next.Next == nil {
			// value does not exists in
			return ReturnType[T]{
				Data:  nil,
				Error: fmt.Errorf("cannot delete Node since value does not exist in this SLL"),
			}
		}
		toDelete = toDelete.Next
	}

	toDelete.Next = toDelete.Next.Next
	l.length--
	return ReturnType[T]{Data: fmt.Sprintf("Deleted node with value %v", value)}
}
