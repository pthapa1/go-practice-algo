package algo

import (
	"github.com/pthapa1/go-practice-algo/algorithms"
	"testing"
)

func TestLinearSearchList(t *testing.T) {
	testCases := []struct {
		list     []int
		value    int
		expected bool
	}{
		{[]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 71, true},
		{[]int{3, 4, 7, 9, 29, 30, 41}, 30, true},
		{[]int{2, 5}, 14, false},
		{[]int{2, 5, 8, 9, 15}, 14, false},
	}

	for _, tc := range testCases {
		result := algo.LinearSearch(tc.list, tc.value)
		if result != tc.expected {
			t.Errorf("Linear Search (%v, %d) = %v, expected: %v", tc.list, tc.value, result, tc.expected)
		}
	}
}
