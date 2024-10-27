package main

import (
	"fmt"
	"math"
)

func FindMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		panic("Array cannot be empty")
	}

	min, max := math.MaxInt, math.MinInt
	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func main() {
	arr := []int{5, 3, 9, 1, 6, -2, -15, 100, 10, 32, 25, 52, 55, 35, 1000, 0 - 100}
	min, max := FindMinMax(arr)
	fmt.Printf("Minimum value: %d\n", min)
	fmt.Printf("Maximum value: %d\n", max)
}
