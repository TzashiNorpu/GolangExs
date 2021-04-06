package lt100

import (
	"fmt"
	"golang/Algorithm/LT/tree"
	"testing"
)

func TestLt100(t *testing.T) {
	ints1 := []interface{}{1, 2}
	ints2 := []interface{}{1, nil, 2}
	l1 := tree.GenTreesByArray(ints1)
	l2 := tree.GenTreesByArray(ints2)
	fmt.Println(isSameTree(l1, l2))
	fmt.Println(isSameTree2(l1, l2))
}
