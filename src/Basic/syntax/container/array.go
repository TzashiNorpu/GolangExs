package main

import (
	"fmt"
)

func printArray(arr *[5]int) {
	(*arr)[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func nonRepeatSubstr(s string) int {
	maxLen := 0
	start := 0
	lastOccured := map[rune]int{}
	for i, v := range []rune(s) {
		if index, ok := lastOccured[v]; ok && index >= start {
			start = lastOccured[v] + 1
		}
		lastOccured[v] = i
		if i+1-start > maxLen {
			maxLen = i + 1 - start
		}
	}
	return maxLen
}

func main() {

	/*	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
		s1 := arr[2:6]
		fmt.Println(s1)
		s2 := s1[3:5]
		fmt.Println(s2)
		s3 := append(s2, 10)
		s4 := append(s3, 11)
		s5 := append(s4, 12)
		fmt.Println("s3, s4 ,s5 = ", s3, s4, s5)
		fmt.Println(arr)*/

	/*	arr := [5]int{1, 2, 3, 4, 5}
		printArray(&arr)
		fmt.Println(arr)*/

	s1 := make([]int, 12)
	s2 := []int{1, 2, 4}
	copy(s1, s2)
	fmt.Println(s1)
	fmt.Println(nonRepeatSubstr("我爱你你你你"))
	// 结构体

	// 通过值调用值接收者方法
	root := treeNode{value: 3}
	root.setValue(2)
	fmt.Println(root)
	// 通过引用调用值接收者方法
	root.left = &treeNode{5, nil, nil}
	root.left.setValue(2)
	fmt.Println(root)
	fmt.Println(root.left)

	// 通过值调用指针接收者方法
	root.setValueX(2)
	fmt.Println(root)
	// 通过引用调用指针接收者方法
	root.left.setValueX(2)
	fmt.Println(root)
	fmt.Println(root.left)

}

type treeNode struct {
	value int
	// left和right为指针类型
	left, right *treeNode
}

// 值接收者方法
func (node treeNode) setValue(value int) {
	node.value = value
}

// 指针接收者方法
func (node *treeNode) setValueX(value int) {
	node.value = value
}
