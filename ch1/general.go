package main

import (
	"fmt"
	"os"
)

func main() {
	// resp := fib(7)
	// fmt.Println(resp)
	var x int = 1
	fmt.Println(x.(int))
}

func gcd(x, y int) (int, int) {
	for y != 0 {
		x, y = y, x%y
	}
	return x, y
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println("----")
		fmt.Fprintf(os.Stdout, "x:%v--", x)
		fmt.Fprintf(os.Stdout, "y:%v\n", y)
		fmt.Println("----")
		x, y = y, x+y
	}
	return x
}
