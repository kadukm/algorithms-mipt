package graphtraversal

import "testing"

type PathSearch struct {
	search   func(Graph, int, int) Path
	name     string
	shortest bool
}

var pathSearches = []PathSearch{
	{BFS, "BFS", true},
	{DFS, "DFS", false},
}

type ExistingPathTestCase struct {
	graph      Graph
	from, to   int
	pathLength int
}

type NotExistingPathTestCase struct {
	graph    Graph
	from, to int
}

func TestSearchNotExistingPath(t *testing.T) {
	notConnectedGraph := Graph{
		{1},
		{0, 2, 3},
		{1},
		{1, 4},
		{3},
		{6, 8},
		{5, 7},
		{6, 8},
		{5, 7},
		{},
	}
	testCases := []NotExistingPathTestCase{
		{notConnectedGraph, 9, 0},
		{notConnectedGraph, 9, 1},
		{notConnectedGraph, 9, 2},
		{notConnectedGraph, 9, 3},
		{notConnectedGraph, 9, 4},
		{notConnectedGraph, 9, 5},
		{notConnectedGraph, 9, 6},
		{notConnectedGraph, 9, 7},
		{notConnectedGraph, 9, 8},

		{notConnectedGraph, 0, 8},
		{notConnectedGraph, 1, 5},
		{notConnectedGraph, 2, 6},
		{notConnectedGraph, 3, 7},
		{notConnectedGraph, 4, 8},
		{notConnectedGraph, 3, 9},
		{notConnectedGraph, 2, 9},

		{notConnectedGraph, 5, 9},
		{notConnectedGraph, 8, 9},
		{notConnectedGraph, 8, 1},
		{notConnectedGraph, 5, 2},
		{notConnectedGraph, 7, 4},
		{notConnectedGraph, 6, 0},
		{notConnectedGraph, 8, 0},
	}

	for _, testCase := range testCases {
		testSearchNotExistingPath(t, testCase)
	}
}

func testSearchNotExistingPath(t *testing.T, tc NotExistingPathTestCase) {
	for _, pathSearch := range pathSearches {
		path := pathSearch.search(tc.graph, tc.from, tc.to)

		if path != nil {
			t.Fatalf("%s: path from %d to %d does not exist but path %v was found",
				pathSearch.name, tc.from, tc.to, path)
		}
	}
}

func TestSingleNodeGraph(t *testing.T) {
	singleNodeGraph := Graph{
		{},
	}
	testCase := ExistingPathTestCase{singleNodeGraph, 0, 0, 1}

	testSearchExistingPath(t, testCase)
}

func TestSmallTreeTraversal(t *testing.T) {
	smallTree := Graph{
		{1, 2},
		{0},
		{0},
	}
	testCases := []ExistingPathTestCase{
		{smallTree, 0, 0, 1},
		{smallTree, 0, 1, 2},
		{smallTree, 0, 2, 2},
		{smallTree, 1, 2, 3},
		{smallTree, 2, 1, 3},
	}

	for _, testCase := range testCases {
		testSearchExistingPath(t, testCase)
	}
}

func TestTreeTraversals(t *testing.T) {
	bigTree := Graph{
		{1, 6, 7},
		{0, 2, 3},
		{1},
		{1, 4, 5},
		{3},
		{3},
		{0},
		{0, 8, 9, 11},
		{7},
		{7, 10},
		{9},
		{7},
	}
	testCases := []ExistingPathTestCase{
		{bigTree, 0, 11, 3},
		{bigTree, 0, 10, 4},
		{bigTree, 0, 5, 4},
		{bigTree, 4, 5, 3},
		{bigTree, 4, 10, 7},
		{bigTree, 10, 4, 7},
		{bigTree, 2, 5, 4},
		{bigTree, 1, 0, 2},
		{bigTree, 7, 1, 3},
	}

	for _, testCase := range testCases {
		testSearchExistingPath(t, testCase)
	}
}

func TestSmallCyclicalGraph(t *testing.T) {
	simpleCyclicalGraph := Graph{
		{1, 2},
		{0, 2},
		{0, 1},
	}
	testCases := []ExistingPathTestCase{
		{simpleCyclicalGraph, 0, 1, 2},
		{simpleCyclicalGraph, 0, 2, 2},
		{simpleCyclicalGraph, 2, 1, 2},
		{simpleCyclicalGraph, 1, 2, 2},
		{simpleCyclicalGraph, 1, 1, 1},
		{simpleCyclicalGraph, 2, 0, 2},
	}

	for _, testCase := range testCases {
		testSearchExistingPath(t, testCase)
	}
}

func TestBigCyclicalGraph(t *testing.T) {
	bigCyclicalGraph := Graph{
		{2, 3},
		{3},
		{0, 3, 5, 6},
		{0, 1, 2, 4, 5},
		{3, 7, 8},
		{2, 3, 6, 7, 9},
		{2, 5, 9},
		{3, 4, 5, 8, 9},
		{4, 7, 9},
		{5, 6, 7, 8},
	}
	testCases := []ExistingPathTestCase{
		{bigCyclicalGraph, 0, 1, 3},
		{bigCyclicalGraph, 0, 9, 4},
		{bigCyclicalGraph, 9, 0, 4},
		{bigCyclicalGraph, 5, 5, 1},
		{bigCyclicalGraph, 5, 4, 3},
		{bigCyclicalGraph, 2, 4, 3},
		{bigCyclicalGraph, 1, 6, 4},
		{bigCyclicalGraph, 0, 5, 3},
		{bigCyclicalGraph, 1, 8, 4},
		{bigCyclicalGraph, 7, 5, 2},
		{bigCyclicalGraph, 2, 6, 2},
		{bigCyclicalGraph, 3, 1, 2},
		{bigCyclicalGraph, 3, 8, 3},
	}

	for _, testCase := range testCases {
		testSearchExistingPath(t, testCase)
	}
}

func testSearchExistingPath(t *testing.T, tc ExistingPathTestCase) {
	for _, pathSearch := range pathSearches {
		path := pathSearch.search(tc.graph, tc.from, tc.to)

		if path == nil {
			t.Fatalf("%s: path from %d to %d exists be found but it is not found",
				pathSearch.name, tc.from, tc.to)
		}
		if pathSearch.shortest && len(path) != tc.pathLength {
			t.Fatalf("%s: wrong length of path from %d to %d: expected %d but was %d",
				pathSearch.name, tc.from, tc.to, tc.pathLength, len(path))
		}
		if !pathCorrect(tc.graph, path) {
			t.Fatalf("%s: path %v from %d to %d does not exist",
				pathSearch.name, path, tc.from, tc.to)
		}
	}
}

func pathCorrect(g Graph, path Path) bool {
	for i := 0; i < len(path)-1; i++ {
		if !nodesAdjacent(g, path[i], path[i+1]) {
			return false
		}
	}
	return true
}

func nodesAdjacent(g Graph, nodeFrom, nodeTo int) bool {
	for _, node := range g[nodeFrom] {
		if node == nodeTo {
			return true
		}
	}
	return false
}
