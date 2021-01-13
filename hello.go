package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello, World!")

	// var x int = 5 // verbose syntax
	// y := 7 // terse syntax for variable declaration and assignment, with inferred type
	// var sum int = x + y
	// fmt.Println(sum) // Println prints multiple passed elements separated by spaces

	// x := 6

	// if x > 6 {
	// 	fmt.Println ("More than 6")
	// } else if x == 6 {
	// 	fmt.Println ("Equal to 6")
	// } else {
	// 	fmt.Println ("Less than 6")
	// }

	// var a[5]int // declare an array of length 5, of integers. Go will automatically insert 5 0's
	// a[2] = 7 // assign third item to be 7

	// a := [5]int{5,4,3,2,1} // shorthand syntax, create an array of 5 integers,
	// // with the provided values, assign to variable a

	// arrays in Go have a fixed number of elements.
	// to have variable-length arrays, simply don't provide a length parameter, and
	// Go will create a "slice", which is an abstraction of the top of an array
	// a := []int{5, 4, 3, 2, 1}
	// a = append(a, 13) // the append function creates and returns a new slice
	// it can only be used on a slice and will be creating a new slice
	// (and underlying array) in the background

	// // to make an object (a hashmap), use the make function and pass it the types
	// vertices := make(map[string]int)
	// // in this case, make a map with string keys and int values and assign it to variable vertices

	// vertices["triangle"] = 2 // assignment is straightforward
	// vertices["square"] = 3

	// // to delete, pass the map reference/variable and the key value to the delete function
	// delete(vertices, "square")
	// fmt.Println(vertices)

	// // The only loop Go has is the for loop. Syntax is similar to JS
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// however if you only give the loop a condition (the second parameter)
	// and not a starting value or an iteration instruction, then you have a while loop
	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// you can loop through arrays by using the range operator
	// arr := []string{"a", "b", "c"}
	// for index, value := range arr {
	// 	fmt.Println("index: ", index, "value: ", value)
	// }

	// // you can do the same thing with a map, just get key instead of index
	// m := make(map[string]string)
	// m["a"] = "alpha"
	// m["b"] = "beta"
	// for key, value := range m {
	// 	fmt.Println("key: ", key, "value: ", value)
	// }

	// result, err := sqrt(-5)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	// // you create a struct similarly to creating an array
	// p := person{name: "Jake", age: 30}
	// // you can access struct properties with dot notation
	// fmt.Println(p.age)

	// pointers are shortcuts to memory addresses
	// you can get the memory address by prefixing the variable with &
	// you can store that address in a variable
	// then you access the value stored at that address by prefixing the
	// variable reference with *
	i := 7
	inc(&i)
	fmt.Println(i)
}

// demonstration of how to access (and then assign) the value at a memory address
func inc(x *int) {
	*x++
}

// go has something called structs, which appear to be something
// between a protocol buffer schema and an object class
// basically just a template for a type
// can't be inside a function
// you list field names and then their type separated by a space
// the type for a field can be another struct
type person struct {
	name string
	age  int
}

// basic functions in go look like this
func basic() {

}

// you take parameters and declare their types.
// if you want to return something, declare its type too
func sum(x int, y int) int {
	return x + y
}

// you can return multiple things
// for example a result and an error (which can be nil)
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
