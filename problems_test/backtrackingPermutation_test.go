package problems

import (
	"reflect"
	"testing"

	"github.com/pthapa1/go-practice-algo/problems"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			input: []int{1, 2},
			expected: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			input: []int{1},
			expected: [][]int{
				{1},
			},
		},
		{
			input:    []int{},
			expected: [][]int{},
		},
	}

	for _, test := range tests {
		result := problems.Permute(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Permute(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
