package algo

import (
	"fmt"
	"math"
)

func TwoCrystalBalls(breaks []bool) int {
	var jumpAmount = math.Floor(math.Sqrt(float64(len(breaks))))
	i := int(jumpAmount)
	fmt.Println("value of i =", i)
	for ; i < len(breaks); i = i + int(jumpAmount) {
		if breaks[i] {
			break
		}
	}

	i = i - int(jumpAmount)
	// when the i stops, either because it found true,
	// or because it is at the end of an array
	for j := i; j < len(breaks); j++ {
		if breaks[j] {
			return j
		}
	}
	return -1
}
