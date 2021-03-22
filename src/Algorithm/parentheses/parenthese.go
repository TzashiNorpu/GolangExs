package main

import "fmt"

func main() {
	genarateP(0, 0, 3, "")
}

func genarateP(left int, right int, len int, res string) {
	if left == len && right == len {
		fmt.Println(res)
		return
	}
	if left < len {
		genarateP(left+1, right, len, res+"(")
	}
	if right < left {
		genarateP(left, right+1, len, res+")")
	}
}
