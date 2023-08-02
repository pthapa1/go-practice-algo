package algo

import (
	"reflect"
	"testing"

	algo "github.com/pthapa1/go-practice-algo/algorithms"
)

func TestRecursion(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}
	mazeResult := []algo.Point{
		{X: 10, Y: 0},
		{X: 10, Y: 1},
		{X: 10, Y: 2},
		{X: 10, Y: 3},
		{X: 10, Y: 4},
		{X: 9, Y: 4},
		{X: 8, Y: 4},
		{X: 7, Y: 4},
		{X: 6, Y: 4},
		{X: 5, Y: 4},
		{X: 4, Y: 4},
		{X: 3, Y: 4},
		{X: 2, Y: 4},
		{X: 1, Y: 4},
		{X: 1, Y: 5},
	}

	start := algo.Point{X: 10, Y: 0}
	end := algo.Point{X: 1, Y: 5}
	path := algo.MazeSolver(maze, "x", start, end)

	if !reflect.DeepEqual(path, mazeResult) {
		t.Errorf("Boy, you have an error on Maze Solving Algo. Expected %v : Got %v", mazeResult, path)
	}
}
