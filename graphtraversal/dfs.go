package graphtraversal

import "github.com/kadukm/algorithms-mipt/stack"

func DFS(g Graph, from, to int) Path {
	if from == to {
		return []int{from}
	}

	parent := make([]int, len(g))
	if tryFindPathDFS(g, from, to, parent) {
		return extractPath(parent, from, to)
	}

	return nil
}

func tryFindPathDFS(g Graph, from, to int, parent []int) bool {
	used := make([]bool, len(g))
	used[from] = true

	s := stack.NewStack(16)
	s.Push(from)

	for s.Size() != 0 {
		node, _ := s.Pop()
		for _, nextNode := range g[node] {
			if used[nextNode] {
				continue
			}

			parent[nextNode] = node
			if nextNode == to {
				return true
			}

			used[nextNode] = true
			s.Push(nextNode)
		}
	}

	return false
}
