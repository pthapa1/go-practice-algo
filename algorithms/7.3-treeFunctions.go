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

func DeleteItemDepthFirst(root *utils.Tree, data int) bool {
	// Find item first, if it does not exist then return false.
	// if the item exist, delete the node. It's parent's left or right side := nil
	// what if the root is the item we are trying to delete?
	// what do you replace the node with?
	return false
}