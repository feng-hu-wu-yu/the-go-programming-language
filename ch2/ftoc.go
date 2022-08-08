package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC \n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC \n", boilingF, fToC(boilingF))

	// test
	x := 1
	p := &x
	fmt.Printf("x: %d, p: %v \n", x, p)
	*p = 2
	fmt.Printf("updated x: %d, updated p: %v ", x, p)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
