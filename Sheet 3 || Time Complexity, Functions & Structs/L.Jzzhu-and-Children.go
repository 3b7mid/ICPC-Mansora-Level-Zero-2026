package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		arr[i] = (arr[i] + m - 1) / m
	}
	res, mn := n, 0
	for i := 0; i < n; i++ {
		if arr[i] >= mn {
			mn = arr[i]
			res = i + 1
		}
	}
	fmt.Fprintln(out, res)
}
