package algo

import (
	"fmt"
	"math"
	"math/rand"
)

func two_cyrstal_balls(breaks []bool) int {

	var jumpAmount = math.Floor(math.Sqrt(float64(len(breaks))))

	i := int(jumpAmount)

	for ; i < len(breaks); i = i + int(jumpAmount) {
		if breaks[i] {
			break
		}
	}

	i = i - int(jumpAmount)

	for j := 0; j < int(jumpAmount) && j < len(breaks); j++ {
		i++
		if breaks[i] {
			return i
		}
	}

	return -1

}

func ExecuteTwoCrystalBalls() {
	// create an array of boolean of 10,000. By default all of them are false.
	data := make([]bool, 10000)
	idx := rand.Intn(10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	fmt.Println(two_cyrstal_balls(data), idx)

}
