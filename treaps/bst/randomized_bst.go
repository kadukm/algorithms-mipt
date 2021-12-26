package bst

import "math/rand"

type Treap struct {
	root *node
}

type node struct {
	value    int
	priority int
	l, r     *node
}

func (t *Treap) Add(value int) {
	newNode := &node{value, rand.Int(), nil, nil}
	l, r := split(t.root, value)
	r = merge(newNode, r)
	t.root = merge(l, r)
}

func (t *Treap) Delete(value int) bool {
	l, r := split(t.root, value)
	l, m := split(l, value-1)
	t.root = merge(l, r)
	return m != nil
}

func (t *Treap) Contains(value int) bool {
	l, r := split(t.root, value)
	l, m := split(l, value-1)
	result := m != nil
	l = merge(l, m)
	t.root = merge(l, r)
	return result
}

func (t *Treap) GetSortedValues() []int {
	result := make([]int, 0, 8)
	bfs(t.root, &result)
	return result
}

func bfs(n *node, result *[]int) {
	if n == nil {
		return
	}
	bfs(n.l, result)
	*result = append(*result, n.value)
	bfs(n.r, result)
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
		return l
	} else {
		r.l = merge(l, r.l)
		return r
	}
}

func split(n *node, value int) (*node, *node) {
	if n == nil {
		return nil, nil
	}

	if n.value <= value {
		l, r := split(n.r, value)
		n.r = l
		return n, r
	} else {
		l, r := split(n.l, value)
		n.l = r
		return l, n
	}
}
