package algo

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
	d.Length++
}

// Add item to the middle or end of the list. O(N)
func (d *DoblyLinkedList[T]) PrependToMiddle(n DLLNode[T]) {

}
