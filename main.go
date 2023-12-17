package main

import (
	"fmt"
)

type Tree struct {
	Data  int
	Right *Tree
	Left  *Tree
}

// Helper function to find the minimum value node in the subtree
// rooted with the provided node
func minValueNode(node *Tree) *Tree {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Deletes a node with the given data and returns the new root of the subtree
// sorted binary search tree
func DeleteItemDepthFirst(root *Tree, dataToDelete int) *Tree {
	if root == nil {
		return root
	}

	// If the data to be deleted is smaller than the root's data,
	// then it lies in the left subtree
	if dataToDelete < root.Data {
		root.Left = DeleteItemDepthFirst(root.Left, dataToDelete)
	} else if dataToDelete > root.Data { // If the data to be deleted is greater than the root's data, then it lies in the right subtree
		root.Right = DeleteItemDepthFirst(root.Right, dataToDelete)
	} else {
		// Node with only one child or no child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Node with two children: Get the inorder successor (smallest
		// in the right subtree)
		temp := minValueNode(root.Right)

		// Copy the inorder successor's content to this node
		root.Data = temp.Data

		// Delete the inorder successor
		root.Right = DeleteItemDepthFirst(root.Right, temp.Data)
	}
	return root
}

func printTree(root *Tree, level int) {
	if root != nil {
		printTree(root.Right, level+1)
		for i := 0; i < level; i++ {
			fmt.Print("   ")
		}
		fmt.Println(root.Data)
		printTree(root.Left, level+1)
	}
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
	fmt.Println(tree1)

	root := &Tree{Data: 20}
	root.Left = &Tree{Data: 10}
	root.Left.Left = &Tree{Data: 5}
	root.Left.Right = &Tree{Data: 15}

	fmt.Println("Original tree:")
	printTree(root, 0)

	root = DeleteItemDepthFirst(root, 20) // Delete root node

	fmt.Println("Tree after deleting 20:")
	printTree(root, 0)
}
