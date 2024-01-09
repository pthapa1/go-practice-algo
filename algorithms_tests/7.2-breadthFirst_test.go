package algo

import (
	"reflect"
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
	"github.com/pthapa1/go-practice-algo/utils"
)

func TestBreadthFirstTreeTraversal(t *testing.T) {
	expectedPath := []int{20, 10, 50, 5, 15, 30, 100, 7, 29, 45}

	path := algo.BreadthFirstTreeTraversal(utils.RandomTree())
	if !reflect.DeepEqual(expectedPath, path) {
		t.Errorf(
			"Boy, you have an error on BreadthFirstTreeTraversal. Expected: %v but got %v",
			expectedPath,
			path,
		)
	}
}
