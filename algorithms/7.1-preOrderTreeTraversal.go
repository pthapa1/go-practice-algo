// general tree traversal
// assuming that tree exists.

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

	path = recurseTheTree(currentNode.Left, path)
	path = recurseTheTree(currentNode.Right, path)

	return path
}

func PreOrderSearch(head *Tree) []int {
	newPath := []int{}
	return recurseTheTree(head, newPath)
}
