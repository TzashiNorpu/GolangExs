package lt206

import (
	"golang/Algorithm/LT/list"
	"testing"
)

func TestLt206(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	l := list.GenLinkedListByInts(ints)
	l = &list.LinkedList{Node: reverseList(l.Node)}
	l.Traverse()
}
