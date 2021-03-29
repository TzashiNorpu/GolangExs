package lt203

import (
	"golang/Algorithm/LT/list"
	"testing"
)

func TestLt203(t *testing.T) {
	ints := []int{1, 2, 3}
	l := list.GenLinkedListByInts(ints)
	r := removeElements2(l.Node, 1)
	(&list.LinkedList{Node: r}).Traverse()
}
