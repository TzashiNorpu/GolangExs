package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	ints1 := []interface{}{1, nil, 2}
	l1 := GenTreesByArray(ints1)
	var a []interface{}
	inOrder(l1, func(n interface{}) {
		a = append(a, n)
	})
	fmt.Println(a)
	ints2 := []interface{}{1, 2}
	l2 := GenTreesByArray(ints2)
	var b []interface{}
	inOrder(l2, func(n interface{}) {
		b = append(b, n)
	})
	fmt.Println(b)
}
