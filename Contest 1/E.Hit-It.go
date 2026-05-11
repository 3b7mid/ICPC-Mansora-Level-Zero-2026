package main

import (
	"fmt"
	"math"
)

func main() {
	var xr, yr, xw, yw float64
	fmt.Scan(&xr, &yr, &xw, &yw)
	rdist := math.Sqrt(math.Pow(0-xr, 2) + math.Pow(0-yr, 2))
	wdist := math.Sqrt(math.Pow(0-xw, 2) + math.Pow(0-yw, 2))
	if rdist < wdist {
		fmt.Println("Russo")
	} else if wdist < rdist {
		fmt.Println("Wil")
	} else {
		fmt.Println("Empate")
	}
}
