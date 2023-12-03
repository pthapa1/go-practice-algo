package main

import (
	"fmt"
)

type Tree struct {
	Data  int
	Right *Tree
	Left  *Tree
}

func recurseTheTree(currentNode *Tree, path []int) []int {
	if currentNode == nil {
		return path
	}

	path = recurseTheTree(currentNode.Left, path)
	path = recurseTheTree(currentNode.Right, path)
	path = append(path, currentNode.Data)

	return path
}

func searchTree(head *Tree) []int {
	newPath := []int{}
	result := recurseTheTree(head, newPath)
	return result
}

func main() {
	tree1 := &Tree{
		Data: 20,
		Right: &Tree{
			Data: 50,
			Right: &Tree{
				Data:  100,
				Right: nil,
				Left:  nil,
			},
			Left: &Tree{
				Data: 30,
				Right: &Tree{
					Data:  45,
					Right: nil,
					Left:  nil,
				},
				Left: &Tree{
					Data:  29,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &Tree{
			Data: 10,
			Right: &Tree{
				Data:  15,
				Right: nil,
				Left:  nil,
			},
			Left: &Tree{
				Data: 5,
				Right: &Tree{
					Data:  7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	}

	result := searchTree(tree1)
	fmt.Println(result)
}
