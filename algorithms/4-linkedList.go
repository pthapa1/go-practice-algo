package algo

import "fmt"

// create a node first.
type node struct {
	data int
	next *node
}

// create a linkedList to connect node. You need head.

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	newSpace := l.head     // put your current head in the new space.
	l.head = n             // now the head space is empty, put your new node as the head.
	l.head.next = newSpace // point your new head to the new space
	l.length++             // increase the length because we added one item. Javascript Array or Slice in Go does this for you but you have to do it yourself in linked list
}

func (l linkedList) printLinkedListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data) // print the data
		toPrint = toPrint.next          // go to the next one.
		l.length--                      // decrease the length so that for loop (while loop in Js or Python) can actually end.
	}
	fmt.Println()
}

// Delete value from the node of a given linkedlist
func (l *linkedList) deleteNodeValue(value int) {
	if l.length == 0 {
		return // return if there is no linked list
	}
	if l.head.data == value {
		// if we are trying to delete the header, we need to update linked list with a new header
		l.head = l.head.next
		l.length--
		return
	}
	toDelete := l.head                // start at thead.
	for toDelete.next.data != value { // assign valueToDelete new next data until, we reach the value
		if toDelete.next.next == nil {
			return // return if the value does not exists in linked list
		}
		toDelete = toDelete.next // once the value is found, assign the value to delete
	}
	toDelete.next = toDelete.next.next // once the value is found, set the pointer to next node
	l.length--                         // decrease the length once done.

}

func LinkedList() {
	node1 := &node{data: 200}
	node2 := &node{data: 20}
	node3 := &node{data: 2}
	node4 := &node{data: 1}
	newLinkedList := linkedList{}
	newLinkedList.prepend(node1)
	newLinkedList.prepend(node2)
	newLinkedList.prepend(node3)
	newLinkedList.prepend(node4)
	fmt.Println(newLinkedList)          // Expected Result {memoryAddress numberOfitems}. {0x140000102b0 3} in my case
	newLinkedList.printLinkedListData() // Expected Result: 1, 2, 20, 200
	newLinkedList.deleteNodeValue(1)    // delete function
	newLinkedList.printLinkedListData() // Expected Result:1, 2, 20

}
