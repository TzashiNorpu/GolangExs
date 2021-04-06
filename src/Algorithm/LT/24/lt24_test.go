package lt24

import (
	"golang/Algorithm/LT/list"
	"testing"
)

func TestLt24(t *testing.T) {
	ints := []int{1, 2, 3}
	l := list.GenLinkedListByInts(ints)
	s := swapPairs(l.Node)
	(&list.LinkedList{Node: s}).Traverse()
}
