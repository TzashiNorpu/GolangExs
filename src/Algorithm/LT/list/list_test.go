package list

import "testing"

func Test_List(t *testing.T) {
	ints := []int{3, 1, 1, 2}
	l := GenLinkedListByInts(ints)
	l = l.DelSpecificValueNode(ListNode{Val: 1})
	l.Traverse()
}
