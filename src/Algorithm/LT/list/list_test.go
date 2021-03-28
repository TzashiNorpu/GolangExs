package list

import "testing"

func Test_List(t *testing.T) {
	ints := []int{1, 1, 2, 3, 3}
	l := GenLinkedListByInts(ints)
	l.Traverse()
}
