package algo

import (
	"github.com/pthapa1/go-practice-algo/algorithms"
	"reflect"
	"testing"
)

func TestExecuteQuickSort(t *testing.T) {
	// unsortedArray := [6]int{8, 7, 0, 2, 13, 78}
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{8, 7, 0, 2, 13, 78}, []int{0, 2, 7, 8, 13, 78}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {

		inputCopy := make([]int, len(tc.input))
		copy(inputCopy, tc.input)

		algo.ExecuteQuickSort(tc.input)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			t.Errorf("ExecuteQuickSort(%v) resulted in %v, expected %v", inputCopy, tc.input, tc.expected)
		}
	}
}
