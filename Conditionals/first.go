package main

import (
	"fmt"
)

/*
~~Conditional statemnts

~~If statemnt
Syntax if x>y
{
	do something
}


~~If else
if condition {
	do somthing
}else{
	do something
}

~~If else if else

if condition1 {
   // code to be executed if condition1 is true
} else if condition2 {
   // code to be executed if condition1 is false and condition2 is true
} else {
   // code to be executed if condition1 and condition2 are both false
}

~~Nested if condition

if condition1 {
   // code to be executed if condition1 is true
  if condition2 {
     // code to be executed if both condition1 and condition2 are true
  }
}



*/

func main() {
	fmt.Println("Starting")
	x := 30
	y := 20
	if x > y {
		fmt.Println("X greater than y")
	}

	if x < y {
		fmt.Println("X lesser than y")
	} else {
		fmt.Println("X greater than y")
	}

	if x < y {
		fmt.Println("X lesser than y")

	} else if x == y {
		fmt.Println("X equal to y")
	} else {
		fmt.Println("X greater than y")
	}

}
