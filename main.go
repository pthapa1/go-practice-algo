package main

import (
	"fmt"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func main() {
	m := &algo.MaxHeap{}
	fmt.Println(m)
	items := []int{1, 5, 7, 8, 9, 14, 16, 30, 34, 50}

	for _, v := range items {
		m.Insert(v)
	}
	fmt.Println("New heap", m.Slice)
}
