package main

import "fmt"

type child struct {
	name string
}

type base struct {
	child
	num int
}

func (b base) details() string {
	return fmt.Sprintf("base num: %d", b.num)
}

func (c child) details() string {
	return fmt.Sprintf("child name: %s", c.name)
}

type derived struct {
	base
	str string
}

func main() {
	d := derived{
		base: base{num: 42, child: child{name: "child name"}},
		str:  "hello",
	}
	fmt.Println(d.details())
	fmt.Println(d.str)
	fmt.Println(d.base.details())
	fmt.Println(d.num)
	fmt.Println(d.child.details())

	type describer interface {
		details() string
	}
	var de describer = d
	fmt.Println(de.details())

}
