package main

import (
	"fmt"

	"github.com/pthapa1/go-practice-algo/problems"
)

func main() {
	input := []int{2, 1}
	output := problems.ComboSum(input, 3)
	fmt.Println("Combosum", output)
}
