package graphtraversal

type Graph [][]int
type Path []int

func extractPath(parent []int, from, to int) Path {
	reversedPath := make(Path, 0, 16)
	reversedPath = append(reversedPath, to)

	node := to
	for node != from {
		node = parent[node]
		reversedPath = append(reversedPath, node)
	}

	path := make(Path, 0, len(reversedPath))
	for i := len(reversedPath) - 1; i >= 0; i-- {
		path = append(path, reversedPath[i])
	}

	return path
}
