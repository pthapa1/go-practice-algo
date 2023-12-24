package algo

import (
	"slices"
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
	"github.com/pthapa1/go-practice-algo/utils"
)

// search both depthFirst and BreadthFirstSearch
func TestFindItmesInTree(t *testing.T) {
	tree := utils.RandomTree()
	includeditems := []int{20, 10, 50, 15, 7, 45}
	missingItems := []int{33, 44, 99}

	for _, val := range includeditems {
		if !algo.FindItemDepthFirst(tree, val) {
			t.Errorf("Expected %v to exist on the tree. DepthFirstSearch", val)
		}
		if !algo.FindItemBreadthFirst(tree, val) {
			t.Errorf(
				"Expected %v to exist on the tree. BreadthFirstSearch",
				val,
			)
		}
	}

	for _, val := range missingItems {
		if algo.FindItemDepthFirst(tree, val) {
			t.Errorf(
				"Expected %v to not exist on the tree. DepthFirstSearch",
				val,
			)
		}

		if algo.FindItemBreadthFirst(tree, val) {
			t.Errorf(
				"Expected %v to not exist on the tree. DreadthFirstSearch",
				val,
			)
		}
	}
}

func TestDeleteItemsInTree(t *testing.T) {
	randomTree := utils.RandomTree()

	nodeData := 20
	leaf1 := 100
	leaf2 := 7
	nodeWithChild := 10

	// deleting leaves

	deletedTree := algo.DeleteDepthFirst(randomTree, leaf1)
	deletedSlice := utils.GenerateSliceOfNodes(deletedTree)
	if slices.Contains(deletedSlice, leaf1) {
		t.Errorf("Expected %v to not exist on TestDeleteItemsInTree", leaf1)
	}

	deletedTree = algo.DeleteDepthFirst(deletedTree, leaf2)
	deletedSlice = utils.GenerateSliceOfNodes(deletedTree)
	if slices.Contains(deletedSlice, leaf2) {
		t.Errorf("Expected %v to not exist on TestDeleteItemsInTree", leaf2)
	}

	// deleting a node with children
	deletedTree = algo.DeleteDepthFirst(deletedTree, nodeWithChild)
	deletedSlice = utils.GenerateSliceOfNodes(deletedTree)
	if slices.Contains(deletedSlice, nodeWithChild) {
		t.Errorf("Tree should not contain node %v after deletion", nodeWithChild)
	}

	// replacement node should be removed
	replacementData := algo.RightOrLeftMostNode(deletedTree).Data
	counter := 0
	for _, v := range deletedSlice {
		if v == replacementData {
			counter++
		}
	}
	// assumes that the tree we are testing with has one distinct element
	// on each node. Works on this test.
	if counter > 1 {
		t.Errorf(
			"Replacement Node exist in the original position. Multiple occurances of: %v",
			replacementData,
		)
	}

	// deleting node

	deletionRootnode := algo.DeleteDepthFirstRootNode(deletedTree, nodeData)
	if !deletionRootnode {
		t.Errorf("Node data '%v' should be deleted but it's not", nodeData)
	}
}
