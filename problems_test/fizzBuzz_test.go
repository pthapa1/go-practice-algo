package problems

import (
	"reflect"
	"testing"

	"github.com/pthapa1/go-practice-algo/problems"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		name          string
		expectedArray []string
		number        int
	}{
		{
			name:   "Standard Case",
			number: 15,
			expectedArray: []string{
				"1", "2", "Fizz", "4", "Buzz",
				"Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz",
			},
		},
		{
			name:          "Zero Case",
			number:        0,
			expectedArray: []string{},
		},
		{
			name:          "Negative Case",
			number:        -3,
			expectedArray: []string{},
		},
		{
			name:   "Large Number Case",
			number: 20,
			expectedArray: []string{
				"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := problems.FizzBuzz(tc.number)
			if !reflect.DeepEqual(tc.expectedArray, result) {
				t.Errorf(
					"Test %v failed. \nExpected %v \nGot %v",
					tc.name,
					tc.expectedArray,
					result,
				)
			}
		})
	}
}
