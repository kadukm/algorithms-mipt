package heaps

type Heap interface {
	Add(int)
	Pop() (int, error)
	GetMin() (int, error)
}
