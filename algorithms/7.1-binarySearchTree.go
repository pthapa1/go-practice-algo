package algo

type BinarySearchNode struct {
	Data  int
	Left  *BinarySearchNode
	Right *BinarySearchNode
}

// insert node to the tree.
func (n *BinarySearchNode) InsertNodeBST(data int) *BinarySearchNode {
	// if the data we are adding is smaller than the node,
	// 	we add it to the left if greater, we move it to right

	// if the head is nil add data to the head
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
			n.Left.InsertNodeBST(data)
		}

	} else if data > n.Data {
		// same as above but on the other side.
		if n.Right == nil {
			n.Right = &BinarySearchNode{Data: data}
		} else {
			n.Right.InsertNodeBST(data)
		}

	}
	return n

}

// takes address of BST, and returns if the items exists or not
func SearchBinaryNode(tree *BinarySearchNode, item int) bool {
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
		return SearchBinaryNode(tree.Left, item)

	} else {
		// go to the right
		if tree.Right == nil {
			return false
		}
		return SearchBinaryNode(tree.Right, item)
	}

}

// create BST with a list of int
func (n *BinarySearchNode) CreateBinarySearchTree(listOfInt []int) *BinarySearchNode {
	var tree *BinarySearchNode

	for _, item := range listOfInt {
		tree = tree.InsertNodeBST(item)
	}
	return tree
}
