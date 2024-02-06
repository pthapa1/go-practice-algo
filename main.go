package main

import (
	"fmt"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func main() {
	// testHashTable := algo.Init(10)
	testHashTable := algo.Init[int](10)
	fmt.Println(testHashTable)
	fmt.Println("Index: ", algo.Hash("Brady", 10))
}
