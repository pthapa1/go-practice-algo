package problems

import (
	"reflect"
	"sort"
	"testing"

	"github.com/pthapa1/go-practice-algo/problems"
)

// Helper function to compare two sets of subsets
func containsAllSubsets(actual, expected [][]int) bool {
	if len(actual) != len(expected) {
		return false
	}

	// Sort the inner slices
	for _, subset := range actual {
		sort.Ints(subset)
	}
	for _, subset := range expected {
		sort.Ints(subset)
	}

	// Sort the outer slices
	sort.Slice(actual, func(i, j int) bool {
		return less(actual[i], actual[j])
	})
	sort.Slice(expected, func(i, j int) bool {
		return less(expected[i], expected[j])
	})

	// Compare each subset
	for i := range actual {
		if !reflect.DeepEqual(actual[i], expected[i]) {
			return false
		}
	}

	return true
}

// Helper function to determine order of slices for sorting
func less(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func TestSubset(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{},
			expected: [][]int{{}},
		},
		{
			input:    []int{1},
			expected: [][]int{{}, {1}},
		},
		{
			input:    []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			input: []int{1, 2, 3},
			expected: [][]int{
				{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3},
			},
		},
		{
			input: []int{1, 2, 3, 4},
			expected: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
				{4},
				{1, 4},
				{2, 4},
				{1, 2, 4},
				{3, 4},
				{1, 3, 4},
				{2, 3, 4},
				{1, 2, 3, 4},
			},
		},
		{
			input:    []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			input: []int{-1, 0, 1},
			expected: [][]int{
				{}, {-1}, {0}, {-1, 0}, {1}, {-1, 1}, {0, 1}, {-1, 0, 1},
			},
		},
		{
			input: []int{5, 10, 15},
			expected: [][]int{
				{}, {5}, {10}, {5, 10}, {15}, {5, 15}, {10, 15}, {5, 10, 15},
			},
		},
	}

	for _, test := range tests {
		t.Run("Testing Subset function", func(t *testing.T) {
			result := problems.Subset(test.input)
			if !containsAllSubsets(result, test.expected) {
				t.Errorf(
					"For input %v, expected subsets %v, but got %v",
					test.input,
					test.expected,
					result,
				)
			}
		})
	}
}
