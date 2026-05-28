package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([][]int, n)
	prim_sum, secd_sum := 0, 0
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
			if i == j {
				prim_sum += a[i][j]
			}
			if i + j == n - 1 {
				secd_sum += a[i][j]
			}
		}
	}
	fmt.Println(math.Abs(float64(prim_sum) - float64(secd_sum)))
}
