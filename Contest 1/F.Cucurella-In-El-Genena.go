package main

import "fmt"

func main() {
	var l, r int64
	fmt.Scan(&l, &r)
	sum_1_to_r := r * (r + 1) / 2
	l--
	sum_1_to_l_mins_1 := l * (l + 1) / 2
	fmt.Println(sum_1_to_r - sum_1_to_l_mins_1)
}
