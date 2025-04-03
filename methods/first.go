package main

import "fmt"

type rect struct {
	w, b int
}

func (r *rect) area() int {
	return r.w * r.b
}

func (r rect) perim() int {
	return 2 * (r.w + r.b)
}

func main() {
	r := rect{w: 10, b: 5}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())

	r.b = 20
	fmt.Println(r.area())
	fmt.Println(r.perim())
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
