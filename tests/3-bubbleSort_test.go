package algo

import (
	"github.com/pthapa1/go-practice-algo/algorithms"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		unSorted []int
		sorted   []int
	}{
		{[]int{9, 3, 7, 4, 69, 420, 42}, []int{3, 4, 7, 9, 42, 69, 420}},
		{[]int{9, 3, 5, 7, 8, 11}, []int{3, 5, 7, 8, 9, 11}},
	}
	for _, tc := range testCases {
		algo.BubbleSortWithPointer(tc.unSorted)
		if !reflect.DeepEqual(tc.sorted, tc.unSorted) {
			t.Errorf("Bubble Sort with Pointer: Expected %v but got %v", tc.sorted, tc.unSorted)
		}
	}
}
