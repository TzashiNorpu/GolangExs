package bst

import "fmt"

type bst struct {
	node *Node
	size int
}
type Node struct {
	Val         int
	Left, Right *Node
}

func (b *bst) IsEmpty() bool {
	return b.size-1 == 0
}

func (b *bst) Add(node Node) {
	b.add1(b.node, node)
	//b.node = b.add2(b.node, node)
}

func (b *bst) add1(root *Node, node Node) {
	if root == nil {
		b.size++
		b.node = &node
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

func (b *bst) add2(root *Node, node Node) *Node {
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

func (b *bst) Size() int {
	return b.size + 1
}

func (b *bst) Contains(node Node) bool {
	return contains(b.node, node)
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

func (b *bst) PreOrder() {
	preOrder(b.node)
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

func (b *bst) InOrder() {
	inOrder(b.node)
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

func (b *bst) PostOrder() {
	postOrder(b.node)
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

func (b *bst) LevelTraverse() {
	levelTraverse(b.node)
}

func levelTraverse(node *Node) {
	if node == nil {
		return
	}
	var stack []*Node
	stack = append(stack, node)
	for len(stack) > 0 {
		curNode := stack[0]
		fmt.Println(curNode.Val)
		stack = stack[1:]
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
	}
}

func (b *bst) GetMin() *Node {
	return getMin(b.node)
}

func getMin(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return getMin(node.Left)
}

func (b *bst) GetMax() *Node {
	return getMax(b.node)
}

func getMax(node *Node) *Node {
	if node.Right == nil {
		return node
	}
	return getMax(node.Right)
}

func (b *bst) DelMin() *Node {
	return b.delMin(b.node)
}

func (b *bst) delMin(node *Node) *Node {
	if node.Left == nil {
		b.size--
		return node.Right
	}
	node.Left = b.delMin(node.Left)
	return node
}

func (b *bst) DelMax() *Node {
	return b.delMax(b.node)
}

func (b *bst) delMax(node *Node) *Node {
	if node.Right == nil {
		b.size--
		return node.Left
	}
	node.Right = b.delMax(node.Right)
	return node
}

func (b *bst) DelNode(node Node) *Node {
	return b.delNode(b.node, node)
}

func (b *bst) delNode(root *Node, node Node) *Node {

	if root == nil {
		return nil
	}
	if node.Val == root.Val {
		if root.Left == nil {
			b.size--
			return root.Right
		} else if root.Right == nil {
			b.size--
			return root.Left
		} else {
			successor := getMin(root.Right)
			sub := b.delMin(root.Right)
			successor.Right = sub
			successor.Left = root.Left
			return successor
		}
	}
	if node.Val < root.Val {
		root.Left = b.delNode(root.Left, node)
	}
	if node.Val > root.Val {
		root.Right = b.delNode(root.Right, node)
	}
	return root
}

func BSTGen() *bst {
	return &bst{
		node: nil,
		size: 0,
	}
}
