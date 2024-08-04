package problems

import (
	"reflect"
	"testing"

	"github.com/pthapa1/go-practice-algo/problems"
)

func TestBTCombinationI(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
		target   int
	}{
		{
			input:  []int{2, 3, 5, 7},
			target: 7,
			expected: [][]int{
				{2, 2, 3},
				{2, 5},
				{7},
			},
		},
		{
			input:  []int{1, 2},
			target: 3,
			expected: [][]int{
				{1, 1, 1},
				{1, 2},
			},
		},
		{
			input:  []int{2, 3, 6, 7},
			target: 7,
			expected: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			input:  []int{2, 4, 6, 8},
			target: 8,
			expected: [][]int{
				{2, 2, 2, 2},
				{2, 2, 4},
				{2, 6},
				{4, 4},
				{8},
			},
		},
		{
			input:  []int{1},
			target: 2,
			expected: [][]int{
				{1, 1},
			},
		},
		{
			input:    []int{},
			target:   3,
			expected: [][]int{},
		},
		{
			input:  []int{1, 2, 3},
			target: 4,
			expected: [][]int{
				{1, 1, 1, 1},
				{1, 1, 2},
				{1, 3},
				{2, 2},
			},
		},
	}

	for _, test := range tests {
		result := problems.ComboSum(test.input, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Error On Combination Sum: Expected %v, Got %v", test.expected, result)
		}
	}
}
