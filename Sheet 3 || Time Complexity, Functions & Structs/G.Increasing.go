package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func increasing(arr []int, sz int) bool {
	if sz == 1 {
		return true
	}
	sort.Ints(arr)
	for i := 1; i < sz; i++ {
		if arr[i] == arr[i - 1] {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int 
	fmt.Fscan(in, &t) 
	for ; t > 0; t-- {
		var n int 
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		if increasing(arr, n) == true {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
