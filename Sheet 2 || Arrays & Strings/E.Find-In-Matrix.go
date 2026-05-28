package main

import "fmt"

func main() {
	var n, m, x int
	fmt.Scan(&n, &m)
	a := make([]int, n*m)
	for i := 0; i < n*m; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&x)
	for _, v := range a {
		if x == v {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
