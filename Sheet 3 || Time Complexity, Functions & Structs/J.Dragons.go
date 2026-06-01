package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Dragon struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s, n int
	fmt.Fscan(in, &s, &n)

	arr := make([]Dragon, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i].x, &arr[i].y)
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].x == arr[j].x {
			return arr[i].y > arr[j].y
		}
		return arr[i].x < arr[j].x
	})

	for _, v := range arr {
		if v.x >= s {
			fmt.Fprintln(out, "NO")
			return
		} else {
			s += v.y
		}
	}
	fmt.Fprintln(out, "YES")
}
