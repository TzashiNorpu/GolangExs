package bst

import (
	"fmt"
	"testing"
)

func TestBst(t *testing.T) {
	bst := BST{Node: Node{Val: 28}}
	bst.Add(Node{Val: 20})
	bst.Add(Node{Val: 30})
	fmt.Println(bst)
	/*var n1 = Node{Val: 28}
	r.Add(n1)
	fmt.Println(r)*/
}
