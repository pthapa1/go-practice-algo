package algo

type BinarySearchNode struct {
	Data  int
	Left  *BinarySearchNode
	Right *BinarySearchNode
}

// insert node to the tree.
func (n *BinarySearchNode) InsertNodeBST(data int) *BinarySearchNode {
	// the key should not be in the node.
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

func BinarySearchTree() {
	listOfInt := []int{100, 60, 59, 19, 190}
	var tree *BinarySearchNode

	for _, item := range listOfInt {
		tree = tree.InsertNodeBST(item)
	}
}
