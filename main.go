package main

import (
	"fmt"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func main() {
	testHashTable := algo.Init()
	fmt.Println(testHashTable)
	fmt.Println("Index: ", algo.Hash("Brady", 10))
}
