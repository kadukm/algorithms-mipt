package bridgefinding

type Graph [][]int
type Bridge struct {
	node1, node2 int
}

// FindBridges finds bridges in connected graph
func FindBridges(g Graph) []Bridge {
	result := make([]Bridge, 0, 16)

	if len(g) > 1 {
		visited := make([]int, len(g))
		depths := make([]int, len(g))

		dfs(g, 0, -1, 0, visited, depths, &result)
	}

	return result
}

func dfs(g Graph, node, parentNode int, depth int, visited []int, depths []int, result *[]Bridge) int {
	visited[node] = 1
	depths[node] = depth

	for _, nextNode := range g[node] {
		if nextNode == parentNode {
			continue
		}

		if visited[nextNode] == 0 {
			nextNodeLevel := dfs(g, nextNode, node, depth+1, visited, depths, result)
			if nextNodeLevel > depth {
				*result = append(*result, Bridge{node, nextNode})
			}
			depths[node] = min(depths[node], nextNodeLevel)
		} else if visited[nextNode] == 1 {
			depths[node] = min(depths[node], depths[nextNode])
		}
	}

	visited[node] = 2
	return depths[node]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
