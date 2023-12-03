package main

import (
	"fmt"
)

type Treemain struct {
	Data  int
	Right *Treemain
	Left  *Treemain
}

func recurseTheTree(currentNode *Treemain, path []int) []int {
	if currentNode == nil {
		return path
	}

	path = recurseTheTree(currentNode.Left, path)
	path = recurseTheTree(currentNode.Right, path)
	path = append(path, currentNode.Data)

	return path
}

func searchTree(head *Treemain) []int {
	newPath := []int{}
	result := recurseTheTree(head, newPath)
	return result
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
	testSlice := []int{2, 4, 9, 19}

	fmt.Println(testSlice[1:], "See the magic", len(testSlice), len(testSlice[1:]))

	result := searchTree(tree1)
	fmt.Println(result)
}
