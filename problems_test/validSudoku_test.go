package problems

import (
	"fmt"
	"testing"

	"github.com/pthapa1/go-practice-algo/problems"
)

func TestValidSudoku(t *testing.T) {
	tests := []struct {
		input    [][]string
		expected bool
	}{
		{
			input: [][]string{
				{"5", "3", ".", ".", "7", ".", ".", ".", "."},
				{"6", ".", ".", "1", "9", "5", ".", ".", "."},
				{".", "9", "8", ".", ".", ".", ".", "6", "."},
				{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
				{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", "6", ".", ".", ".", ".", "2", "8", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "5"},
				{".", ".", ".", ".", "8", ".", ".", "7", "9"},
			},
			expected: true,
		},
		{
			input: [][]string{
				{"5", "3", ".", ".", "7", ".", ".", ".", "."},
				{"6", ".", ".", "1", "9", "5", ".", ".", "."},
				{".", "9", "8", ".", ".", ".", ".", "6", "."},
				{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
				{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", "6", ".", ".", ".", ".", "2", "8", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "5"},
				{".", ".", ".", ".", "8", ".", ".", "7", "8"}, // Invalid: "8" repeated in last row
			},
			expected: false,
		},
		{
			input: [][]string{
				{"8", "3", ".", ".", "7", ".", ".", ".", "."},
				{"6", ".", ".", "1", "9", "5", ".", ".", "."},
				{".", "9", "8", ".", ".", ".", ".", "6", "."},
				{
					"8",
					".",
					".",
					".",
					"6",
					".",
					".",
					".",
					"3",
				}, // Invalid: "8" repeated in first column
				{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", "6", ".", ".", ".", ".", "2", "8", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "5"},
				{".", ".", ".", ".", "8", ".", ".", "7", "9"},
			},
			expected: false,
		},
		{
			input: [][]string{
				{"5", "3", ".", ".", "7", ".", ".", ".", "."},
				{"6", ".", ".", "1", "9", "5", ".", ".", "."},
				{".", "9", "8", ".", ".", ".", ".", "6", "."},
				{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
				{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", "6", ".", ".", ".", ".", "2", "8", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "5"},
				{
					"5",
					".",
					".",
					".",
					"8",
					".",
					".",
					"7",
					"9",
				}, // Invalid: "5" repeated in first column
			},
			expected: false,
		},
		{
			input: [][]string{
				{"5", "3", ".", ".", "7", ".", ".", ".", "."},
				{"6", ".", ".", "1", "9", "5", ".", ".", "."},
				{".", "9", "8", ".", ".", ".", ".", "6", "."},
				{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
				{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", "6", ".", ".", ".", ".", "2", "8", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "5"},
				{".", ".", ".", ".", "8", ".", ".", "7", "."},
			},
			expected: true,
		},
		{
			input: [][]string{
				{".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", "."},
				{".", ".", ".", ".", ".", ".", ".", ".", "."},
			},
			expected: true,
		},
		{
			input: [][]string{
				{"5", "3", ".", ".", "7", ".", ".", ".", "."},
				{"6", ".", ".", "1", "9", "5", ".", ".", "."},
				{".", "9", "8", ".", ".", ".", ".", "6", "."},
				{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
				{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
				{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
				{".", "6", ".", ".", ".", ".", "2", "8", "."},
				{".", ".", ".", "4", "1", "9", ".", ".", "5"},
				{
					"5",
					".",
					".",
					".",
					"8",
					".",
					".",
					"7",
					"9",
				}, // Invalid: Duplicate in last row and first column
			},
			expected: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			output := problems.ValidSudoku(test.input)
			if output != test.expected {
				t.Errorf("Test %d failed: got %v, expected %v", i+1, output, test.expected)
			}
		})
	}
}
