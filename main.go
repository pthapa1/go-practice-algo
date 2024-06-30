package main

import (
	"fmt"

	"github.com/pthapa1/go-practice-algo/problems"
)

func main() {
	input := []int{2, 3, 5, 7}
	output := problems.ComboSum(input, 7)
	fmt.Println("Combosum", output)
}
