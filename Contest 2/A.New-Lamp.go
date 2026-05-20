package main

import "fmt"

func main() {
	var i, v, p int64
	fmt.Scan(&i, &v, &p)
	if i*v <= p {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}