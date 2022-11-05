package main

import "fmt"
// Here we are call the int value and print return by using the func main
func main() {
	var result int

	result = add(3, 5)

	fmt.Println(result)
}

// function is a reusable 'code block' which is having same name
// function take some 'inpute arguments' and return some 'output arguments'
// syntax: func <function name> (inpute arguments) (output arguments){code block}
func add(a int, b int) int {

	var sum int

	sum = a + b

	return sum
}
