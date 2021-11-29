package graphtraversal

func BFS(g Graph, from, to int) Path {
	if from == to {
		return []int{from}
	}

	parent := make([]int, len(g))
	if tryFindPathBFS(g, from, to, parent) {
		return extractPath(parent, from, to)
	}

	return nil
}

func tryFindPathBFS(g Graph, from, to int, parent []int) bool {
	used := make([]bool, len(g))
	used[from] = true

	currentLayer := make([]int, 0, 16)
	currentLayer = append(currentLayer, from)

	for len(currentLayer) != 0 {
		nextLayer := make([]int, 0, 16)

		for _, node := range currentLayer {
			for _, nextNode := range g[node] {
				if used[nextNode] {
					continue
				}

				parent[nextNode] = node
				if nextNode == to {
					return true
				}

				used[nextNode] = true
				nextLayer = append(nextLayer, nextNode)
			}
		}

		currentLayer = nextLayer
	}

	return false
}
