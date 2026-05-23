package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	
	var prv, cur int64 
	fmt.Fscan(in, &prv)
	length, max_sub := 1, 1

	for i := 1; i < n; i++ {
		fmt.Fscan(in, &cur)

		if cur > prv {
			length++
		} else {
			length = 1
		}

		if length > max_sub {
			max_sub = length
		}

		prv = cur
	}
	fmt.Println(max_sub)
}
