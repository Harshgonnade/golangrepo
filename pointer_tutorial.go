package main

import "fmt"

//Decleration statement
type Person struct {
	name string
	age  int
}

//function decleration 
func setAge(p1 *Person) {
	if (*p1).name == "neel" {
	
		// this line is for cheaking 
		fmt.Println("if block is run")
	
		(*p1).age = 26
	} else {
		(*p1).age = 99
	}
}

Main function body
func main() {

	// initialization state
	//pointer statement 
	var p *Person
	
	//creat new instent
	p = &Person{}
	
	//input statement
	fmt.Scan(&p.name)

	//print statement before calling setage function 
	fmt.Println("person instace/Example before setting age: ", p)
	
	//calling setage function
	setAge(p)

	//print statement after calling setage function
	fmt.Println("person instance/Ehample after setting age: ", p)

}
//Pointers in Go programming language or Golang is a variable that is used to store the memory address of another variable.
// Pointers in Golang is also termed as the special variables. 
//The variables are used to store some data at a particular memory address in the system.
