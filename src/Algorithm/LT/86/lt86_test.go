package lt86

import (
	"golang/Algorithm/LT/list"
	"testing"
)

func TestLt86(t *testing.T) {
	ints := []int{1, 2, 4, 3, 2, 5, 2}
	l := list.GenLinkedListByInts(ints)
	s := partition(l.Node, 3)
	(&list.LinkedList{Node: s}).Traverse()
}
