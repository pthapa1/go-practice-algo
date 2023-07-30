package main

import "fmt"

// algo "github.com/pthapa1/go-practice-algo/algorithms"
type Point struct {
	x int
	y int
}

func walk(path []Point) {
	fmt.Println(path)
}

func main() {
	myPoint := []Point{{x: 1, y: 1}}
	walk(myPoint)
}
