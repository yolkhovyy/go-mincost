package main

import (
	"fmt"
	"math"
)

// See https://www.geeksforgeeks.org/min-cost-path-dp-6/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCost(cost [][]int, m int, n int) int {
	if n < 0 || m < 0 {
		return math.MaxInt64
	} else if m == 0 && n == 0 {
		return cost[m][n]
	} else {
		return cost[m][n] +
			min(minCost(cost, m-1, n),
				min(minCost(cost, m, n-1),
					minCost(cost, m-1, n-1)))
	}
}

func main() {
	tiles := [][]int{
		{5, 0, 8, 3},
		{0, 2, 1, 3},
		{1, 0, 9, 4},
		{0, 1, 0, 5}}

	fmt.Printf("%d", minCost(tiles, 3, 3))
}
