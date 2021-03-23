package maze

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

func (p1 point) add(p2 point) point {
	return point{p1.i + p2.i,
		p1.j + p2.j}
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) in(grid [][]int) (val int, in bool) {
	if p.i >= len(grid) || p.i < 0 {
		return 0, false
	}
	if p.j >= len(grid[0]) || p.j < 0 {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(m [][]int, start_point point, end_point point) [][]int {
	steps := make([][]int, len(m))
	for i := range steps {
		steps[i] = make([]int, len(m[i]))
	}
	q := []point{start_point}
	for len(q) >= 1 {
		point := q[0]
		q = q[1:]
		for _, dir := range dirs {
			next := point.add(dir)
			val, ok := next.in(m)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.in(steps)
			if !ok || val != 0 {
				continue
			}
			// 第二个节点为0时没有该条件判断会往回走
			if next == start_point {
				continue
			}
			val, _ = point.in(steps)
			steps[next.i][next.j] = val + 1
			q = append(q, next)
		}
	}
	return steps
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var rows, cols int
	fmt.Fscanf(file, "%d %d\n", &rows, &cols)
	fmt.Printf("rows=%d,cols=%d\n", rows, cols)
	maze := make([][]int, rows)
	for row := range maze {
		maze[row] = make([]int, cols)
		for col := range maze[row] {
			_, err := fmt.Fscanf(file, "%d", &maze[row][col])
			if err != nil {
				panic(err)
			}
		}
		fmt.Fscanf(file, "\n")
	}
	return maze
}
