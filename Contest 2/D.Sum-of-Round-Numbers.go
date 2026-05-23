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
	var t int 
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n int 
		fmt.Fscan(in, &n)
		arr := make([]int, 5)
		ten, sz := 1, 0
		for n > 0 {
			dig := n % 10;
			if dig > 0 {
				arr[sz] = dig * ten
				sz++
			}
			n /= 10
			ten *= 10
		}
		fmt.Println(sz)
		for i := 0; i < sz; i++ {
			fmt.Print(arr[i], " ")
		}
		fmt.Printf("\n")
	}
	
}
