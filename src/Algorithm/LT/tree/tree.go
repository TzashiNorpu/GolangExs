package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenTreesByArray(ints []int) *TreeNode {
	/*	left = 2^k - 1
		right = 2^(k+1) - 2 */
}

// [1,2,3]
// level =

func preOrder(node *TreeNode) {
	fmt.Println(node.Val)
	if node.Left != nil {
		preOrder(node.Left)
	}
	if node.Right != nil {
		preOrder(node.Right)
	}
}

func inOrder(node *TreeNode) {
	if node.Left != nil {
		inOrder(node.Left)
	}
	fmt.Println(node.Val)
	if node.Right != nil {
		inOrder(node.Right)
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
