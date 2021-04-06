package tree

import (
	"fmt"
)

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

// [0,1,2,3,4,5,6,7,8,9,10,12,13,14,15,16]
// level 0
// i = 0 left  1 right  2
// level 1
// i = 1 left  3 right  4
// i = 2 left  5 right  6
// level 2
// i = 3 left 7 right  8
// i = 4 left 9 right  10
// i = 5 left 11 right 12
// i = 6 left 13 right  14
// level 3
// i = 7 left 15 right  16
// i = 8 left 17 right  18
// i = 9 left 19 right 20
// i = 10 left 21 right 22
// i = 11 left 23 right  24
// i = 12 left 25 right  26
// i = 13 left 27 right  28
// i = 14 left 29 right  30

func GenTreesByArray(ints []interface{}) *TreeNode {
	/*	left = 2^k - 1
		right = 2^(k+1) - 2 */
	nodeMap := make(map[interface{}]*TreeNode)
	var node *TreeNode
	for k, v := range ints {
		node1, ok := nodeMap[k]
		if !ok {
			node = &TreeNode{Val: v}
			nodeMap[k] = node
		} else {
			node = node1
		}
		if 2*k+1 < len(ints) {
			left := &TreeNode{Val: ints[2*k+1]}
			nodeMap[2*k+1] = left
			(*node).Left = left
		}
		if 2*k+2 < len(ints) {
			right := &TreeNode{Val: ints[2*k+2]}
			nodeMap[2*k+2] = right
			(*node).Right = right
		}
	}
	return nodeMap[0]
}

func preOrder(node *TreeNode) {
	fmt.Println(node.Val)
	if node.Left != nil {
		preOrder(node.Left)
	}
	if node.Right != nil {
		preOrder(node.Right)
	}
}

func inOrder(node *TreeNode, f func(n interface{})) {
	if node.Left != nil {
		inOrder(node.Left, f)
	}
	f(node.Val)
	//fmt.Println(node.Val)
	if node.Right != nil {
		inOrder(node.Right, f)
	}
}

func postOrder(node *TreeNode) {
	if node.Left != nil {
		postOrder(node.Left)
	}
	if node.Right != nil {
		postOrder(node.Right)
	}
	fmt.Println(node.Val)
}
