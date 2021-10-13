package stack

import "errors"

type Stack interface {
	Pop() (int, error)
	Push(value int)
	Peek() (int, error)
	Size() int
}

type stack struct {
	data []int
}

func NewStack(capacity int) Stack {
	return &stack{make([]int, 0, capacity)}
}

func (s *stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("attempt to pop from empty stack is incorrect")
	}

	result := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return result, nil
}

func (s *stack) Peek() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("attempt to peek from empty stack is incorrect")
	}

	return s.data[len(s.data)-1], nil
}

func (s *stack) Size() int {
	return len(s.data)
}
