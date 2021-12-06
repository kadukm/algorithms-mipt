package pathfinding

import "testing"

func TestGraphWithoutExistingPath(t *testing.T) {
	graph := Graph{
		{{2, 8}},
		{},
		{},
		{{0, 2}, {4, 4}},
		{{0, 9}},
	}

	result := FindDijkstraPath(graph, 3, 1)

	if result.exist {
		t.Fatalf("path from 3 to 1 doesn't exist but %v was found", result.path)
	}
}

func TestGraphWithIsolatedLastNode(t *testing.T) {
	graph := Graph{
		{{1, 3}},
		{{2, 4}},
		{{1, 6}},
	}

	result := FindDijkstraPath(graph, 2, 0)

	if result.exist {
		t.Fatalf("path from 2 to 0 doesn't exist but %v was found", result.path)
	}
}

func TestSmallGraph(t *testing.T) {
	graph := Graph{
		{{1, 9}, {2, 1}},
		{{3, 3}},
		{{1, 2}},
		{},
	}

	testFindDijkstraPath(t, graph, 0, 3, 6)
}

func TestGraphWithOneNode(t *testing.T) {
	graph := Graph{
		{},
	}

	testFindDijkstraPath(t, graph, 0, 0, 0)
}

func TestTwoNodesGraph(t *testing.T) {
	graph := Graph{
		{{1, 9}},
		{{0, 8}},
	}

	testFindDijkstraPath(t, graph, 0, 1, 9)
}

func TestGraphWithTwoMinPaths(t *testing.T) {
	graph := Graph{
		{{1, 5}, {2, 6}, {3, 9}},
		{{3, 5}},
		{{3, 3}},
		{},
	}

	testFindDijkstraPath(t, graph, 0, 3, 9)
}

func TestBigGraph(t *testing.T) {
	graph := Graph{
		{{1, 6}, {2, 3}, {7, 100}},
		{{2, 8}, {3, 7}, {4, 5}},
		{{5, 9}},
		{{4, 2}, {6, 3}},
		{{5, 4}, {6, 1}},
		{{1, 6}, {4, 3}, {6, 5}, {7, 10}},
		{{7, 2}},
		{{4, 11}},
	}

	testFindDijkstraPath(t, graph, 0, 7, 14)
}

func testFindDijkstraPath(t *testing.T, g Graph, from, to int, expectedWeight int) {
	result := FindDijkstraPath(g, from, to)

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
