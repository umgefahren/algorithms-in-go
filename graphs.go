package algorithms_in_go

import (
	"github.com/yourbasic/graph"
)

const MaxInt64 = int64(9223372036854775807)

func MinIdxArr(inp []int64) int {
	min := MaxInt64
	minIdx := 0
	for idx, el := range inp {
		if el < min {
			min = el
			minIdx = idx
		}
	}
	return minIdx
}

func Dijkstra(g *graph.Mutable, start int, end int) int64 {
	visited := make([]bool, g.Order())
	distance := make([]int64, g.Order())
	for idx, _ := range distance {
		distance[idx] = MaxInt64
	}
	distance[start] = 0
	current := start
	for {
		currentDistance := distance[current]

		g.Visit(current, func(w int, c int64) bool {
			if visited[w] {
				return false
			}
			newDistance := currentDistance + c
			if newDistance < distance[w] {
				distance[w] = newDistance
			}
			return false
		})

		visited[current] = true

		// fmt.Printf("Current => %v\n", current)
		if visited[end] {
			break
		}
		distance[current] = MaxInt64
		// fmt.Printf("Current %v\n", current)
		current = MinIdxArr(distance)
	}
	return distance[end]
}
