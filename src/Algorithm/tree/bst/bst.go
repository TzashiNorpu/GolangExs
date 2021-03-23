package bst

import "fmt"

type BST struct {
	Node *Node
	size int
}
type Node struct {
	Val         int
	Left, Right *Node
}

func (b *BST) IsEmpty() bool {
	return b.size-1 == 0
}

func (b *BST) Add(node Node) {
	b.add1(b.Node, node)
	//b.Node = b.add2(b.Node, node)
}

func (b *BST) add1(root *Node, node Node) {
	if root == nil {
		b.size++
		b.Node = &node
		//nil 传递只能进行复制，不能修改外部的变量
		//root = node
		return
	} else {
		if root.Val == node.Val {
			return
		} else if node.Val < root.Val && root.Left == nil {
			root.Left = &node
			b.size++
			return
		} else if node.Val > root.Val && root.Right == nil {
			root.Right = &node
			b.size++
			return
		}
	}

	if node.Val < root.Val {
		b.add1(root.Left, node)
	} else if node.Val > root.Val {
		b.add1(root.Right, node)
	}
}

func (b *BST) add2(root *Node, node Node) *Node {
	if root == nil {
		b.size++
		return &node
	}
	if node.Val < root.Val {
		root.Left = b.add2(root.Left, node)
	} else if node.Val > root.Val {
		root.Right = b.add2(root.Right, node)
	}
	return root
}

func (b *BST) Size() int {
	return b.size + 1
}

func (b *BST) Contains(node Node) bool {
	return contains(b.Node, node)
}

func contains(root *Node, n Node) bool {
	if root == nil {
		return false
	}
	if n.Val < root.Val {
		return contains(root.Left, n)
	} else if n.Val > root.Val {
		return contains(root.Right, n)
	} else {
		return true
	}
}

func (b *BST) PreOrder() {
	preOrder(b.Node)
}

func preOrder(node *Node) {
	fmt.Println(node.Val)
	if node.Left != nil {
		preOrder(node.Left)
	}
	if node.Right != nil {
		preOrder(node.Right)
	}
}

func (b *BST) InOrder() {
	inOrder(b.Node)
}

func inOrder(node *Node) {
	if node.Left != nil {
		inOrder(node.Left)
	}
	fmt.Println(node.Val)
	if node.Right != nil {
		inOrder(node.Right)
	}
}

func (b *BST) PostOrder() {
	postOrder(b.Node)
}

func postOrder(node *Node) {
	if node.Left != nil {
		postOrder(node.Left)
	}
	if node.Right != nil {
		postOrder(node.Right)
	}
	fmt.Println(node.Val)
}

func BSTGen() *BST {
	return &BST{
		Node: nil,
		size: 0,
	}
}
