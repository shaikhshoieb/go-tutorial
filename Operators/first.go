package main

import (
	"fmt"
)

/*
~~Operators
1.There are different types of operators
	a. Arithematic
		i. + :- Adition
		ii. - :- subtraction
		iii. * :- multiplication
		iv. / :- division
		v. % :- modulus
		vi. ++ :- increment
		vii. -- :decrement
	b. Assignment
		i. = :- a=10
		ii += :- addition assignment a+=10
		iii. Every other operator can be used in th similar manner
	c. Comparison
		i. > :- greater than
		ii. < :- less than
		ii. And so on like the other languages
	d. Logical
		i. && :- Logical and
		ii. || logical or
		iii. ! :- logical not
	e. Bitwise :- refer https://www.w3schools.com/go/go_bitwise_operators.php

*/

func main() {
	fmt.Print("Starting \n")
	var a = 10 + 15

	var (
		b = 15 + 20
		c = b + 10
		d = a + b + c + 10
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
	a += 10
	fmt.Println(a)

	// fmt.Println(a++)
	// fmt.Println(a--)

}
