package list

import "errors"

type List interface {
	AddLeft(int)
	AddRight(int)
	RemoveLeft() (int, error)
	RemoveRight() (int, error)
	GetLeft() (int, error)
	GetRight() (int, error)
}

type list struct {
	left  *node
	right *node
}

type node struct {
	value int
	left  *node
	right *node
}

func NewList() List {
	return new(list)
}

func (l *list) AddLeft(i int) {
	if l.left == nil {
		l.left = &node{i, nil, nil}
		l.right = l.left
	} else {
		oldLeft := l.left
		l.left = &node{i, nil, oldLeft}
		oldLeft.left = l.left
	}
}

func (l *list) AddRight(i int) {
	if l.right == nil {
		l.left = &node{i, nil, nil}
		l.right = l.left
	} else {
		oldRight := l.right
		l.right = &node{i, oldRight, nil}
		oldRight.right = l.right
	}
}

func (l *list) RemoveLeft() (int, error) {
	if l.left == nil {
		return 0, errors.New("list is empty")
	}

	value := l.left.value

	if l.left.right == nil {
		l.left = nil
		l.right = nil
	} else {
		l.left = l.left.right
		l.left.left = nil
	}

	return value, nil
}

func (l *list) RemoveRight() (int, error) {
	if l.right == nil {
		return 0, errors.New("list is empty")
	}

	value := l.right.value

	if l.right.left == nil {
		l.right = nil
		l.left = nil
	} else {
		l.right = l.right.left
		l.right.right = nil
	}

	return value, nil
}

func (l *list) GetLeft() (int, error) {
	if l.left == nil {
		return 0, errors.New("list is empty")
	}

	return l.left.value, nil
}

func (l *list) GetRight() (int, error) {
	if l.right == nil {
		return 0, errors.New("list is empty")
	}

	return l.right.value, nil
}
