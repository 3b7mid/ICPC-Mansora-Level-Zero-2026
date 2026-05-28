package main

import (
	"fmt"
	"math"
)

func main() {
	a := make([][]int, 5)
	iIdx, jIdx := 0, 0
	for i := 0; i < 5; i++ {
		a[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			fmt.Scan(&a[i][j])
			if a[i][j] == 1 {
				iIdx, jIdx = i, j
			}
		}
	}
	fmt.Println(math.Abs(float64(2-iIdx)) + math.Abs(float64(2-jIdx)))
}
