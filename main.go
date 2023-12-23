package main

import (
	"fmt"

	"github.com/pthapa1/go-practice-algo/utils"
)

type Tree struct {
	Data  int
	Right *Tree
	Left  *Tree
}

// returns either the right most or the left most pointer to the node
func rightOrLeftMostValue(node *Tree) *Tree {
	if node == nil {
		return node
	}

	if node.Right != nil {
		return rightOrLeftMostValue(node.Right)
	} else if node.Left != nil {
		return rightOrLeftMostValue(node.Left)
	}
	return node
}

func printTree(root *utils.Tree, level int) {
	if root != nil {
		printTree(root.Right, level+1)
		for i := 0; i < level; i++ {
			fmt.Print("   ")
		}
		fmt.Println(root.Data)
		printTree(root.Left, level+1)
	}
}

func deleteDepthFirst(root *Tree, dataToDelete int) *Tree {
	if root == nil {
		return root
	}

	if root.Right != nil && root.Right.Data == dataToDelete {
		// leaf: node without any children
		if root.Right.Right == nil && root.Right.Left == nil {
			root.Right = nil
			return root
		} else {
			// node with children
			replacementNode := rightOrLeftMostValue(root)
			root.Right.Data = replacementNode.Data
			root.Right = deleteDepthFirst(root.Right, replacementNode.Data)
			return root
		}
	}

	if root.Left != nil && root.Left.Data == dataToDelete {
		// leaf: node without any children
		if root.Left.Right == nil && root.Left.Left == nil {
			root.Left = nil
			return root
		} else {
			replacementNode := rightOrLeftMostValue(root)
			root.Left.Data = replacementNode.Data
			root.Left = deleteDepthFirst(root.Left, replacementNode.Data)
			return root
		}
	}

	deleteDepthFirst(root.Right, dataToDelete)
	deleteDepthFirst(root.Left, dataToDelete)

	return root
}

func main() {
	root := &Tree{Data: 20}
	root.Left = &Tree{Data: 10}
	root.Left.Left = &Tree{Data: 5}
	root.Left.Right = &Tree{Data: 15}
	root.Right = &Tree{Data: 11}
	root.Right.Left = &Tree{Data: 12}
	root.Right.Right = &Tree{Data: 13}

	// fmt.Println("Original tree:")
	// printTree(root, 0)

	// myLeaf := deleteDepthFirst(root, 11)
	fmt.Println("✍️")
	// printTree(myLeaf, 0)

	randomTree := utils.RandomTree()
	fmt.Println("This is a random tree")
	printTree(randomTree, 0)
	// check deleting random tree's node
}
