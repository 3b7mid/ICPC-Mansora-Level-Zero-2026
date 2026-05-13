package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var l, r, x int
	for q > 0 {
		fmt.Scan(&l, &r, &x)
		cnt := 0
		for i := l - 1; i < r; i++ {
			if x == a[i] {
				cnt++
			}
		}
		fmt.Println(cnt)
		q--
	}
}
