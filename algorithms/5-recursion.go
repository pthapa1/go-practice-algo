// maze solving algorithm with recursion
package algo

/*
[
'##########E##',
'#           #',
'##S##########'
]
*/

type Point struct {
	x int
	y int
}

var direction = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func walk(maze []string, wall string, curr Point, seen [][]bool, end Point, path *[]Point) bool {
	// if it is a wall.
	if string(maze[curr.y][curr.x]) == wall {
		return false
	}
	// if the current point is off the map.
	if curr.y < 0 || curr.x < 0 || curr.y >= len(maze) || curr.x >= len(maze[curr.y]) {
		return false
	}
	// if we have seen it.
	if seen[curr.y][curr.x] {
		return false
	}
	// return true, when it's the end.
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, end)
		return true
	}

	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	for i := 0; i < len(direction); i++ {
		x := direction[i][0]
		y := direction[i][1]

		if walk(maze, wall, Point{x: curr.x + x, y: curr.y + y}, seen, end, path) {
			*path = append(*path, curr)
			return true
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}

func MazeSover(maze []string, wall string, start Point, end Point) []Point {

	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	var path []Point
	walk(maze, wall, start, seen, end, &path)
	return path
}
