package disjointset

import (
	"errors"
	"fmt"
)

type DisjointSet interface {
	MakeSet(int) error
	Union(int, int) error
	Find(int) (int, error)
}

type disjointSet struct {
	data  map[int]int
	count map[int]int
}

func NewDisjointSet() DisjointSet {
	return &disjointSet{make(map[int]int), make(map[int]int)}
}

func (s *disjointSet) MakeSet(value int) error {
	if r, ok := s.data[value]; ok {
		return errors.New(fmt.Sprintf("value %d already in set with representative %d", value, r))
	}

	s.data[value] = value
	s.count[value] = 1

	return nil
}

func (s *disjointSet) Union(value1, value2 int) error {
	r1, ok := s.data[value1]
	if !ok {
		return errors.New(fmt.Sprintf("value %d is not in any set", value1))
	}

	r2, ok := s.data[value2]
	if !ok {
		return errors.New(fmt.Sprintf("value %d is not in any set", value2))
	}

	if r1 == r2 {
		return nil
	}

	if s.count[r1] < s.count[r2] {
		s.data[r1] = r2
		s.count[r2] += s.count[r1]
	} else {
		s.data[r2] = r1
		s.count[r1] += s.count[r2]
	}

	return nil
}

func (s *disjointSet) Find(value int) (int, error) {
	_, ok := s.data[value]
	if !ok {
		return 0, errors.New(fmt.Sprintf("value %d is not in any set", value))
	}

	return s.findInternal(value), nil
}

func (s *disjointSet) findInternal(value int) int {
	r, _ := s.data[value]
	if value == r {
		return r
	}

	s.data[value] = s.findInternal(r)
	return s.data[value]
}
