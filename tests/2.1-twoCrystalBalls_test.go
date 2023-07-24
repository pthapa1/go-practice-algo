package algo

import (
	"github.com/pthapa1/go-practice-algo/algorithms"
	"math/rand"
	"testing"
)

func TestTwoCrystalBallProblem(t *testing.T) {
	limit := 13
	boolArr := make([]bool, limit)
	randomNumber := rand.Intn(limit)
	for i := randomNumber; i < limit; i++ {
		boolArr[i] = true
	}
	testCases := []struct {
		boolArr  []bool
		expected int
	}{
		{make([]bool, 10), -1},
		{boolArr, randomNumber},
	}

	for index, tc := range testCases {
		result := algo.TwoCrystalBalls(tc.boolArr)
		if result != tc.expected {
			t.Errorf("Two Crystal Ball: Expected %v but got %v for %v th array", tc.expected, result, index+1)
		}
	}

}
