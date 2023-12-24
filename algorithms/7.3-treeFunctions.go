// all functions assume that the tree is not sorted

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

// returns either the right most or the left most pointer of the node
func RightOrLeftMostNode(node *utils.Tree) *utils.Tree {
	if node == nil {
		return node
	}

	if node.Right != nil {
		return RightOrLeftMostNode(node.Right)
	} else if node.Left != nil {
		return RightOrLeftMostNode(node.Left)
	}
	return node
}

func DeleteDepthFirst(root *utils.Tree, dataToDelete int) *utils.Tree {
	if root == nil {
		return root
	}

	if root.Right != nil && root.Right.Data == dataToDelete {
		// leaf: node without any children
		if root.Right.Right == nil && root.Right.Left == nil {
			root.Right = nil
			return root
		} else {
			// node with children, has at least 1
			replacementNode := RightOrLeftMostNode(root.Right)
			root.Right.Data = replacementNode.Data
			root.Right = DeleteDepthFirst(root.Right, replacementNode.Data)
			return root
		}
	}

	if root.Left != nil && root.Left.Data == dataToDelete {
		// leaf: node without any children
		if root.Left.Right == nil && root.Left.Left == nil {
			root.Left = nil
			return root
		} else {
			replacementNode := RightOrLeftMostNode(root.Left)
			root.Left.Data = replacementNode.Data
			root.Left = DeleteDepthFirst(root.Left, replacementNode.Data)
			return root
		}
	}

	DeleteDepthFirst(root.Right, dataToDelete)
	DeleteDepthFirst(root.Left, dataToDelete)

	return root
}

// after writing tests for the funciton above, I realized that my function does not
// delete the root node. So, I wrote another function to delete the root node.
// I can probably combine these two function but I am too lazy to do that
func DeleteDepthFirstRootNode(root *utils.Tree, rootVal int) bool {
	if root == nil {
		return false
	}

	if root.Data != rootVal {
		return false
	}

	var replacementNode *utils.Tree
	var fromLeftSubtree bool

	// Check if the left subtree exists
	if root.Left != nil {
		replacementNode = RightOrLeftMostNode(root.Left)
		fromLeftSubtree = true
	} else if root.Right != nil {
		replacementNode = RightOrLeftMostNode(root.Right)
		fromLeftSubtree = false
	} else {
		// Root is a leaf node, nothing to replace with
		return false
	}

	// Replace root data with replacement node data
	root.Data = replacementNode.Data

	// Delete the replacement node
	if fromLeftSubtree {
		root.Left = deleteSpecificNode(root.Left, replacementNode.Data)
	} else {
		root.Right = deleteSpecificNode(root.Right, replacementNode.Data)
	}

	return true
}

// Function to delete a specific node by value
func deleteSpecificNode(node *utils.Tree, data int) *utils.Tree {
	if node == nil {
		return nil
	}

	if node.Data == data {
		if node.Left == nil && node.Right == nil {
			// Node is a leaf
			return nil
		} else if node.Left != nil {
			// Promote the left subtree
			return node.Left
		} else {
			// Promote the right subtree
			return node.Right
		}
	}

	node.Left = deleteSpecificNode(node.Left, data)
	node.Right = deleteSpecificNode(node.Right, data)
	return node
}
