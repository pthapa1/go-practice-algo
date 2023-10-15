package algo

import "fmt"

type BinarySearchNode struct {
	Data  int
	Left  *BinarySearchNode
	Right *BinarySearchNode
}

// insert node to the tree.
func (n *BinarySearchNode) insertNodeBST(data int) *BinarySearchNode {
	// if the data we are adding is smaller than the node,
	// 	we add it to the left if greater, we move it to right

	// if the head is nil
	if n == nil {
		n = &BinarySearchNode{Data: data}
	}

	// if data exists on the node already, exit.
	if data == n.Data {
		return n
	}

	if data < n.Data {
		// if the node is empty, put the item there.
		if n.Left == nil {
			n.Left = &BinarySearchNode{Data: data}
		} else {
			// if it's not empty,
			// compare the data again with the right and left & repeat
			n.Left.insertNodeBST(data)
		}

	} else if data > n.Data {
		// same as above but on the other side.

		if n.Right == nil {
			n.Right = &BinarySearchNode{Data: data}
		} else {
			n.Right.insertNodeBST(data)
		}

	}
	return n

}

func searchBinaryNode(tree *BinarySearchNode, item int) bool {
	// Assuming
	// the tree is balanced &
	// the tree is BST

	// if the tree/node itself is empty
	if tree == nil {
		return false
	}
	// if the item matches the key node
	if item == tree.Data {
		return true
	}
	// if the item is less than the key node, go left
	if item < tree.Data {
		if tree.Left == nil {
			return false
		}
		return searchBinaryNode(tree.Left, item)

	} else {
		// go to the right
		if tree.Right == nil {
			return false
		}
		return searchBinaryNode(tree.Right, item)
	}

}

func CreateBinarySearchTree() *BinarySearchNode {
	listOfInt := []int{100, 60, 59, 19, 190}
	var tree *BinarySearchNode

	for _, item := range listOfInt {
		tree = tree.insertNodeBST(item)
	}
	return tree
}

func SearchBinarySearchTree() bool {
	addressOfTree := CreateBinarySearchTree()
	trueOrFalse := searchBinaryNode(addressOfTree, 19)
	fmt.Println(trueOrFalse)
	return trueOrFalse
}
