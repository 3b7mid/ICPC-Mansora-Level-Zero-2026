package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	left := a[:n/2]
	right := a[n/2:]
	for _, v := range left {
		fmt.Print(v, " ")
	}
	for i := len(left) - 1; i >= 0; i-- {
		fmt.Print(left[i], " ")
	}
	fmt.Print("\n")
	for i := len(right) - 1; i >= 0; i-- {
		fmt.Print(right[i], " ")
	}
	for _, v := range right {
		fmt.Print(v, " ")
	}
}