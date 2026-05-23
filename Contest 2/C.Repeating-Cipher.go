package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var str string
	fmt.Scan(&str)
	add := 1
	res := ""
	for i := 0; i < n; i++ {
		res += string(str[i])
		i += add
		add++
	}
	fmt.Println(res)
}