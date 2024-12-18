package algorithms

import (
	"github.com/UltraGnome/AdventOfCode2024/pkg/math"
	"github.com/UltraGnome/AdventOfCode2024/pkg/sets"
)

// Dijkstra is an implementation of Dijkstra's algorithm for finding the shortest path between two nodes in a graph.
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func Dijkstra[T comparable](g Graph[T], a, b T) int {
	nodeSet, edgeMap := g.optimal()

	if !nodeSet.Contains(a) || !nodeSet.Contains(b) {
		return -1
	}

	distances := make(map[T]int, len(nodeSet))
	for node := range nodeSet {
		distances[node] = math.MaxInt
	}
	distances[a] = 0

	visitable := sets.SetOf(a)
	visited := make(sets.Set[T], len(nodeSet))
	for len(visited) < len(nodeSet) {
		current := a
		minDistance := math.MaxInt
		for node := range visitable {
			if distances[node] < minDistance {
				current = node
				minDistance = distances[node]
			}
		}
		if current == b {
			return distances[current]
		}
		visited.Add(current)
		visitable.Remove(current)
		for node, distance := range edgeMap[current] {
			if !visited.Contains(node) {
				visitable.Add(node)
				newDistance := distances[current] + distance
				distances[node] = min(distances[node], newDistance)
			}
		}
	}

	return -1
}
