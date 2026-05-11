package main

import "fmt"

func main()  {
	var n int 
	fmt.Scan(&n)
	if n % 2 == 0 {
		fmt.Println("Ahmed welcomes you to ICPC Mansoura Community")
	} else {
		fmt.Println("Amr welcomes you to ICPC Mansoura Community")
	}
}