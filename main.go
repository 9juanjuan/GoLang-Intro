package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
}

/////////// Declaring variables with var

// func main() {
// 	var x int = 5
// 	var y int = 7
// 	var sum int = x + y
// 	fmt.Println(sum)
// }

////////// Declaring variables with :=

// func main() {
// 	x := 5
// 	y := 7
// 	sum := x + y
// 	fmt.Println(sum)

// }

/////// Else / Else  if

// func main() {
// 	x := 5

// 	if x > 6 {
// 		fmt.Println("More than 6")
// 	} else if x < 6 {
// 		fmt.Println("Less than 6")
// 	} else {
// 		fmt.Println("Well....I guess you're not a number?")
// 	}
// }

//////// Arrays

// func main() {
// 	// a := [5]int{0, 1, 2, 3, 4}
// 	// a[2] = 7
// 	//
// 	//Slice of Integers
// 	a := []int{0, 1, 2, 3, 4}
// 	// built in append function does not modify but return a new slice
// 	a = append(a, 13)

// 	fmt.Println(a)
// }

/////// Mapping
// func main() {
// 	vertices := make(map[string]int)

// 	vertices["triangle"] = 2
// 	vertices["square"] = 3
// 	vertices["octagon"] = 8

// 	delete(vertices, "square")
// 	fmt.Println(vertices)
// }

///// Loops
// func main() {
// For loop
// for i := 0; i < 5; i++ {
// 	fmt.Println(i)
// }

/// While loop
// i := 0
// for i < 5 {
// 	fmt.Println(i)
// 	i++
// }

// Looping through array
// arr := []string{"a", "b", "c"}
// for index, value := range arr {
// 	fmt.Println("index", index, "value", value)
// }

// Loop through map, similar but with key instead of index
// m := make(map[string]string)
// m["a"] = "alpha"
// m["b"] = "beta"

// for key, value := range m {
// 	fmt.Println("key:", key, "value:", value)
// }
// }

// func main() {
// 	result := sum(2, 3)
// 	fmt.Println(result)
// }

// func sum(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	result, err := sqrt(16)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}
// }

// //Error handling and return multiple values
// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("Undefined for negative numbers")
// 	}
// 	return math.Sqrt(x), nil

// }

// type person struct {
// 	name string
// 	age  int
// }

// func main() {
// 	p := person{name: "Jake", age: 23}
// 	fmt.Println(p.age)
// }

//// Pointers

// func main() {
// 	i := 7
// 	inc(&i)
// 	fmt.Println(i)
// }

// func inc(x *int) {
// 	*x++
// }
