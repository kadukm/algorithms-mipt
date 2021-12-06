package pathfinding

import "math"

func FindDijkstraPath(g Graph, from, to int) WeightedPath {
	parent := make([]int, len(g))
	visited := make([]bool, len(g))
	weights := createDijkstraWeights(g, from)

	for {
		node, exists := findMinWeightNode(g, weights, visited)
		if !exists {
			return WeightedPath{nil, -1, false}
		}
		if node == to {
			return extractPath(parent, from, to, weights[to])
		}

		for _, edge := range g[node] {
			weight := weights[node] + edge.weight
			if weight < weights[edge.to] {
				weights[edge.to] = weight
				parent[edge.to] = node
			}
		}

		visited[node] = true
	}
}

func createDijkstraWeights(g Graph, from int) []int {
	weights := make([]int, len(g))
	for i := 0; i < len(g); i++ {
		weights[i] = math.MaxInt
	}
	weights[from] = 0

	return weights
}

func findMinWeightNode(g Graph, weights []int, visited []bool) (int, bool) {
	minWeight := math.MaxInt
	minWeightIdx := -1

	for i := 0; i < len(g); i++ {
		if weights[i] < minWeight && !visited[i] {
			minWeightIdx = i
			minWeight = weights[i]
		}
	}

	if minWeightIdx == -1 {
		return minWeightIdx, false
	} else {
		return minWeightIdx, true
	}
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
