package heap

type MaxHeap struct {
	data *[]int
}

func (h *MaxHeap) Add(node int) *MaxHeap {
	n := append(*(h.data), node)
	h.data = &n
	currIndex := len(*h.data) - 1
	parentIndex := getParentIndex(len(*h.data) - 1)
	h.siftUp(currIndex, parentIndex)
	return h
}

func (h *MaxHeap) Heapfy(arr *[]int) *MaxHeap {
	(h.data) = arr
	index := getParentIndex(len(*arr) - 1)
	for i := index; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (h *MaxHeap) siftUp(currIndex int, parentIndex int) {
	for parentIndex >= 0 && (*h.data)[currIndex] > (*h.data)[parentIndex] {
		(*h.data)[currIndex], (*h.data)[parentIndex] = (*h.data)[parentIndex], (*h.data)[currIndex]
		currIndex = parentIndex
		parentIndex = getParentIndex(currIndex)
	}
}
func (h *MaxHeap) RemoveMax() int {
	if len(*h.data) == 0 {
		panic("no element...")
	}
	pop := (*h.data)[0]
	(*h.data)[0], (*h.data)[len(*h.data)-1] = (*h.data)[len(*h.data)-1], (*h.data)[0]
	(*h.data) = (*h.data)[:len(*h.data)-1]
	h.siftDown(0)
	return pop
}

func (h *MaxHeap) siftDown(i int) {
	for getLeftChildIndex(i) < len(*h.data) {
		j := getLeftChildIndex(i)
		if j+1 < len(*h.data) && (*h.data)[j+1] > (*h.data)[j] {
			j = getRightChildIndex(i)
		}
		if (*h.data)[i] > (*h.data)[j] {
			break
		}
		(*h.data)[i], (*h.data)[j] = (*h.data)[j], (*h.data)[i]
		i = j
	}
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftChildIndex(index int) int {
	return index*2 + 1
}

func getRightChildIndex(index int) int {
	return index*2 + 2
}

func GetMaxHeap() *MaxHeap {
	var s []int
	return &MaxHeap{data: &s}
}
