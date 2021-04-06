package trap

import (
	"fmt"
	"math"
)

func trap(height []int) int {
	level := 1
	len := len(height)
	l, r := 0, len-1
	vol := 0
	sum := 0
	for l <= r && r < len {
		for l <= r && height[l] < level {
			l++
		}
		for l <= r && height[r] < level {
			r--
		}
		vol = vol + r - l + 1
		level++
	}
	for _, v := range height {
		sum = sum + v
	}

	return vol - sum
}

func trap1(height []int) int {
	leftMax := make([]int, len(height))
	for k, _ := range height {
		for i := 0; i < k; i++ {
			if height[i] > leftMax[k] {
				leftMax[k] = height[i]
			}
		}
	}
	//fmt.Println(leftMax)
	rightMax := make([]int, len(height))
	for k, _ := range height {
		for i := k + 1; i < len(height); i++ {
			if height[i] > rightMax[k] {
				rightMax[k] = height[i]
			}
		}
	}
	//fmt.Println(rightMax)
	res := 0
	areai := 0
	for i := 0; i < len(height); i++ {
		min := int(math.Min(float64(leftMax[i]), float64(rightMax[i])))
		areai = min - height[i]
		if areai < 0 {
			areai = 0
		}
		res = res + areai
	}
	return res
}

// DP
func trap2(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}
	leftMax := make([]int, len(height))
	leftMax[0] = height[0]

	for i := 1; i < len(height); i++ {
		leftMax[i] = int(math.Max(float64(leftMax[i-1]), float64(height[i])))
	}
	fmt.Println(leftMax)

	rightMax := make([]int, len(height))
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = int(math.Max(float64(rightMax[i+1]), float64(height[i])))
	}
	fmt.Println(rightMax)
	res := 0
	areai := 0
	for i := 0; i < len(height); i++ {
		areai = int(math.Min(float64(leftMax[i]), float64(rightMax[i]))) - height[i]
		if areai < 0 {
			areai = 0
		}
		res = res + areai
	}
	return res
}

// 单调栈
func trap3(height []int) int {
	/*	stack := []int{}
		for i, v := range height {

			stack := append(stack, i)
		}*/
	return 0
}
