package pathfinding

type Edge struct {
	to     int
	weight int
}
type Graph [][]Edge

type WeightedPath struct {
	path   []int
	weight int
	exist  bool
}
