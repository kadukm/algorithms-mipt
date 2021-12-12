package pathfinding

import "math"

type FloydWarshallPath struct {
	weight int
	exist  bool
	cycled bool
}

func FindFloydWarshallPaths(g Graph) [][]FloydWarshallPath {
	weights := createFloydWarshallWeights(g)

	for i := 0; i < len(g); i++ {
		for u := 0; u < len(g); u++ {
			for v := 0; v < len(g); v++ {
				if weights[u][i] == math.MaxInt || weights[i][v] == math.MaxInt {
					continue
				}

				weight := weights[u][i] + weights[i][v]
				if weight < weights[u][v] {
					weights[u][v] = weight
				}
			}
		}
	}

	result := createFloydWarshallPaths(g, weights)

	for i := 0; i < len(g); i++ {
		for u := 0; u < len(g); u++ {
			for v := 0; v < len(g); v++ {
				if weights[u][i] == math.MaxInt || weights[i][v] == math.MaxInt {
					continue
				}

				weight := weights[u][i] + weights[i][v]
				if weight < weights[u][v] {
					result[u][v].cycled = true
					weights[u][v] = weight
				}
			}
		}
	}

	return result
}

func createFloydWarshallWeights(g Graph) [][]int {
	weights := make([][]int, len(g))

	for node := 0; node < len(g); node++ {
		weights[node] = createWeights(g, node)
	}

	for node := 0; node < len(g); node++ {
		weights[node][node] = 0
		for _, edge := range g[node] {
			weights[node][edge.to] = edge.weight
		}
	}

	return weights
}

func createFloydWarshallPaths(g Graph, weights [][]int) [][]FloydWarshallPath {
	result := make([][]FloydWarshallPath, len(g))

	for u := 0; u < len(g); u++ {
		result[u] = make([]FloydWarshallPath, len(g))
		for v := 0; v < len(g); v++ {
			if weights[u][v] == math.MaxInt {
				result[u][v] = FloydWarshallPath{-1, false, false}
			} else {
				result[u][v] = FloydWarshallPath{weights[u][v], true, false}
			}
		}
	}

	return result
}
