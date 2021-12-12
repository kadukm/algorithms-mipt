package pathfinding

import "testing"

func TestFindFloydWarshallPathsInGraphWhereAllPathsExist(t *testing.T) {
	graph := Graph{
		{{1, 1}, {2, 6}},
		{{2, 7}},
		{{3, -1}},
		{{0, -4}, {1, 6}},
	}

	expected := [][]FloydWarshallPath{
		{exist(0), exist(1), exist(6), exist(5)},
		{exist(2), exist(0), exist(7), exist(6)},
		{exist(-5), exist(-4), exist(0), exist(-1)},
		{exist(-4), exist(-3), exist(2), exist(0)},
	}

	actual := FindFloydWarshallPaths(graph)

	comparePaths(t, expected, actual)
}

func TestFindFloydWarshallPathsInGraphWhereSomePathsNotExistOrCycled(t *testing.T) {
	graph := Graph{
		{},
		{{0, 9}, {2, 6}},
		{{4, -1}},
		{{2, -1}},
		{{3, -1}},
		{{4, 9}, {6, 6}},
		{},
	}

	expected := [][]FloydWarshallPath{
		{exist(0), notExist(), notExist(), notExist(), notExist(), notExist(), notExist()},
		{exist(9), exist(0), cycled(), cycled(), cycled(), notExist(), notExist()},
		{notExist(), notExist(), cycled(), cycled(), cycled(), notExist(), notExist()},
		{notExist(), notExist(), cycled(), cycled(), cycled(), notExist(), notExist()},
		{notExist(), notExist(), cycled(), cycled(), cycled(), notExist(), notExist()},
		{notExist(), notExist(), cycled(), cycled(), cycled(), exist(0), exist(6)},
		{notExist(), notExist(), notExist(), notExist(), notExist(), notExist(), exist(0)},
	}

	actual := FindFloydWarshallPaths(graph)

	comparePaths(t, expected, actual)
}

func exist(weight int) FloydWarshallPath {
	return FloydWarshallPath{weight, true, false}
}

func notExist() FloydWarshallPath {
	return FloydWarshallPath{-1, false, false}
}

func cycled() FloydWarshallPath {
	return FloydWarshallPath{-1, true, true}
}

func comparePaths(t *testing.T, expected, actual [][]FloydWarshallPath) {
	for u := 0; u < len(expected); u++ {
		for v := 0; v < len(expected); v++ {
			if expected[u][v].exist && !actual[u][v].exist {
				t.Fatalf("path from %d to %d exists but was not found %v", u, v, actual[u][v])
			}
			if !expected[u][v].exist && actual[u][v].exist {
				t.Fatalf("path from %d to %d does not exist but was found %v", u, v, actual[u][v])
			}

			if expected[u][v].cycled && !actual[u][v].cycled {
				t.Fatalf("path from %d to %d is cycled by was found %v", u, v, actual[u][v])
			}
			if !expected[u][v].cycled && actual[u][v].cycled {
				t.Fatalf("path from %d to %d is not cycled by was found %v", u, v, actual[u][v])
			}

			if expected[u][v].exist && !expected[u][v].cycled && expected[u][v].weight != actual[u][v].weight {
				t.Fatalf("path from %d to %d has weight %d but was found %v",
					u, v, expected[u][v].weight, actual[u][v])
			}
		}
	}
}
