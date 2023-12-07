package algo

import "github.com/pthapa1/go-practice-algo/utils"

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

	// search if the item item exist on the left side of the node.
	answerFromTheLeft := FindItemDepthFirst(root.Left, data)

	if answerFromTheLeft {
		return true
	}

	if !answerFromTheLeft {
		answerFromTheRight := FindItemDepthFirst(root.Right, data)
		if answerFromTheRight {
			return true
		}
	}

	// if the item exist on the right side of the node.

	return false
}
