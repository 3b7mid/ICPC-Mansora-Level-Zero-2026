package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 5003

var isPal [N][N]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var str string
	fmt.Fscan(in, &str)

	n := len(str)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			isPal[i][j] = false
		}
	}
	for i := 0; i < n; i++ {
		isPal[i][i] = true
	}
	for i := 0; i < n-1; i++ {
		if str[i] == str[i+1] {
			isPal[i][i+1] = true
		}
	}
	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && str[l] == str[r] {
			isPal[l][r] = true
			l--
			r++
		}
	}
	for i := 0; i < n-1; i++ {
		l, r := i, i+1
		for l >= 0 && r < n && str[l] == str[r] {
			isPal[l][r] = true
			l--
			r++
		}
	}

	var q int
	fmt.Fscan(in, &q)

	for q > 0 {
		var l, r int
		fmt.Fscan(in, &l, &r)
		l--
		r--
		if isPal[l][r] {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
		q--
	}
}