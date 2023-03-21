package main

import "fmt"

func main() {
	// x := 1
	// p := &x
	// fmt.Println(*p)
	// *p = 2
	// fmt.Println(x)

	// var x, y int
	// fmt.Println(&x == &x, &x == &y, &x == nil)

	// var p = f()
	// fmt.Println("------")
	// fmt.Println(p)
	// fmt.Println(f())
	// fmt.Println(f() == f())
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}

func f() *int {
	v := 1
	fmt.Println(v)
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
