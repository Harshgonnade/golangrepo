// packge decleration line
package main
// import declaration line
import "fmt"
// Here we are call the int value and print return by using the func main
func main() {
	//Variable decleration statement
	var result int
	
	//Return statement
	result = add(3, 5)

	fmt.Println(result)
}

// function decleration line
// function is a reusable 'code block' which is having some name
// function take some 'inpute arguments' and return some 'output arguments'
// syntax: func <function name> (inpute arguments) (output arguments){code block}
func add(a int, b int) int {

	var sum int

	sum = a + b

	return sum
}


