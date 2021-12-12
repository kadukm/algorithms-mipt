package pathfinding

import "testing"

func testPathIsCorrect(t *testing.T, g Graph, from int, to int, expectedWeight int, result WeightedPath) {
	weight, exist := checkPath(g, result)
	if !exist {
		t.Fatalf("resulting path %v doesn't exist in graph", result.path)
	}
	if result.path[0] != from || result.path[len(result.path)-1] != to {
		t.Fatalf("resulting path should be from %d to %d but found %v", from, to, result.path)
	}
	if result.weight != weight {
		t.Fatalf("presented weight %d for path %v but actual weight is %d", result.weight, result.path, weight)
	}
	if result.weight != expectedWeight {
		t.Fatalf("resulting path %v has weight %d but expected %d", result.path, weight, expectedWeight)
	}
}

func checkPath(g Graph, p WeightedPath) (int, bool) {
	weight := 0
	for i := 0; i < len(p.path)-1; i++ {
		node := p.path[i]
		nextNode := p.path[i+1]

		nextWeight, exist := getWeight(g, node, nextNode)
		if !exist {
			return -1, false
		}

		weight += nextWeight
	}

	return weight, true
}

func getWeight(g Graph, node, nextNode int) (int, bool) {
	for _, edge := range g[node] {
		if edge.to == nextNode {
			return edge.weight, true
		}
	}

	return -1, false
}
