package main

import (
	"fmt"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func main() {
	testHashTable := algo.Init[int](10)
	fmt.Println(testHashTable)
	fmt.Println("Index: ", algo.Hash("Brady", 10))
	fmt.Println("Index: ", algo.Hash(1889181, 10))
}
