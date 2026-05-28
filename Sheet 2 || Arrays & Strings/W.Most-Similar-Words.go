package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n, m int
		fmt.Scan(&n, &m)
		arr := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		min_cost := int(1e9)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				cost := 0
				for k := 0; k < m; k++ {
					cost += abs(int(arr[i][k]) - int(arr[j][k]))
				}
				min_cost = min(min_cost, cost)
			}
		}
		fmt.Println(min_cost)
	}
}
