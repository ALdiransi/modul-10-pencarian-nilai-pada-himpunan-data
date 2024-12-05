package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	weights := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	totalWeights := []float64{}
	sum := 0.0
	count := 0

	for _, weight := range weights {
		sum += weight
		count++
		if count == y {
			totalWeights = append(totalWeights, sum)
			sum = 0
			count = 0
		}
	}

	if count > 0 {
		totalWeights = append(totalWeights, sum)
	}

	average := 0.0
	for _, total := range totalWeights {
		average += total
	}
	average /= float64(len(totalWeights))

	for _, total := range totalWeights {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()
	fmt.Printf("%.2f\n", average)
}
