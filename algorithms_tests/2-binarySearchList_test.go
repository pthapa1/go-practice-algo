package algo

import (
	"github.com/pthapa1/go-practice-algo/algorithms"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	testCases := []struct {
		list     []int
		value    int
		expected bool
	}{
		{[]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 69, true},
		{[]int{0, 3, 7, 8, 9, 23, 34}, 91, false},
		{[]int{0, 3, 7, 8, 9, 23, 34}, 9, true},
	}
	for _, tc := range testCases {
		result := algo.BinarySearch(tc.list, tc.value)
		if result != tc.expected {
			t.Errorf("Binary Search (%v, %v) expected: %v, but got %v", tc.list, tc.value, tc.expected, result)
		}

	}

}
