package lt2

import (
	"golang/Algorithm/LT/list"
	"testing"
)

func Test(t *testing.T) {
	ints1 := []int{9, 9, 9, 9, 9, 9, 9}
	l1 := list.GenLinkedListByInts(ints1)
	ints2 := []int{9, 9, 9, 9}
	l2 := list.GenLinkedListByInts(ints2)
	l := addTwoNumbers(l1.Node, l2.Node)
	(&list.LinkedList{Node: l}).Traverse()
	l = addTwoNumbers2(l1.Node, l2.Node)
	(&list.LinkedList{Node: l}).Traverse()
}
