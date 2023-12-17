package algo

import (
	"github.com/pthapa1/go-practice-algo/utils"
)

/*
To find an item in a binary tree, in an efficient manner,
we should be aware of how the tree is structured.
* If the item you are trying to find is deep down from the root,
depth first search finds the item in an efficient manner. For example,
a character in game could choose an axe as a weapon, if axe then the type of blade,
if blade then the material, so on and so forth.
* If you are trying to find an item, closer to the root,
breadth first approach is efficient. For example, when you are trying to
find shortest path to the next node (GPS system, web crawlers).
*/

func FindItemDepthFirst(root *utils.Tree, data int) bool {
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

func FindItemBreadthFirst(root *utils.Tree, data int) bool {
	if root == nil {
		return false
	}
	if root.Data == data {
		return true
	}
	nodeContainer := []*utils.Tree{root}

	for len(nodeContainer) > 0 {
		node := nodeContainer[0]          // copy the first node
		nodeContainer = nodeContainer[1:] // remove the first node

		if node.Left != nil && node.Left.Data == data {
			return true
		}

		if node.Right != nil && node.Right.Data == data {
			return true
		}
		// if the childrens do not have items, queue more nodes
		if node.Left != nil {
			nodeContainer = append(nodeContainer, node.Left)
		}

		if node.Right != nil {
			nodeContainer = append(nodeContainer, node.Right)
		}
	}

	return false
}

func findMax(node *utils.Tree) *utils.Tree {
	current := node
	if current.Right != nil {
		current = current.Right
	}
	return current
}

func findMin(node *utils.Tree) *utils.Tree {
	current := node
	if current.Left != nil {
		current = current.Left
	}
	return current
}

func DeleteItemDepthFirst(root *utils.Tree, dataToDelete int) *utils.Tree {
	if root == nil {
		return root
	}

	// if the root is the data we are trying to delete and
	if root.Data == dataToDelete {
		// Node with no children
		if root.Left == nil && root.Right == nil {
			root = nil
		}
		// Node with only one child
		if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		}

		/* If node has two childeren, we need to pick which child should be new root. It could be either one.
		   Or it could be based on some criteria, for now lets make left child as root if available.
		*/
		if root.Left != nil {
			replacement := findMax(
				root.Left,
			) // find the maximum in the left subtree
			root.Data = replacement.Data // Replace root data with replacement data
			root.Left = DeleteItemDepthFirst(
				root.Left,
				replacement.Data,
			) // Recursively delete the root node
			return root
		} else {
			replacement := findMin(root.Right)                              // Find the minimum in the right subtree
			root.Data = replacement.Data                                    // Replace root data with replacement data
			root.Right = DeleteItemDepthFirst(root.Right, replacement.Data) // Recursively delete the replacement node
			return root
		}

	}

	// Recurse on left and right children
	root.Left = DeleteItemDepthFirst(root.Left, dataToDelete)
	root.Right = DeleteItemDepthFirst(root.Right, dataToDelete)

	return root
}
