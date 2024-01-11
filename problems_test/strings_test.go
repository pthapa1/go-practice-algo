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
