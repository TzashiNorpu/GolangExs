package lt100

import (
	"golang/Algorithm/LT/tree"
)

func isSameTree(p *tree.TreeNode, q *tree.TreeNode) bool {
	/*	if p == nil && q == nil {
		return true
	}*/
	var a []interface{}
	inOrder(p, func(n interface{}) {
		a = append(a, n)
	})
	var b []interface{}
	inOrder(q, func(n interface{}) {
		b = append(b, n)
	})
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func inOrder(node *tree.TreeNode, f func(n interface{})) {
	if node == nil {
		return
	}
	if node.Left != nil {
		inOrder(node.Left, f)
	}
	f(node.Val)
	if node.Right != nil {
		inOrder(node.Right, f)
	}
}

func isSameTree2(p *tree.TreeNode, q *tree.TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if p.Left != nil && q.Left != nil {
		return isSameTree2(p.Left, q.Left)
	}
	if p.Val != q.Val {
		return false
	}
	if p.Right != nil && q.Right != nil {
		return isSameTree2(p.Right, q.Right)
	}
	return true
	/*
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return isSameTree2(p.Left, q.Left) && isSameTree2(p.Right, q.Right)
	*/
}
