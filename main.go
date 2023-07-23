package main

import (
	// algo "github.com/pthapa1/go-practice-algo/algorithms"
	"fmt"
)

type TestCase struct {
	input    []int
	expected []int
}

func main() {
	// algo.ExecuteQuickSort()

	testCases := []TestCase{
		{[]int{1, 4, 2}, []int{1, 2, 4}},
		{[]int{2, 9, 3}, []int{2, 3, 9}},
	}

	fmt.Println(testCases)
}
