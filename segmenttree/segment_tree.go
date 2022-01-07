package segmenttree

import (
	"errors"
	"fmt"
	"math"
)

type SegmentTree interface {
	GetSum(findL, findR int) int
	Set(index int, value int) error
}

type segmentTree struct {
	tree  []int
	count int
}

func NewSegmentTree(values []int) SegmentTree {
	log := math.Log2(float64(len(values)))
	log = math.Ceil(log)

	treeLength := int(math.Pow(2, log))*2 - 1
	result := &segmentTree{make([]int, treeLength), len(values)}
	result.makeTree(0, 0, len(values), values)

	return result
}

func (t *segmentTree) makeTree(node int, l, r int, values []int) {
	if l+1 == r {
		t.tree[node] = values[l]
		return
	}

	t.makeTree(node*2+1, l, (l+r)/2, values)
	t.makeTree(node*2+2, (l+r)/2, r, values)
	t.tree[node] = t.tree[node*2+1] + t.tree[node*2+2]
}

func (t *segmentTree) GetSum(findL, findR int) int {
	if findL < 0 {
		findL = 0
	}
	if findR > len(t.tree)+1 {
		findR = len(t.tree)
	}

	return t.getSum(0, 0, t.count, findL, findR)
}

func (t *segmentTree) getSum(node int, l, r int, findL, findR int) int {
	if findL >= r || findR <= l {
		return 0
	}

	if findL <= l && r <= findR {
		return t.tree[node]
	}

	return t.getSum(node*2+1, l, (l+r)/2, findL, findR) +
		t.getSum(node*2+2, (l+r)/2, r, findL, findR)
}

func (t *segmentTree) Set(idx int, value int) error {
	if idx < 0 || idx >= t.count {
		return errors.New(fmt.Sprintf("index out of range error: %d of %d", idx, t.count))
	}

	t.set(0, 0, t.count, idx, value)

	return nil
}

func (t *segmentTree) set(node int, l, r int, idx int, value int) {
	if idx < l || idx >= r {
		return
	}

	if l+1 == r && idx == l {
		t.tree[node] = value
		return
	}

	t.set(node*2+1, l, (l+r)/2, idx, value)
	t.set(node*2+2, (l+r)/2, r, idx, value)
	t.tree[node] = t.tree[node*2+1] + t.tree[node*2+2]
}
