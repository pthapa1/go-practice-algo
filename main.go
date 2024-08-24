package main

import (
	"fmt"

	"github.com/pthapa1/go-practice-algo/problems"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	trueOrFalse := problems.RecursiveBinarySearch(arr, -1)
	fmt.Println(trueOrFalse)
	problems.LoopBinarySearch(arr, -1)
	fmt.Println(trueOrFalse)
}
