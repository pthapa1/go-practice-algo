package algo

import (
	"github.com/pthapa1/go-practice-algo/utils"
)

func BreadthFirstTreeTraversal(root *utils.Tree) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	// create an slice that takes in nodes.
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
