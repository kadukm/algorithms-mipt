package pathfinding

import "math"

func FindBellmanFordPath(g Graph, from, to int) WeightedPath {
	weights := createWeights(g, from)
	parent := make([]int, len(g))

	for i := 0; i < len(g); i++ {
		for node := 0; node < len(g); node++ {
			if weights[node] == math.MaxInt {
				continue
			}

			for _, edge := range g[node] {
				weight := weights[node] + edge.weight
				if weight < weights[edge.to] {
					parent[edge.to] = node
					weights[edge.to] = weight
				}
			}
		}
	}

	cycled := make([]bool, len(g))
	for i := 0; i < len(g); i++ {
		for node := 0; node < len(g); node++ {
			if weights[node] == math.MaxInt {
				continue
			}

			for _, edge := range g[node] {
				weight := weights[node] + edge.weight
				if weight < weights[edge.to] {
					cycled[edge.to] = true
					weights[edge.to] = weight
				}
			}
		}
	}

	if cycled[to] || weights[to] == math.MaxInt {
		return WeightedPath{nil, -1, false}
	}

	return extractPath(parent, from, to, weights[to])
}
