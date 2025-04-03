package main

import (
	"fmt"
	"reflect"
)

func main() {

	m := make(map[string]int)

	m["A"] = 1
	m["B"] = 2
	m["C"] = 3
	fmt.Println(m)

	v1 := m["A"]
	v2 := m["B"]
	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println("len:", len(m))
	delete(m, "B")
	fmt.Println("len:", len(m))
	fmt.Println(m)
	_, prs := m["d"]
	fmt.Println("prs:", prs)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	n2 := map[string]int{"foo": 1, "bar": 2}
	n2["bar"] = 3
	if reflect.DeepEqual(n, n2) {
		fmt.Println("n == n2")
	} else {
		fmt.Println("n != n2")
	}

}
