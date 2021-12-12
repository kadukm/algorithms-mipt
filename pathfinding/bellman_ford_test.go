package pathfinding

import "testing"

func TestFindBellmanFordPathInGraphWithoutExistingPath(t *testing.T) {
	graph := Graph{
		{{2, 8}},
		{},
		{},
		{{0, -2}, {4, 4}},
		{{0, -9}},
	}

	result := FindBellmanFordPath(graph, 3, 1)

	if result.exist {
		t.Fatalf("path from 3 to 1 doesn't exist but %v was found", result.path)
	}
}

func TestFindBellmanFordPathInGraphWithIsolatedLastNode(t *testing.T) {
	graph := Graph{
		{{1, 3}},
		{{2, -4}},
		{{1, 6}},
	}

	result := FindBellmanFordPath(graph, 2, 0)

	if result.exist {
		t.Fatalf("path from 2 to 0 doesn't exist but %v was found", result.path)
	}
}

func TestFindBellmanFordPathInSmallGraph(t *testing.T) {
	graph := Graph{
		{{1, 10}, {2, 20}},
		{},
		{{1, -25}},
	}

	testFindBellmanFordPath(t, graph, 0, 1, -5)
}

func TestFindBellmanFordPathInGraphWithOneNode(t *testing.T) {
	graph := Graph{
		{},
	}

	testFindBellmanFordPath(t, graph, 0, 0, 0)
}

func TestFindBellmanFordPathInTwoNodesGraph(t *testing.T) {
	graph := Graph{
		{{1, -9}},
		{{0, 18}},
	}

	testFindBellmanFordPath(t, graph, 0, 1, -9)
}

func TestFindBellmanFordPathInGraphWithTwoMinPaths(t *testing.T) {
	graph := Graph{
		{{1, 10}, {2, 60}, {3, 15}},
		{{2, 40}, {3, 30}},
		{{3, -35}},
		{},
	}

	testFindBellmanFordPath(t, graph, 0, 3, 15)
}

func TestFindBellmanFordPathInBigGraph(t *testing.T) {
	graph := Graph{
		{{4, 7}},
		{{0, 4}, {3, 9}, {6, 8}},
		{{0, 2}, {1, 10}, {5, 30}, {7, 40}},
		{{0, 6}, {5, 11}, {6, 33}, {7, 20}},
		{{5, 9}},
		{},
		{{5, 1}},
		{{1, 3}, {6, -50}},
	}

	testFindBellmanFordPath(t, graph, 2, 5, -10)
}

func TestFindBellmanFordPathWhenNegativeCycleInResultPath(t *testing.T) {
	graph := Graph{
		{{1, 1}},
		{{2, 1}},
		{{3, -1}, {4, 1}},
		{{1, -1}},
		{},
	}

	result := FindBellmanFordPath(graph, 0, 4)

	if result.exist {
		t.Fatalf("path from 0 to 4 infinite due to negative cycle but pat %v was found", result.path)
	}
}

func TestFindBellmanFordPathInTwoNodesGraphWithNegativeCycle(t *testing.T) {
	graph := Graph{
		{{1, -2}},
		{{0, 1}},
	}

	result := FindBellmanFordPath(graph, 0, 1)

	if result.exist {
		t.Fatalf("path from 0 to 1 infinite due to negative cycle but pat %v was found", result.path)
	}
}

func TestFindBellmanFordPathWhenNegativeCycleNotInResultPath(t *testing.T) {
	graph := Graph{
		{{1, 10}},
		{},
		{{3, -1}},
		{{4, -1}},
		{{2, -1}},
	}

	testFindBellmanFordPath(t, graph, 0, 1, 10)
}

func testFindBellmanFordPath(t *testing.T, g Graph, from, to int, expectedWeight int) {
	result := FindBellmanFordPath(g, from, to)

	testPathIsCorrect(t, g, from, to, expectedWeight, result)
}
