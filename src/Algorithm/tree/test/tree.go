package main

import (
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
}

/*          4
	2				6
1		3		5		7*/
func main() {
	// 先访问根节点再访问左右节点 -->前序
	// 先访问左节点，然后访问根节点，再访问右节点 -->中序
	// 先访问左右节点再访问根节点 -->后序
	// 根节点位置的不同
	tree := treeGen()
	fmt.Print("Inorder:")
	inOrder(tree)
	fmt.Println()
	fmt.Print("Preorder:")
	preOrder(tree)
	fmt.Println()
	fmt.Print("Postorder:")
	postOrder(tree)
	fmt.Println()
}

func preOrder(node node) {
	//root -> left -> right
	fmt.Printf("%d ", node.val)
	if node.left != nil {
		preOrder(*node.left)
	}
	if node.right != nil {
		preOrder(*node.right)
	}
}

func inOrder(node node) {
	//left -> root -> right
	if node.left != nil {
		inOrder(*node.left)
	}
	fmt.Printf("%d ", node.val)
	if node.right != nil {
		inOrder(*node.right)
	}
}
func postOrder(node node) {
	//left -> right -> root
	if node.left != nil {
		postOrder(*node.left)
	}
	if node.right != nil {
		postOrder(*node.right)
	}
	fmt.Printf("%d ", node.val)
}

func treeGen() node {
	root := node{
		val: 4,
	}
	root.left = &node{
		val: 2,
	}
	root.left.left = &node{
		val: 1,
	}
	root.left.right = &node{
		val: 3,
	}

	root.right = &node{
		val: 6,
	}
	root.right.left = &node{
		val: 5,
	}
	root.right.right = &node{
		val: 7,
	}
	return root
}
