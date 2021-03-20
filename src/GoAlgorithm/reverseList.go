package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*输入: 1->2->3->4->5->NULL
		null<-1 2->3->4->5->NULL
输出: 5->4->3->2->1->NULL*/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

/*输入: 1->2->NULL
输出: 2->1->NULL*/
func reverseListByRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	tmp := head
	if curr.Next != nil {
		tmp = reverseListByRecursive(curr.Next)
		curr.Next.Next = curr
		curr.Next = nil
	}
	return tmp
}

func main() {
	node := ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	//node.Next.Next = nil
	node.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next = &ListNode{Val: 4}
	node.Next.Next.Next.Next = &ListNode{Val: 5}
	node.Next.Next.Next.Next.Next = nil
	list := reverseList(&node)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
	fmt.Println("++++++++++++++++++++++++++")
	node = ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	//node.Next.Next = nil
	node.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next = &ListNode{Val: 4}
	node.Next.Next.Next.Next = &ListNode{Val: 5}
	node.Next.Next.Next.Next.Next = nil
	listByRecursive := reverseListByRecursive(&node)
	for listByRecursive != nil {
		fmt.Println(listByRecursive.Val)
		listByRecursive = listByRecursive.Next
	}

}
