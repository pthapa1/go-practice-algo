package algo

import "fmt"

type nodeForQueue[T any] struct {
	data T
	next *nodeForQueue[T]
}

type queue[T any] struct {
	head   *nodeForQueue[T]
	tail   *nodeForQueue[T]
	length int
}

func (q *queue[T]) enqueue(value T) {
	newNode := &nodeForQueue[T]{data: value, next: nil}
	// if the queue is empty,
	if q.length == 0 {
		q.head = newNode
		q.tail = newNode
	}
	q.tail.next = newNode
	q.tail = newNode
	q.length++
}

func (q *queue[T]) dequeue() (T, error) {
	// when the queue is empty.
	if q.length == 0 {
		var nope T
		return nope, fmt.Errorf("Queue is Empty")
	}
	removedValue := q.head.data

	// only 1 item
	if q.length == 1 {
		q.head = nil
		q.tail = nil
	}

	q.head = q.head.next
	q.length--
	return removedValue, nil
}

func (q *queue[T]) peek() T {
	value := q.head.data
	return value
}

func (q queue[T]) printQueuedData() {
	valueToPrint := q.head

	for q.length != 0 {
		fmt.Println(valueToPrint.data)
		valueToPrint = valueToPrint.next
		q.length--
	}
}

func ExecuteQueue() {

	newQueue := queue[int]{}

	newQueue.enqueue(88)
	newQueue.enqueue(78)
	newQueue.enqueue(68)
	newQueue.enqueue(58)
	fmt.Println("Queue Memory Address:", newQueue)
	newQueue.printQueuedData()

	headValue := newQueue.peek()
	fmt.Println("Head Value: ", headValue)

	removedValue, error := newQueue.dequeue()
	fmt.Println("Removed Value:", removedValue)
	if error != nil {
		fmt.Println(error)
	}

	removedValue2, err2 := newQueue.dequeue()
	fmt.Println("Removed Value 2:", removedValue2)
	if err2 != nil {
		fmt.Println(err2)
	}

}
