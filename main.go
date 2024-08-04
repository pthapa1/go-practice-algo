package main

import (
	"fmt"

	"github.com/pthapa1/go-practice-algo/problems"
)

func main() {
	input := []int{3, 1, 3, 5, 1, 1}
	output := problems.ComboSum2(input, 8)
	fmt.Println("Output1", output)
}
