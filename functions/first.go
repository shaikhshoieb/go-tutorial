package main

import (
	"fmt"
)

func test(a string, b string) {
	fmt.Println(a, b)
}

func add(a, b, c int) int {
	return a + b + c
}

func returnMultipleValues() (int, string) {
	return 1, "hello"
}

func variadicFunc(a string, b ...int) {
	fmt.Println(a)
	total := 0
	for _, v := range b {
		total += v
	}
	fmt.Println("total:", total)
}

// The variadic function can take zero or more arguments of the same type.

func intVariadicFunc(b ...int) int {
	fmt.Println("int variadic function")
	fmt.Println(b)
	total := 0
	for _, v := range b {
		total += v
	}
	fmt.Printf("total: %d\n", total)
	return total
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	test("hello", "world")
	res := add(1, 2, 3)
	fmt.Println(res)
	a, b := returnMultipleValues()
	fmt.Println(a, b)
	_, c := returnMultipleValues()
	fmt.Println(c)
	variadicFunc("hello", 1, 2, 3, 4, 5)
	variadicFunc("hello")
	variadicFunc("hello", 1, 2, 3, 4, 5, 6, 7, 8, 9)
	intVariadicFunc(1, 2)
	intVariadicFunc(1, 2, 3, 4, 5)
	nums := []int{1, 2, 3, 4, 5}
	intVariadicFunc(nums...)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(fact(7))

	var fibo func(n int) int

	fibo = func(n int) int {
		if n < 2 {
			return n
		}
		return fibo(n-1) + fibo(n-2)
	}
	// fib = func(n int) int {
	// 	if n < 2 {
	// 		return n
	// 	}
	// 	return fib(n-1) + fib(n-2)
	// }
	// fib = func(i int) int {
	// 	if i < 2 {
	// 		return i
	// 	}
	// 	return fib(i-1) + fib(i-2)

	// }

	fmt.Println(fibo(7))
	numsSum := []int{1, 2, 3, 4, 5}
	sumTotal := 0
	for _, v := range numsSum {
		sumTotal += v
	}
	fmt.Println("sum total:", sumTotal)

	newMap := map[string]string{"a": "a1", "b": "b1", "c": "c1"}
	for k, v := range newMap {
		fmt.Printf("%s : %s \n", k, v)
	}

	for i, k := range "go" {
		fmt.Println(i, k)
	}

}
