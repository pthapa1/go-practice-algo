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

// Read and print out all the items in the DLL.
func (d *DoblyLinkedList[T]) PrintAllDllData() []T {

	toPrint := d.Head
	dllValueList := []T{}

	for toPrint != nil {
		data := toPrint.Data
		fmt.Printf("%v =>", data)
		dllValueList = append(dllValueList, data)
		toPrint = toPrint.Next
	}
	fmt.Println()
	fmt.Println(dllValueList)
	return dllValueList
}

// Delete a node:

func (d *DoblyLinkedList[T]) DeleteDLLNode(value T) ReturnType[T] {
	// empty length
	if d.Length == 0 {
		return ReturnType[T]{Data: nil, Error: fmt.Errorf("list is empty. nothing to delete")}
	}

	if reflect.DeepEqual(d.Head.Data, value) {
		d.Head = d.Head.Next
		d.Length--
		return ReturnType[T]{Data: d.Head.Data, Error: nil}
	}

	toDelete := d.Head
	for !reflect.DeepEqual(toDelete.Next.Data, value) {
		// if the value does not exists
		if toDelete.Next.Next == nil {
			return ReturnType[T]{Data: nil, Error: fmt.Errorf("cannot perform delete operation since the %v does not exists on node", value)}
		}
		toDelete = toDelete.Next
	}
	newPrevious := toDelete
	toDelete.Next = toDelete.Next.Next
	toDelete.Next.Next.Previous = newPrevious
	d.Length--
	return ReturnType[T]{Data: fmt.Sprintf("Deleted Node with Value %v: ", value)}
}
