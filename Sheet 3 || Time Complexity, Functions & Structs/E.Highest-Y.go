package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	arr := make([]Point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i].x, &arr[i].y)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].y > arr[j].y
	})

	for _, v := range arr {
		fmt.Fprintln(out, v.x, v.y)
	}
}
