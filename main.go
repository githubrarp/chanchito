package main

import (
	"fmt"
	"math/rand"
)

func main() {

	amount := 0.0
	posibleValues := [3]float64{0.05, 0.10, 0.25}

	for amount < 20 {
		amount += posibleValues[rand.Intn(4-1)]
		fmt.Printf("%5.2f\n", amount)
	}

}
