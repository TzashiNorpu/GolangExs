package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertTobin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("Error")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(convertTobin(4))
	fmt.Println(convertTobin(13))

	printFile("./syntax/branch/abc.txt")
}
