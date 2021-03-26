package lt83

import (
	list2 "golang/Algorithm/LT/list"
	"testing"
)

func TestLt83(t *testing.T) {
	ints := []int{1, 2, 3}
	l := list2.GenLinkedListByInts(ints)
	l.AddFirst(list2.ListNode{Val: 100})
	/*	for listInts != nil {
		fmt.Println(listInts.Val)
		listInts = listInts.Next
	}*/
	/*n := deleteDuplicates(listInts)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}*/
}
