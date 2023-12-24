package utils

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

// returns a random tree
func RandomTree() *Tree {
	tree1 := &Tree{
		Data: 20,
		Right: &Tree{
			Data: 50,
			Right: &Tree{
				Data:  100,
				Right: nil,
				Left:  nil,
			},
			Left: &Tree{
				Data: 30,
				Right: &Tree{
					Data:  45,
					Right: nil,
					Left:  nil,
				},
				Left: &Tree{
					Data:  29,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &Tree{
			Data: 10,
			Right: &Tree{
				Data:  15,
				Right: nil,
				Left:  nil,
			},
			Left: &Tree{
				Data: 5,
				Right: &Tree{
					Data:  7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	}
	return tree1
}

func GenerateSliceOfNodes(root *Tree) []int {
	values := []int{}
	if root == nil {
		return values
	}
	values = append(values, root.Data)

	values = append(values, GenerateSliceOfNodes(root.Left)...)
	values = append(values, GenerateSliceOfNodes(root.Right)...)

	return values
}
