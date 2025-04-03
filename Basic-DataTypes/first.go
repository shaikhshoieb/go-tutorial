package main

import (
	"fmt"
)

/*
Multiline comment
Here is the second one


~~Basic
1.Need to have package main
2. Packages can be imported using import("package")
3. fmt is used for basic io operation
4. main function is the one executed
5. go run file name runs the code

~~Variables
1. While defining a variable you eitehr need to assign a value or give the type of both
2.If assigning a value  := can be used so that the type is directly infrred by the compiler
3. var variablename type = value is the basic syntax
4. There are 4 basic type of datatypes
	a. string
	b. int
	c. float32
	d. bool
5. := can only be run inside a function where as var can be used anywhere
6. Multiple variables can be decared in comma separated mannaeer
7. Variables can be defined in block using var key word and parenthesis
8 . var (
		m int     = 1
		n string  = "String"
		o bool    = true
		p float64 = 10.2743
	)




~~NAMING CONVENTION
1. A variable name must start with a letter or an underscore character (_)
2. A variable name cannot start with a digit
3. A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
4. Variable names are case-sensitive (age, Age and AGE are three different variables)
5. There is no limit on the length of the variable name
6. A variable name cannot contain spaces
7. The variable name cannot be any Go keywords


~~Coonstants
1. If a value needs to be fixed we can use the const keyword
2.Two types of constants
	a. Typed constants :- That are defined with type
	b. untyped constants :-  That are defined without type
3. Constants can also be defined in a block

~~ Output statments
1. fmt.Print(i,j) :- Adds no whitespaces between i and j
2. fmt.Println(i,j) :- Adds a whitespace between the two variables
3. fmt.Printf( "Value : %v and tyep :%T\n"i,i) :- uses format to print variables where %v defined the value and %T defined the type
4. Types of formatting
	a. %v :- Prints value
	b. %#v :- Print the value in Go-syntax format
	c. %T :- Print the type of value
	d. %% :- Prints % sign
5. For more formatting, see https://www.w3schools.com/go/go_formatting_verbs.php
6.
*/

func main() {
	// Comment
	var test = "acbd"
	var test2 string = "string"
	var test3 int32
	test3 = 1
	var test4 bool = false // type is bool
	var test5 float64 = 10.02
	test6 := true
	var student1 string = "John"    //type is string
	var student2 = "Jane"           //type is inferred
	x := 2                          //type is inferred
	var a, b, c, d int = 1, 3, 5, 7 // type and vallues is defined
	var e, f, g, h int              // type is defined
	i, j, k, l := 11, 12, 14, 14    // type is inferred value is assigned direclty
	var (
		m = 1
		n = "String"
		o = true
		p = 10.2743
	)
	const test7 = 8
	const test9 string = "String"
	const (
		q        = 1
		r        = true
		s string = "Test"
	)

	t, u := "Hello", "World"

	var i1 = 15.5
	var txt = "Hello World!"

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println("Hey there")
	fmt.Println(test)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println(test4)
	fmt.Println(test5)
	fmt.Println(test6)
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g, h)
	fmt.Println(i, j, k, l)
	fmt.Println(m, n, o, p)
	fmt.Println(test7)
	fmt.Println(test9)
	fmt.Println(q, r, s)
	fmt.Print(t, u)
	fmt.Printf("%v\n", i1)
	fmt.Printf("%#v\n", i1)
	fmt.Printf("%v%%\n", i1)
	fmt.Printf("%T\n", i1)
	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	// fmt.Println("This line does not execute")
}
