package lt21

import (
	"golang/Algorithm/LT/list"
	"testing"
)

func TestLt21(t *testing.T) {
	ints1 := []int{3, 3, 4}
	l1 := list.GenLinkedListByInts(ints1)
	ints2 := []int{1, 2, 5}
	l2 := list.GenLinkedListByInts(ints2)
	n := mergeTwoLists2(l1.Node, l2.Node)
	(&(list.LinkedList{Node: n})).Traverse()
}
