package algo

import (
	"math/rand"
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func TestBinarySearchTree(t *testing.T) {
	// create a binary search tree with Insert Node from a list.
	numArr := [8]int{1, 3, 5, 8, 12, 13, 11, 10}

	var sampleBST *algo.BinarySearchNode
	treeAddress := sampleBST.CreateBinarySearchTree(numArr[:])
	randomIndex := rand.Intn(len(numArr))
	randomNumber := numArr[randomIndex]
	if !algo.SearchBinaryNode(treeAddress, randomNumber) {
		t.Errorf("Tree Does not contain number: %d", randomNumber)
	}

	// assert a duplicate value does not get added.
	duplicateValue := numArr[randomIndex]
	treeAddressDuplicate := sampleBST.InsertNodeBST(duplicateValue)
	if !algo.SearchBinaryNode(treeAddressDuplicate, duplicateValue) {
		t.Errorf("The tree does not handle duplicates correctly. Adding: %d", duplicateValue)
	}
}
