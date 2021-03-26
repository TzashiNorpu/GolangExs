package heap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := GetMaxHeap()
	n := 1000000
	for i := 0; i < n; i++ {
		item := rand.Intn(n)
		heap.Add(item)
	}
	var test []int
	for i := 0; i < n; i++ {
		item := heap.RemoveMax()
		test = append(test, item)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(test[i])
	}
	fmt.Println(len(test))
	for i := 0; i < n-1; i++ {
		i1 := test[i]
		i2 := test[i+1]
		if i1 < i2 {
			fmt.Println("testing failure")
			break
		}
	}
	fmt.Println("testing success")
	h2 := GetMaxHeap()
	arr := []int{60, 40, 50, 70, 10, 20}
	h2.Heapfy(&arr)

	fmt.Println(h2.RemoveMax())
	fmt.Println(h2.RemoveMax())
	fmt.Println(h2.RemoveMax())
	fmt.Println(h2.RemoveMax())
	fmt.Println(h2.RemoveMax())
	fmt.Println(h2.RemoveMax())
}
