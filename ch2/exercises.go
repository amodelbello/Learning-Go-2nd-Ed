package main

import "fmt"

// Disable electric-indent-mode and aggressive-indent-mode

func main() {
	i := 20
	var f = float64(i)
	fmt.Printf("i: %d, f:	%f\n", i, f)
	two()
	three()
}

func two() {
	const value = 4
	i := value
	var f float64 = value
	fmt.Printf("i: %d, f:	%f\n", i, f)
}

func three() {
	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Printf("b: %d, smallI:	%d, bigI: %d\n", b, smallI, bigI)
}
