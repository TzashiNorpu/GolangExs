package bst

import (
	"fmt"
	"testing"
)

func TestBst(t *testing.T) {
	bst := BSTGen()
	bst.Add(Node{Val: 28})
	bst.Add(Node{Val: 30})
	bst.Add(Node{Val: 30})
	bst.Add(Node{Val: 44})
	bst.Add(Node{Val: 25})
	bst.Add(Node{Val: 20})
	bst.Add(Node{Val: 77})
	bst.Add(Node{Val: 80})
	fmt.Println(bst)
	fmt.Println(bst.Contains(Node{Val: 25}))
	fmt.Println(bst.Contains(Node{Val: 45}))
	fmt.Println(bst.Size())
	fmt.Println(bst.IsEmpty())
	fmt.Println()
	bst.InOrder()
	fmt.Println()
	bst.PreOrder()
	fmt.Println()
	bst.PostOrder()
}
