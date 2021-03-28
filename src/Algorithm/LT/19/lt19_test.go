package lt19

import (
	"golang/Algorithm/LT/list"
	"testing"
)

func TestLt19(t *testing.T) {
	ints1 := []int{1, 2}
	l := list.GenLinkedListByInts(ints1)
	s := removeNthFromEnd(l.Node, 2)
	(&list.LinkedList{Node: s}).Traverse()
}
