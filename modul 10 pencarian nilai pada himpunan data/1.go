package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	weights := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i])
	}

	minWeight, maxWeight := weights[0], weights[0]
	for _, weight := range weights[1:] {
		if weight < minWeight {
			minWeight = weight
		}
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	fmt.Printf("%.2f %.2f\n", minWeight, maxWeight)
}
