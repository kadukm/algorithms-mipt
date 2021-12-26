package implicit

import (
	"errors"
	"fmt"
	"math/rand"
)

type ImplicitTreap struct {
	root *node
}

type node struct {
	value    int
	count    int
	priority int
	l, r     *node
}

func (t *ImplicitTreap) Add(value int) {
	newNode := &node{value, 1, rand.Int(), nil, nil}
	t.root = merge(t.root, newNode)
}

func (t *ImplicitTreap) Insert(value int, index int) error {
	if err := t.checkIndex(index, true); err != nil {
		return err
	}

	newNode := &node{value, 1, rand.Int(), nil, nil}
	l, r := split(t.root, index)
	l = merge(l, newNode)
	t.root = merge(l, r)
	return nil
}

func (t *ImplicitTreap) Delete(index int) error {
	if err := t.checkIndex(index, false); err != nil {
		return err
	}

	l, r := split(t.root, index)
	_, r = split(r, 1)
	t.root = merge(l, r)

	return nil
}

func (t *ImplicitTreap) Get(index int) (int, error) {
	if err := t.checkIndex(index, false); err != nil {
		return 0, err
	}

	l, r := split(t.root, index)
	m, r := split(r, 1)
	result := m.value
	r = merge(m, r)
	t.root = merge(l, r)

	return result, nil
}

func (t *ImplicitTreap) checkIndex(index int, allowRightBoundary bool) error {
	count := getCount(t.root)
	if index < 0 || count < index || (!allowRightBoundary && count == index) {
		return errors.New(fmt.Sprintf("index out of range: %d of %d", index, count))
	}
	return nil
}

func merge(l, r *node) *node {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}

	if l.priority <= r.priority {
		l.r = merge(l.r, r)
		l.count = getCount(l.l) + 1 + getCount(l.r)
		return l
	} else {
		r.l = merge(l, r.l)
		r.count = getCount(r.l) + 1 + getCount(r.r)
		return r
	}
}

func split(n *node, count int) (*node, *node) {
	if n == nil {
		return nil, nil
	}

	if getCount(n.l) < count {
		l, r := split(n.r, count-getCount(n.l)-1)
		n.r = l
		n.count = getCount(n.l) + 1 + getCount(n.r)
		return n, r
	} else {
		l, r := split(n.l, count)
		n.l = r
		n.count = getCount(n.l) + 1 + getCount(n.r)
		return l, n
	}
}

func getCount(n *node) int {
	if n == nil {
		return 0
	}
	return n.count
}
