package problems

import (
	"testing"

	"github.com/pthapa1/go-practice-algo/problems"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		input    []int
		target   int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:   5,
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 5, 6, 7, 9},
			target:   8,
			expected: false,
		},
		{
			input:    []int{1, 2, 3, 5, 6, 7, 9},
			target:   11,
			expected: false,
		},
		{
			input:    []int{2, 2, 2, 2, 2, 2, 2},
			target:   11,
			expected: false,
		},
		{
			input:    []int{},
			target:   11,
			expected: false,
		},
	}

	for _, v := range testCases {
		result := problems.LoopBinarySearch(v.input, v.target)
		if result != v.expected {
			t.Errorf(
				"Error on Loop Binary Search.\n Expected %v \n Got %v",
				v.expected,
				result,
			)
		}
		resultRecurse := problems.RecursiveBinarySearch(v.input, v.target)
		if resultRecurse != v.expected {
			t.Errorf(
				"Error on Recursive Binary Search.\n Expected %v \n Got %v",
				v.expected,
				result,
			)
		}
	}
}
