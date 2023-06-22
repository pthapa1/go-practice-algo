package algo

import "fmt"

type eachStackNode[T any] struct {
	data T
	next *eachStackNode[T]
}

type Stack[T any] struct {
	length int
	head   *eachStackNode[T]
	tail   *eachStackNode[T]
}

func (stack *Stack[T]) peek() T {
	return stack.tail.data
}

func (stack *Stack[T]) push(value T) {
	newNode := &eachStackNode[T]{data: value}
	stack.length++

	if stack.length == 0 {
		stack.head = newNode
		stack.tail = newNode
	}

	newNode.next = stack.tail
	stack.tail = newNode

}

func (stack *Stack[T]) pop() T {
	if stack.tail == nil && stack.head == nil {
		var sampleValue T
		return sampleValue
	}
	removedValue := stack.tail.data
	stack.tail = stack.tail.next
	stack.tail.next = nil
	return removedValue
}

func ExecuteStack() {

	newStack := &Stack[int]{}

	newStack.push(98)
	newStack.push(88)
	newStack.push(78)
	newStack.push(68)
	tailValue := newStack.peek()
	fmt.Println("Tail Value: ", tailValue) // Expected Value is the last one we added. 68 in this case
	removedValue := newStack.pop()
	fmt.Println("Removed Value: ", removedValue) // Expected Value is 68 as it is the tail
}
