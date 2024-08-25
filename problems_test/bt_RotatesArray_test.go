package problems

import (
	"testing"

	"github.com/pthapa1/go-practice-algo/problems"
)

func TestBsRotatedArray(t *testing.T) {
	// Define a struct for test cases
	type testCase struct {
		arr      []int
		expected int
	}

	// Create a slice of test cases
	tests := []testCase{
		{arr: []int{4, 5, 6, 7, 0, 1, 2}, expected: 0},
		{arr: []int{6, 7, 1, 2, 3, 4, 5}, expected: 1},
		{arr: []int{10}, expected: 10},
		{arr: []int{1, 2, 3, 4, 5, 6, 7}, expected: 1},
		{arr: []int{2, 1}, expected: 1},
		{arr: []int{1, 2}, expected: 1},
		{arr: []int{1, 1, 1, 1, 1, 1, 1}, expected: 1},
		{arr: []int{2, 2, 2, 0, 1, 2, 2}, expected: 0},
		{arr: []int{10, 11, 12, 13, 14, 5, 6, 7, 8, 9}, expected: 5},
		{arr: []int{1, 2, 3, 4, 5, 6, 7}, expected: 1},
		{arr: []int{2, 3, 4, 5, 6, 7, 1}, expected: 1},
		{arr: append([]int{2, 3, 4, 5, 6, 7}, 1), expected: 1},
		{arr: []int{30, 40, 50, 10, 20}, expected: 10},
		{arr: []int{2, 2, 2, 2, 2, 2, 2}, expected: 2},
	}

	// Loop over the test cases
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := problems.RotatedArrayBinarySearchWithRecursion(test.arr)
			if result != test.expected {
				t.Errorf("Expected %d, but got %d, for %v", test.expected, result, test.arr)
			}
		})
	}
}
