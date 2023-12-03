package algo

import (
	"reflect"
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func RandomTree1() *algo.Tree {
	tree1 := &algo.Tree{
		Data: 20,
		Right: &algo.Tree{
			Data: 50,
			Right: &algo.Tree{
				Data:  100,
				Right: nil,
				Left:  nil,
			},
			Left: &algo.Tree{
				Data: 30,
				Right: &algo.Tree{
					Data:  45,
					Right: nil,
					Left:  nil,
				},
				Left: &algo.Tree{
					Data:  29,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &algo.Tree{
			Data: 10,
			Right: &algo.Tree{
				Data:  15,
				Right: nil,
				Left:  nil,
			},
			Left: &algo.Tree{
				Data: 5,
				Right: &algo.Tree{
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

func RandomTree2() *algo.Tree {
	// Tree 2
	tree2 := &algo.Tree{
		Data: 20,
		Right: &algo.Tree{
			Data:  50,
			Right: nil,
			Left: &algo.Tree{
				Data: 30,
				Right: &algo.Tree{
					Data: 45,
					Right: &algo.Tree{
						Data:  49,
						Right: nil,
						Left:  nil,
					},
					Left: nil,
				},
				Left: &algo.Tree{
					Data:  29,
					Right: nil,
					Left: &algo.Tree{
						Data:  21,
						Right: nil,
						Left:  nil,
					},
				},
			},
		},
		Left: &algo.Tree{
			Data: 10,
			Right: &algo.Tree{
				Data:  15,
				Right: nil,
				Left:  nil,
			},
			Left: &algo.Tree{
				Data: 5,
				Right: &algo.Tree{
					Data:  7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	}
	return tree2
}

func TestPreordertreeTraversal(t *testing.T) {
	// the path returned by algo.preOrder should deep equal
	// to the list below.
	preOrderPath := []int{
		20,
		10,
		5,
		7,
		15,
		50,
		30,
		29,
		45,
		100,
	}

	resultPreOrderPath := algo.PreOrderSearch(RandomTree1())

	if !reflect.DeepEqual(preOrderPath, resultPreOrderPath) {
		t.Errorf(
			"Boy, you have an error on preorder search. Expected: %v but got: %v",
			preOrderPath,
			resultPreOrderPath,
		)
	}
	// in order tree traversal tests
	inOrderPath := []int{
		5, 7, 10, 15, 20, 29, 30, 45, 50, 100,
	}
	inOrderResultPath := algo.InOrderSearch(RandomTree1())

	if !reflect.DeepEqual(inOrderPath, inOrderResultPath) {
		t.Errorf(
			"Boy, you have an error on inorder search. Expected: %v but got: %v",
			inOrderPath,
			inOrderResultPath,
		)
	}

	// post order tree traversal tests
	postOrderPath := []int{
		7, 5, 15, 10, 29, 45, 30, 100, 50, 20,
	}

	postOrderResult := algo.PostOrderSearch(RandomTree1())
	if !reflect.DeepEqual(postOrderPath, postOrderResult) {
		t.Errorf(
			"Boy, you have an error on postOrderSearch. Expected %v bug got: %v",
			postOrderPath,
			postOrderResult,
		)
	}
}
