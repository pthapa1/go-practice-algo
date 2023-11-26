// general tree traversal

package algo

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

func recurseTheTree(currentNode *Tree, path []int) []int {
	if currentNode == nil {
		return path
	}

	path = append(path, currentNode.Data)

	recurseTheTree(currentNode.Left, path)
	recurseTheTree(currentNode.Right, path)

	return path

}

// make a tree
// search the tree.
// write tests to prove that the it is working.
// insert node to the tree
// func (n *BinarySearchNode) InsertNodeInTree(data int) *Tree {
//  return recurseTheTree(, path []int)
// }
