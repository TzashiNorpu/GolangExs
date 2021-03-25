package bst

import (
	"fmt"
	"testing"
)

func TestBst(t *testing.T) {
	bst := BSTGen()
	bst.Add(Node{Val: 30})
	bst.Add(Node{Val: 20})
	bst.Add(Node{Val: 25})
	bst.Add(Node{Val: 40})
	bst.Add(Node{Val: 33})
	bst.Add(Node{Val: 42})
	bst.Add(Node{Val: 41})
	bst.Add(Node{Val: 45})
	fmt.Println(bst)
	/*fmt.Println(bst.Contains(node{Val: 25}))
	fmt.Println(bst.Contains(node{Val: 45}))
	fmt.Println(bst.Size())
	fmt.Println(bst.IsEmpty())
	fmt.Println()
	bst.InOrder()
	fmt.Println()
	bst.PreOrder()
	fmt.Println()
	bst.PostOrder()*/
	//bst.LevelTraverse()
	/*min := bst.GetMin()
	fmt.Println(min)
	max := bst.GetMax()
	fmt.Println(max)*/
	/*	bst.DelMin()
		bst.InOrder()
		fmt.Println(bst)
		bst.DelMax()
		bst.InOrder()
		fmt.Println(bst)*/

	bst.InOrder()
	fmt.Println(bst)
	bst.DelNode(Node{Val: 40})
	/*
				30
		20				40
			25		33		42
						41		45
	*/
	/*
				30
		20				41
			25		33		42
								45
	*/
	bst.InOrder()
	fmt.Println(bst)
	/*
				30
		20				42
			25		33		45

	*/
	bst.DelNode(Node{Val: 41})
	bst.InOrder()
	fmt.Println(bst)
	bst.Add(Node{Val: 32})
	fmt.Println()
	bst.InOrder()
}
