package main

import (
	"fmt"
)

type Treemain struct {
	Data  int
	Right *Treemain
	Left  *Treemain
}

func FindItemDepthFirst(root *Treemain, data int) bool {
	// assuming the tree is not binary search tree

	// safety first
	if root == nil {
		return false
	}

	// if the node has the item we are looking for.
	if data == root.Data {
		return true
	}

	// search if the item item exist on the left subtree
	answerFromTheLeft := FindItemDepthFirst(root.Left, data)

	if answerFromTheLeft {
		return true
	}

	return FindItemDepthFirst(root.Right, data)
}

func main() {
	tree1 := &Treemain{
		Data: 20,
		Right: &Treemain{
			Data: 50,
			Right: &Treemain{
				Data:  100,
				Right: nil,
				Left:  nil,
			},
			Left: &Treemain{
				Data: 30,
				Right: &Treemain{
					Data:  45,
					Right: nil,
					Left:  nil,
				},
				Left: &Treemain{
					Data:  29,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &Treemain{
			Data: 10,
			Right: &Treemain{
				Data:  15,
				Right: nil,
				Left:  nil,
			},
			Left: &Treemain{
				Data: 5,
				Right: &Treemain{
					Data:  7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	}

	result := FindItemDepthFirst(tree1, 33)

	fmt.Println("result", result)
}
