package main

import (
	"fmt"
)

/*
~~Arrays

~~Declaration
1.In go there are two types in which an array can be created
	a. var keyword
		i. With length and values :- var array_name = [length]datatype{values} here length is defined
		ii. Without length and values:-  var array_name = [...]datatype{values}  here length is inferred
	b. := operator
		i. With length and values :- array_name := [length]datatype{values} // here length is defined
		ii. Without length and values :- array_name := [...]datatype{values} // here length is inferred


~~Accessing element
1. Elements can be accessed using the index operator
2. Indexing in go starts with 0
3. Value can be chages by referring to an index number
4. While defining an array one can
	a. Fully inititalise
	b. Partiallry initialise
	c. Not intialise
5. If and element on the array in not initialised it takes the dataypes default value
	a. int 0
	b. string ""
	c. bool false
	d. float64 0
6. One can also initialise a specific element using the index as the key while defining it
7. Use len(array_name) to find length of an array

*/

func main() {
	fmt.Println("Starting")
	//Array with var keyword
	var arr1 = [3]int{1, 2, 3}
	var arr2 = []int{4, 5, 6}

	//Array with := operator
	arr3 := [4]string{"a", "a", "a", "a"}
	arr4 := []string{"a", "a", "a", "a", "b", "c"}
	arr5 := [5]string{}
	arr6 := [5]int{1, 2, 3}
	arr7 := [6]bool{true}
	arr8 := [6]float64{1.1}
	arr9 := [6]string{0: "a", 4: "b"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr1[0])
	fmt.Println(arr2[1])
	fmt.Println(arr3[2])
	fmt.Println(arr4[5])
	arr1[1] = 4
	fmt.Println(arr1)
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)
	fmt.Println(arr8)
	fmt.Println(arr9)
	fmt.Println(len(arr9))

}
