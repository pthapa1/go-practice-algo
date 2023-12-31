package algo

import (
	"reflect"
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func TestHep(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput []int
	}{
		{[]int{10, 20, 30, 40, 16, 27}, []int{40, 30, 27, 10, 16, 20}},
		{[]int{1, 5, 7, 8, 9, 14, 16, 30, 34, 50}, []int{50, 34, 14, 16, 30, 5, 9, 1, 8, 7}},
	}
	result := &algo.MaxHeap{}
	for _, tc := range testCases {
		for _, v := range tc.input {
			result.Insert(v)
		}
		if reflect.DeepEqual(tc.expectedOutput, result) {
			t.Errorf(
				"Error on Inserting element on a heap. Expected: %v \n Got: %v",
				tc.expectedOutput,
				result,
			)
		}
	}
}
