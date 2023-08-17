package algo

import (
	"fmt"
	"reflect"
)

type DLLNode[T any] struct {
	Data     T
	Next     *DLLNode[T]
	Previous *DLLNode[T]
}

type DoblyLinkedList[T any] struct {
	Head   *DLLNode[T]
	Length int
}

// Add item to the head of the linked list. Constant Time O(1)
func (d *DoblyLinkedList[T]) Prepend(n DLLNode[T]) {
	oldHead := d.Head
	d.Head = &n
	n.Next = oldHead
	n.Previous = nil
	if oldHead != nil {
		oldHead.Previous = d.Head
	}
	d.Length++
}

func (d *DoblyLinkedList[T]) GetLength() int {
	length := d.Length
	return length
}

// Read and print out all the items in the DLL.
func (d *DoblyLinkedList[T]) PrintAllDllData() []T {

	toPrint := d.Head
	dllValueList := []T{}

	for toPrint != nil {
		data := toPrint.Data
		fmt.Printf("%v <=> ", data)
		dllValueList = append(dllValueList, data)
		toPrint = toPrint.Next
	}
	fmt.Println()
	fmt.Println(dllValueList)
	return dllValueList
}

// Delete a node given the node value:
func (d *DoblyLinkedList[T]) DeleteDLLNode(value T) ReturnType[T] {
	// nothing to delete if the dll is empty.
	if d.Length == 0 {
		return ReturnType[T]{Data: nil, Error: fmt.Errorf("list is empty. nothing to delete")}
	}
	// if the value we are deleting is in the head.
	if reflect.DeepEqual(d.Head.Data, value) {
		d.Head = d.Head.Next
		if d.Head != nil { // Check if the list is now empty
			d.Head.Previous = nil
		}
		d.Length--
		return ReturnType[T]{Data: d.Head.Data, Error: nil}
	}

	toDelete := d.Head
	for toDelete.Next != nil && !reflect.DeepEqual(toDelete.Next.Data, value) {
		// move to the next node
		toDelete = toDelete.Next
	}

	// if reached the end and didn't find the node
	if toDelete.Next == nil {
		return ReturnType[T]{Data: nil, Error: fmt.Errorf("cannot perform delete operation since the %v does not exist in the list", value)}
	}

	// now, toDelete.Next is the node we want to delete
	newNext := toDelete.Next.Next
	if newNext != nil {
		newNext.Previous = toDelete
	}
	toDelete.Next = newNext
	d.Length--
	return ReturnType[T]{Data: d.Head.Data, Error: nil}
}
