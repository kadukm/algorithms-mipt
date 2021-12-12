package pathfinding

import "testing"

func TestFindDijkstraPathInGraphWithoutExistingPath(t *testing.T) {
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

func TestFindDijkstraPathInGraphWithIsolatedLastNode(t *testing.T) {
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

func TestFindDijkstraPathInSmallGraph(t *testing.T) {
	graph := Graph{
		{{1, 9}, {2, 1}},
		{{3, 3}},
		{{1, 2}},
		{},
	}

	testFindDijkstraPath(t, graph, 0, 3, 6)
}

func TestFindDijkstraPathInGraphWithOneNode(t *testing.T) {
	graph := Graph{
		{},
	}

	testFindDijkstraPath(t, graph, 0, 0, 0)
}

func TestFindDijkstraPathInTwoNodesGraph(t *testing.T) {
	graph := Graph{
		{{1, 9}},
		{{0, 8}},
	}

	testFindDijkstraPath(t, graph, 0, 1, 9)
}

func TestFindDijkstraPathInGraphWithTwoMinPaths(t *testing.T) {
	graph := Graph{
		{{1, 5}, {2, 6}, {3, 9}},
		{{3, 5}},
		{{3, 3}},
		{},
	}

	testFindDijkstraPath(t, graph, 0, 3, 9)
}

func TestFindDijkstraPathInBigGraph(t *testing.T) {
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

	testPathIsCorrect(t, g, from, to, expectedWeight, result)
}
