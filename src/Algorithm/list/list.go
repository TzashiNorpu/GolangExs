package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListGenByInts(ints []int) *ListNode {
	return listGenByInts(ints, 0)
}
func listGenByInts(ints []int, level int) *ListNode {
	if level == len(ints)-1 {
		return &ListNode{Val: ints[level]}
	}
	curr := ListNode{Val: ints[level]}
	curr.Next = listGenByInts(ints, level+1)
	return &curr
}
