package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const DELTA = 1e-15

func sqrt(x float64) float64 {
	z, d := float64(1), float64(1)

	for d > DELTA {
		z0 := z
		z = z - (z*z-x)/(2*z)
		d = math.Abs(z - z0)
	}
	return z
}

func main() {
	if len(os.Args) > 1 {
		arg, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			log.Fatal("Error:", err)
		}
		fmt.Println("The Sqrt of", arg, "is", sqrt(arg))
	} else {
		fmt.Println("Please provide a number")
	}
}
