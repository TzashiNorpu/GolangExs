package main

import (
	"fmt"
	"golang/Algorithm/tree/bst"
)

func main() {
	r := bst.BSTGen()
	r.Add(bst.Node{Val: 55})
	r.Add(bst.Node{Val: 35})
	r.Add(bst.Node{Val: 23})
	r.Add(bst.Node{Val: 12})
	r.Add(bst.Node{Val: 67})
	r.Add(bst.Node{Val: 61})
	r.Add(bst.Node{Val: 72})
	r.Add(bst.Node{Val: 81})
	r.Add(bst.Node{Val: 89})
	r.Add(bst.Node{Val: 102})
	r.InOrder()
	fmt.Println()
	r.DelNode(bst.Node{Val: 67})
	r.InOrder()
	fmt.Println()
	r.LevelTraverse()
	fmt.Println()
}
