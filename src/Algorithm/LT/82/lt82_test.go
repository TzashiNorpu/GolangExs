package lt82

import (
	"golang/Algorithm/LT/list"
	"testing"
)

func TestLt82(t *testing.T) {
	ints := []int{1, 1, 2, 3}
	l := list.GenLinkedListByInts(ints)
	nodes := deleteDuplicates2(l.Node)
	l = &list.LinkedList{Node: nodes}
	l.Traverse()
}
