package algo

import (
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

func TestDeleteItemsInTree() {
	// test that when deleting node or any other element in the random tree, the tree
	// there are no errors

	// what happens when you delete node
	// leaf and
	// two random nodes with one and two children each

	// create a function that generates the slice.
	// and compare those two slices.

	// find the difference between pre order recursion 7
	// and the binary tree utils
}
