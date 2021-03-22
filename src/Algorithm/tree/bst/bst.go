package bst

type BST struct {
	Node Node
	size int
}
type Node struct {
	Val         int
	Left, Right *Node
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

func (b *BST) Add(node Node) {
	add(b.Node, node)
}

func add(root Node, node Node) Node {
	if &root == nil {
		return node
	}
	if node.Val < root.Val {
		n := add(*root.Left, node)
		root.Left = &n
	}
	if node.Val > root.Val {
		n := add(*root.Right, node)
		root.Right = &n
	}
	return root
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) Contains(node Node) bool {
	return contains(b.Node, node)
}

func contains(root Node, n Node) bool {
	if &root == nil {
		return false
	}
	if root.Val > n.Val {
		return contains(*root.Right, n)
	} else if root.Val < n.Val {
		return contains(*root.Left, n)
	} else {
		return true
	}
}
