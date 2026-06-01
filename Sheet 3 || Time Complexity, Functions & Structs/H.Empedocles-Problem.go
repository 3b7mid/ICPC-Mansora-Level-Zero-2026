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

	var n int
	fmt.Fscan(in, &n)
	
	a := make([]string, n)
	b := make([]string, n)
	c := make([]string, n)
	d := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		fmt.Fscan(in, &b[i])
		fmt.Fscan(in, &c[i])
		fmt.Fscan(in, &d[i])
	}
	

	
	good := true
	A := a[0]
	B := b[0]
	C := c[0]
	D := d[0]
	for i := 1; i < n; i++ {
		if a[i] !=  A || b[i] != B || c[i] != C || d[i] != D {
			good = false
			break
		}
 	}

	if good == true {
		fmt.Fprint(out, "Empedocles was right") 
	} else {
		fmt.Fprint(out, "you were right")
	}
}