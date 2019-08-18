package main

import (
	"fmt"
	"time"
)

var s [51]int // for 50 fibonacci numbers

func fib(n int) int {
	if n < 3 {
		return 1
	}
	if s[n] != 0 {
		return s[n]
	}
	s[n] = fib(n-1) + fib(n-2)
	return s[n]
}

func main() {
	var i int
	t0 := time.Now()
	for i = 1; i <= 50; i++ {
		fmt.Printf("Fibonnaci number No. %d = %d\n", i, fib(i))
	}
	println()
	println("Computing time =", time.Since(t0).Seconds())
}
