package heaps

import "errors"

type binaryHeap struct {
	data []int
}

func NewBinaryHeap() Heap {
	return &binaryHeap{make([]int, 0, 8)}
}

func (h *binaryHeap) Add(value int) {
	h.data = append(h.data, value)
	h.siftUp()
}

func (h *binaryHeap) siftUp() {
	index := len(h.data) - 1
	for {
		if index == 0 {
			break
		}

		nextIndex := (index - 1) / 2
		if h.data[nextIndex] <= h.data[index] {
			break
		}

		h.data[index], h.data[nextIndex] = h.data[nextIndex], h.data[index]
		index = nextIndex
	}
}

func (h *binaryHeap) Pop() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("heap is empty")
	}

	result := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.siftDown()

	return result, nil
}

func (h *binaryHeap) siftDown() {
	index := 0
	for {
		minIndex := index

		l := index*2 + 1
		if l < len(h.data) && h.data[l] < h.data[minIndex] {
			minIndex = l
		}

		r := index*2 + 2
		if r < len(h.data) && h.data[r] < h.data[minIndex] {
			minIndex = r
		}

		if minIndex == index {
			break
		}

		h.data[index], h.data[minIndex] = h.data[minIndex], h.data[index]
		index = minIndex
	}
}

func (h *binaryHeap) GetMin() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("heap is empty")
	}

	return h.data[0], nil
}
