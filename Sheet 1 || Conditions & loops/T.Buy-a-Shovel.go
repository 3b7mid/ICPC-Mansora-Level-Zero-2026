package main

import "fmt"

func main() {
	var k, r int
	fmt.Scan(&k, &r)
	cnt := 0 
	kk := k
	for true {
		cnt++
		if k%10 == 0 || k%10 == r {
			break
		}
		k += kk
	}
	fmt.Println(cnt)
}
