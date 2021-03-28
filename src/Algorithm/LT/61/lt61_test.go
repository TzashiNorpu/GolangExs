package lt61

import (
	"golang/Algorithm/LT/list"
	"testing"
)

func TestLt61(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	l := list.GenLinkedListByInts(ints)
	n := rotateRight(l.Node, 7)
	(&(list.LinkedList{Node: n})).Traverse()
}
