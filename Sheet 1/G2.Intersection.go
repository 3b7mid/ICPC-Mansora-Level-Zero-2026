package main

import "fmt"

func main() {
	var l1, l2, r1, r2 int 
	fmt.Scan(&l1, &r1, &l2, &r2)
	if l2 > r1{
		fmt.Println(-1)
	}else {
		fmt.Println(max(l1, l2), min(r1, r2))
	}

}