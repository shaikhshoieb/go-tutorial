package main

import (
	"fmt"
	"math"
)

type geo interface {
	area() float64
	perim() float64
}

type cir struct {
	radius float64
}

type rec struct {
	w, b float64
}

func (c cir) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c cir) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (r rec) area() float64 {
	return r.w * r.b
}
func (r rec) perim() float64 {
	return 2 * (r.w + r.b)
}

func measure(g geo) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCir(g geo) bool {
	_, ok := g.(cir)
	return ok
}

func detectRect(g geo) bool {
	_, ok := g.(rec)
	return ok
}
func main() {
	c := cir{radius: 5}
	r := rec{w: 10, b: 5}
	measure(c)
	measure(r)
	fmt.Println(r.perim())
	fmt.Println(c.perim())
	fmt.Println(r.area())
	fmt.Println(c.area())
	fmt.Println(r)
	fmt.Println(c)
	fmt.Println(detectCir(c))
	fmt.Println(detectCir(r))
	fmt.Println(detectRect(c))
	fmt.Println(detectRect(r))

}
