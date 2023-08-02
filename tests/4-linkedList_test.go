package algo

import (
	"reflect"
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func TestLinkedList(t *testing.T) {

	node1 := algo.Node[int]{Data: 2000}
	node2 := algo.Node[int]{Data: 200}
	node3 := algo.Node[int]{Data: 20}
	node4 := algo.Node[int]{Data: 2}
	node5 := algo.Node[int]{Data: 33}

	singleLinkedList := algo.SinglyLinkedList[int]{}
	singleLinkedList.Prepend(node1)
	singleLinkedList.Prepend(node2)
	singleLinkedList.Prepend(node3)
	singleLinkedList.Prepend(node4)
	singleLinkedList.Prepend(node5)

	expectedOutput := []int{33, 2, 20, 200, 2000}
	output := singleLinkedList.PrintAllSLLData()

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Expected %v, Got %v", expectedOutput, output)
	}
	// testing delete
	result := singleLinkedList.DeleteNode(2)
	if result.Data != "Deleted node with value 2" {
		t.Errorf(" Expected: Deleted node with value 2, Got: %v, Error: %v", result.Data, result.Error)
	}

	// non existant node.
	nonExistNodeResult := singleLinkedList.DeleteNode(10)
	if nonExistNodeResult.Data != nil || nonExistNodeResult.Error == nil {
		t.Errorf("Expected: {Data: %v, Error: %v} but Got:  {Data: %v, Error: %v}", nil,
			"Cannot delete Node since value does not exist in this SLL", nonExistNodeResult.Data, nonExistNodeResult.Error)
	}
	//node with length 0
	sll := algo.SinglyLinkedList[int]{}

	emptySLL := sll.DeleteNode(1)
	if emptySLL.Data != nil || emptySLL.Error == nil {
		t.Errorf("Expected Linked List to be empty but it is not.")
	}

}
