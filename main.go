package main

import (
	"fmt"

	"github.com/pthapa1/go-practice-algo/problems"
)

func main() {
	input := []int{1, 2, 3}
	output := problems.Permute(input)
	fmt.Println("Permutations", output)
}
