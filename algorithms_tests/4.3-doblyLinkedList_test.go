package algo

import (
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func TestDLLPrepend(t *testing.T) {
	list := algo.DoublyLinkedList[int]{}
	list.Prepend(algo.DLLNode[int]{Data: 1})

	if list.Length != 1 {
		t.Fatalf("expected length to be 1, got %d", list.Length)
	}

	if list.Head.Data != 1 {
		t.Fatalf("expected head data to be 1, got %d", list.Head.Data)
	}

	list.Prepend(algo.DLLNode[int]{Data: 2})
	if list.Head.Data != 2 {
		t.Fatalf("expected head data to be 2, got %d", list.Head.Data)
	}
	if list.Head.Next.Data != 1 {
		t.Fatalf("expected next node data to be 1, got %d", list.Head.Next.Data)
	}
	if list.Head.Next.Previous.Data != 2 {
		t.Fatalf("expected previous node data of next node to be 2, got %d", list.Head.Next.Previous.Data)
	}
}

func TestDeleteDLLNode(t *testing.T) {
	list := algo.DoublyLinkedList[int]{}
	list.Prepend(algo.DLLNode[int]{Data: 1})
	list.Prepend(algo.DLLNode[int]{Data: 2})
	list.Prepend(algo.DLLNode[int]{Data: 3})

	// delete middle node
	ret := list.DeleteDLLNode(2)
	if ret.Error != nil {
		t.Fatalf("unexpected error: %v", ret.Error)
	}
	if list.Length != 2 {
		t.Fatalf("expected length to be 2, got %d", list.Length)
	}
	if list.Head.Next.Data != 1 {
		t.Fatalf("expected next node data to be 1, got %d", list.Head.Next.Data)
	}

	// delete head
	ret = list.DeleteDLLNode(3)
	if ret.Error != nil {
		t.Fatalf("unexpected error: %v", ret.Error)
	}
	if list.Length != 1 {
		t.Fatalf("expected length to be 1, got %d", list.Length)
	}
	if list.Head.Data != 1 {
		t.Fatalf("expected head data to be 1, got %d", list.Head.Data)
	}

	// delete non-existing value
	ret = list.DeleteDLLNode(4)
	if ret.Error == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestGetlength(t *testing.T) {
	list := algo.DoublyLinkedList[int]{}
	length := list.GetLength()
	if length != 0 {
		t.Fatalf("expected length to be 0 but got %v", length)
	}
	list.Prepend(algo.DLLNode[int]{Data: 1})
	list.Prepend(algo.DLLNode[int]{Data: 2})
	list.Prepend(algo.DLLNode[int]{Data: 3})

	newLength := list.GetLength()
	if newLength != 3 {
		t.Fatalf("expected length to be 3 but got %v", newLength)
	}
}
