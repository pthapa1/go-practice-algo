package algo

import (
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
	"github.com/pthapa1/go-practice-algo/utils"
)

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
