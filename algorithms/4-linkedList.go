package algo

// import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type SinglyLinkedList[T any] struct {
	head   *Node[T]
	length int
}

// Add new node on the head of linked list. Constant time O(1)
func (l *SinglyLinkedList[T]) prepend(n *Node[T]) {
	oldHead := l.head // save current head as old head
	l.head = n        // make n our new head.
	n.next = oldHead  // and point it to our old head
	l.length++        // increase length for prepend
}

// // create a node first.
// type node struct {
// 	data int
// 	next *node
// }

// // create a linkedList to connect node. You need head.
// type linkedList struct {
// 	head   *node
// 	length int
// }

// func (l *linkedList) prepend(n *node) {
// 	newSpace := l.head     // put your current head in the new space.
// 	l.head = n             // now the head space is empty, put your new node as the head.
// 	l.head.next = newSpace // point your new head to the new space
// 	l.length++             // increase the length because we added one item. Javascript Array or Slice in Go does this for you but you have to do it yourself in linked list
// }

// func (l SinglyLinkedList[T]) printLinkedListData() {
// 	toPrint := l.head
// 	for l.length != 0 {
// 		fmt.Printf("%d ", toPrint.data) // print the data
// 		toPrint = toPrint.next          // go to the next one.
// 		l.length--                      // decrease the length so that for loop (while loop in Js or Python) can actually end.
// 	}
// 	fmt.Println()
// }

// // Delete value from the node of a given linkedlist
// func (l *SinglyLinkedList[T]) deleteNodeValue(value T) {
// 	if l.length == 0 {
// 		return // return if there is no linked list
// 	}

// 	if l.head.data == value {
// 		// if we are trying to delete the header, we need to update linked list with a new header
// 		l.head = l.head.next
// 		l.length--
// 		return
// 	}
// 	toDelete := l.head                // start at thead.
// 	for toDelete.next.data != value { // assign valueToDelete new next data until, we reach the value
// 		if toDelete.next.next == nil {
// 			return // return if the value does not exists in linked list
// 		}
// 		toDelete = toDelete.next // keep moving forward
// 	}
// 	toDelete.next = toDelete.next.next // once the value is found, set the pointer to next node
// 	l.length--                         // decrease the length once done.

// }

// func LinkedList() {
// 	node1 := Node[int]{data: 200}
// 	node2 := Node[int]{data: 20}
// 	node3 := Node[int]{data: 2}
// 	node4 := Node[int]{data: 1}
// 	newLinkedList := SinglyLinkedList[int]{}
// 	newLinkedList.prepend(&node1)
// 	newLinkedList.prepend(&node2)
// 	newLinkedList.prepend(&node3)
// 	newLinkedList.prepend(&node4)
// 	fmt.Println(newLinkedList)          // Expected Result {memoryAddress numberOfitems}. {0x140000102b0 4} in my case
// 	newLinkedList.printLinkedListData() // Expected Result: 1, 2, 20, 200
// 	newLinkedList.deleteNodeValue(1)    // delete function
// 	newLinkedList.printLinkedListData() // Expected Result:1, 2, 20

// }
