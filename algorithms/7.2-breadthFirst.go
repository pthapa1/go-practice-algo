package algo

import (
	"github.com/pthapa1/go-practice-algo/utils"
)

func BreadthFirstTreeTraversal(root *utils.Tree) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	// create a slice that takes in nodes.
	queue := []*utils.Tree{root}

	for len(queue) > 0 {
		node := queue[0]  // copy the first node
		queue = queue[1:] // remove the first node
		result = append(result, node.Data)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

/*
there is no recursion in breadth first search
if you recurse in a tree of nodes, you create a stack
you are going in depth first.
So, here we create a queue slice and then we push and pop like we do
with a queue.
*/
