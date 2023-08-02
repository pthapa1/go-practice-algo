// maze solving algorithm with recursion
package algo

// import "fmt"

type Point struct {
	X int
	Y int
}

var direction = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func walk(maze []string, wall string, curr Point, seen [][]bool, end Point, path *[]Point) bool {
	// if the current point is off the map.
	if curr.Y < 0 || curr.X < 0 || curr.Y >= len(maze) || curr.X >= len(maze[curr.Y]) {
		return false
	}
	// if it is a wall.
	if string(maze[curr.Y][curr.X]) == wall {
		return false
	}
	// if we have seen it.
	if seen[curr.Y][curr.X] {
		return false
	}
	// return true when it's the end.
	if curr.X == end.X && curr.Y == end.Y {
		*path = append(*path, end)
		return true
	}

	seen[curr.Y][curr.X] = true

	for i := 0; i < len(direction); i++ {
		x := direction[i][0]
		y := direction[i][1]

		if walk(maze, wall, Point{X: curr.X + x, Y: curr.Y + y}, seen, end, path) {
			*path = append(*path, curr)
			return true
		}
	}

	return false
}

func MazeSolver(maze []string, wall string, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	var path []Point
	walk(maze, wall, start, seen, end, &path)
	// The path will be in reverse order, so we need to reverse it before returning.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
