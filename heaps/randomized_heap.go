package heaps

import (
	"errors"
	"math/rand"
)

type randomizedHeap struct {
	root *node
}

type node struct {
	value int
	l, r  *node
}

func NewRandomizedHeap() Heap {
	return &randomizedHeap{nil}
}

func (h *randomizedHeap) Add(value int) {
	newNode := &node{value, nil, nil}
	h.root = merge(h.root, newNode)
}

func (h *randomizedHeap) Pop() (int, error) {
	if h.root == nil {
		return 0, errors.New("heap is empty")
	}

	result := h.root.value
	h.root = merge(h.root.l, h.root.r)
	return result, nil
}

func (h *randomizedHeap) GetMin() (int, error) {
	if h.root == nil {
		return 0, errors.New("heap is empty")
	}

	return h.root.value, nil
}

func merge(node1, node2 *node) *node {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	if node2.value < node1.value {
		node1, node2 = node2, node1
	}
	if rand.Int()%2 == 1 {
		node1.l, node1.r = node1.r, node1.l
	}

	node1.l = merge(node1.l, node2)
	return node1
}
