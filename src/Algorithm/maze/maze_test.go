package maze

import (
	"fmt"
	"testing"
)

func TestMaze(t *testing.T) {
	m := readMaze("./maze.in")
	fmt.Println(m)
	steps := walk(m, point{
		0, 0,
	}, point{
		len(m) - 1,
		len(m[0]) - 1,
	})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
