package main

import (
	"fmt"
)

func main() {

	// call function1
	function1()

	// call greet
	greet("Bart")

	// call multiply
	m := multiply(2, 4)
	fmt.Printf("multiply returned %d\n", m)

	// call add
	a := add()
	fmt.Printf("add returned %d\n", a)

	// call findFruit
	myFruits := []string{"apple", "banana", "lemon"}
	key, value, err := findFruit("lemon", myFruits)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Fruit: %s found as %d key on the list.\n", value, key)
	}

	// anonymous function
	myVar := func() {
		fmt.Println("Hello from anon function")
	}
	myVar()
	fmt.Printf("Type of myVar is: %T", myVar)

}

// function1 takes no parameters and has no return values
// prints directly to std.output
func function1() {
	fmt.Println("Hello from function1")
}

// greet takes one parameter of type string, has no return
// prints directly to std.output
func greet(name string) {
	fmt.Printf("Hello %s.\n", name)
}

// multiply takes two parameters of type integer and returns interger
func multiply(num1 int, num2 int) int {
	result := num1 * num2
	return result
}

// add takes any number of integer type parameter and returns integer
func add(num ...int) int {
	var result int
	for _, i := range num {
		result += i
	}
	return result
}

// findFruit takes two params: fruit string that we are looking for
// and array or fruit strings to search in
// the return is the array key where the fruit is, the value at that key
// and error if the fruit was not found
func findFruit(fruit string, fruits []string) (key int, value string, err error) {
	for i, f := range fruits {
		if fruit == f {
			return i, f, nil
		}
	}
	return 0, "", fmt.Errorf("Fruit: %s not in the list", fruit)

}
