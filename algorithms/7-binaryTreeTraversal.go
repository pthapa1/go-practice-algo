// general tree traversals
// perOerder
// inOrder &
// postOrder
// assuming that tree exists.

package algo

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

func preOrderRecursion(currentNode *Tree, path []int) []int {
	if currentNode == nil {
		return path
	}

	path = append(path, currentNode.Data)

	path = preOrderRecursion(currentNode.Left, path)
	path = preOrderRecursion(currentNode.Right, path)

	return path
}

func PreOrderSearch(head *Tree) []int {
	newPath := []int{}
	return preOrderRecursion(head, newPath)
}

// in order tree traversal
func inOrderRecursion(currentNode *Tree, path []int) []int {
	if currentNode == nil {
		return path
	}

	path = inOrderRecursion(currentNode.Left, path)
	path = append(path, currentNode.Data)
	path = inOrderRecursion(currentNode.Right, path)

	return path
}

func InOrderSearch(head *Tree) []int {
	newPath := []int{}
	result := inOrderRecursion(head, newPath)
	return result
}

// postOrder tree traversal

func postOrderRecursion(currentNode *Tree, path []int) []int {
	if currentNode == nil {
		return path
	}
	path = postOrderRecursion(currentNode.Left, path)
	path = postOrderRecursion(currentNode.Right, path)
	path = append(path, currentNode.Data)

	return path
}

func PostOrderSearch(head *Tree) []int {
	newPath := []int{}
	return postOrderRecursion(head, newPath)
}
