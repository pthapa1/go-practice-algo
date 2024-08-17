package problems

import (
	"reflect"
	"testing"

	"github.com/pthapa1/go-practice-algo/problems"
)

func TestComboSum2(t *testing.T) {
	tests := []struct {
		candidates []int
		expected   [][]int
		target     int
	}{
		{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected: [][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected: [][]int{
				{7},
			},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			expected: [][]int{
				{3, 5},
			},
		},
		{
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			candidates: []int{},
			target:     3,
			expected:   [][]int{},
		},
		{
			candidates: []int{1},
			target:     1,
			expected: [][]int{
				{1},
			},
		},
		{
			candidates: []int{1, 1, 1, 1},
			target:     2,
			expected: [][]int{
				{1, 1},
			},
		},
	}

	for _, test := range tests {
		result := problems.ComboSum2(test.candidates, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For candidates %v and target %d, expected %v but got %v",
				test.candidates, test.target, test.expected, result)
		}
	}
}
