package test1

import (
	"fmt"
	"golang/Algorithm/list"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	ints := []int{1, 2, 3, 3, 4, 4, 5}
	listInts := list.ListGenByInts(ints)
	/*	for listInts != nil {
		fmt.Println(listInts.Val)
		listInts = listInts.Next
	}*/
	n := deleteDuplicates(listInts)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
