package main

import (
	"fmt"
)

/*
~~Sices

~~Declaration
1. Slices are similar to arrays, but are more powerful and flexible.
2. Like arrays, slices are also used to store multiple values of the same type in a single variable.
3. However, unlike arrays, the length of a slice can grow and shrink as you see fit.
4. In Go, there are several ways to create a slice:
	a. Using the []datatype{values} format
	b. Create a slice from an array
	c. Using the make() function
5. There are two important function related to length of a slice
	a. len() :- Shows the number of value in slice
	b. cap() :- Shows the capacity of slice
6. Slices can be created froman array using the syntac var mySlice = array[start:end]
	a. Start and end is the index
	b. Capacity of slice is defined as number of array from start to the end of the orifinal array
7. Slice can be created using the make function
	a. slice_name := make([]type, length, capacity)
	b. Capacity is variable and ifnot defined length is taken as the capacity

~~Accessing the element
1. Here the elements can be accessed using the indexes
2. Elements can be modified using the method same as array


~~Appending elements to slice
1. We can append elements to slice
	a. slice_name = append(slice_name, element1, element2, ...)
2. The size of slice doubles as the we append the slice using append function
3. We can also combine the different slices using the append function by passing the slices
4. 2 Slices can be appnded and by adding ... operator after second slice
5. Capacity  of slice doubles the size of largest slice


~~Copy function
1. When using slices, Go loads all the underlying elements into the memory.
2. If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
3. The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.
4. Syntax :- copy(dest, src)
*/

func main() {
	fmt.Println("Starting")

	slice := []int{1, 2, 3, 4}
	var slice2 []int
	myslice1 := []int{}

	fmt.Println(slice)
	// slice2[0] = 5
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	fmt.Println(len(slice2))
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)
	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]
	myslice3 := arr1[1:3]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
	fmt.Printf("myslice = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	//Creating with capacity
	myslice4 := make([]int, 5, 10)
	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	// without omitted capacity
	myslice5 := make([]int, 5)
	fmt.Printf("myslice5 = %v\n", myslice5)
	fmt.Printf("length = %d\n", len(myslice5))
	fmt.Printf("capacity = %d\n", cap(myslice5))

	//Appending
	myslice6 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice6 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n", cap(myslice6))
	myslice6 = append(myslice6, 20, 21)
	fmt.Printf("myslice6 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n", cap(myslice6))

	//Merging two slices
	myslice7 := []int{1, 2, 3, 5}
	myslice8 := []int{4, 5, 6}
	myslcie9 := append(myslice7, myslice8...)

	fmt.Printf("myslcie9=%v\n", myslcie9)
	fmt.Printf("length=%d\n", len(myslcie9))
	fmt.Printf("capacity=%d\n", cap(myslcie9))

	//Copy function
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]     //Creating array leaving last 10 number
	numbersCopy := make([]int, len(neededNumbers)) //creating an slice with length of above created array and no capacity
	copy(numbersCopy, neededNumbers)               //Copy neededNumbers to numbersCopy slice

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
