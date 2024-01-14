package problems

import (
	"testing"

	"github.com/pthapa1/go-practice-algo/problems"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		original string
		reversed string
	}{
		{original: "hello", reversed: "olleh"},
		{original: "hello world", reversed: "dlrow olleh"},
	}

	for _, v := range testCases {
		result := problems.ReverseString(v.original)
		if result != v.reversed {
			t.Errorf("Reversed string does not match, Expected %v \n Got %v", v.reversed, result)
		}
	}
}

func TestFindConcurrentChar(t *testing.T) {
	type testCase struct {
		testString string
		expected   problems.CharProtocol
	}
	testCases := []testCase{
		{
			testString: "helLo",
			expected: problems.CharProtocol{
				Character:     "l",
				StartingIndex: 2,
				Count:         2,
			},
		},
	}

	for _, tc := range testCases {
		result := problems.FindConcurrentChar(tc.testString)
		if result.Character != tc.expected.Character ||
			result.StartingIndex != tc.expected.StartingIndex ||
			result.Count != tc.expected.Count {
			t.Errorf(
				"Test for concurrent char failed for string %v. Expected %v Got %v",
				tc.testString,
				tc.expected,
				result,
			)
		}
	}
}
