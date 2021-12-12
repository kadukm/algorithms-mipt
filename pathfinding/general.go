package pathfinding

import "math"

type Edge struct {
	to     int
	weight int
}
type Graph [][]Edge

type WeightedPath struct {
	path   []int
	weight int
	exist  bool
}

func createWeights(g Graph, from int) []int {
	weights := make([]int, len(g))
	for i := 0; i < len(g); i++ {
		weights[i] = math.MaxInt
	}
	weights[from] = 0

	return weights
}

func extractPath(parent []int, from, to int, weight int) WeightedPath {
	reversedPath := make([]int, 0, 16)
	reversedPath = append(reversedPath, to)

	node := to
	for node != from {
		node = parent[node]
		reversedPath = append(reversedPath, node)
	}

	path := make([]int, 0, len(reversedPath))
	for i := len(reversedPath) - 1; i >= 0; i-- {
		path = append(path, reversedPath[i])
	}

	return WeightedPath{path, weight, true}
}
