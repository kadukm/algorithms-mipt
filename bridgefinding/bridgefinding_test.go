package bridgefinding

import "testing"

func TestEmptyGraph(t *testing.T) {
	graph := Graph{}
	expected := []Bridge{}

	actual := FindBridges(graph)

	if !bridgeArraysEqual(expected, actual) {
		t.Fatalf("wrong bridges found: expected %v but was %v", expected, actual)
	}
}

func TestOneNodeGraph(t *testing.T) {
	graph := Graph{
		{},
	}
	expected := []Bridge{}

	actual := FindBridges(graph)

	if !bridgeArraysEqual(expected, actual) {
		t.Fatalf("wrong bridges found: expected %v but was %v", expected, actual)
	}
}

func TestGraphWithoutBridges(t *testing.T) {
	graph := Graph{
		{1, 2},
		{0, 2},
		{0, 1},
	}
	expected := []Bridge{}

	actual := FindBridges(graph)

	if !bridgeArraysEqual(expected, actual) {
		t.Fatalf("wrong bridges found: expected %v but was %v", expected, actual)
	}
}

func TestSmallGraph(t *testing.T) {
	graph := Graph{
		{1},
		{0},
	}
	expected := []Bridge{
		{0, 1},
	}

	actual := FindBridges(graph)

	if !bridgeArraysEqual(expected, actual) {
		t.Fatalf("wrong bridges found: expected %v but was %v", expected, actual)
	}
}

func TestMediumGraph(t *testing.T) {
	graph := Graph{
		{1, 2},
		{0, 2},
		{0, 1, 3},
		{2, 4, 5, 6, 7},
		{3, 5},
		{3, 4, 9},
		{3, 7},
		{3, 6, 8},
		{7},
		{5},
	}
	expected := []Bridge{
		{2, 3},
		{5, 9},
		{7, 8},
	}

	actual := FindBridges(graph)

	if !bridgeArraysEqual(expected, actual) {
		t.Fatalf("wrong bridges found: expected %v but was %v", expected, actual)
	}
}

func TestLargeGraph(t *testing.T) {
	graph := Graph{
		{1, 2},
		{0, 2, 4},
		{0, 1, 3},
		{2},
		{1, 5, 7, 8},
		{4, 6, 8},
		{5, 7, 8},
		{4, 6, 8},
		{4, 5, 6, 7, 9},
		{8, 10},
		{9, 11, 13},
		{10, 12},
		{11, 13},
		{10, 12},
	}
	expected := []Bridge{
		{2, 3},
		{1, 4},
		{8, 9},
		{9, 10},
	}

	actual := FindBridges(graph)

	if !bridgeArraysEqual(expected, actual) {
		t.Fatalf("wrong bridges found: expected %v but was %v", expected, actual)
	}
}

func bridgeArraysEqual(bridges1, bridges2 []Bridge) bool {
	if len(bridges1) != len(bridges2) {
		return false
	}

	for _, bridge := range bridges1 {
		if !bridgesContainBridge(bridges2, bridge) {
			return false
		}
	}

	return true
}

func bridgesContainBridge(bridges []Bridge, bridge Bridge) bool {
	for _, b := range bridges {
		if bridgesEqual(bridge, b) {
			return true
		}
	}
	return false
}

func bridgesEqual(b1, b2 Bridge) bool {
	return (b1.node1 == b2.node1 && b1.node2 == b2.node2) ||
		(b1.node1 == b2.node2 && b1.node2 == b2.node1)
}
