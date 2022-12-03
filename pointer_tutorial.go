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
		fmt.Println("if block is run")
		(*p1).age = 26
	} else {
		(*p1).age = 99
	}
}

func main() {

	// initialization state
	//pointer statement 
	var p *Person
	
	//creat new instent
	p = &Person{}
	
	//input statement
	fmt.Scan(&p.name)

	//print statement before calling setage function 
	fmt.Println("person instace before setting age: ", p)
	
	//calling setage function
	setAge(p)

	//print statement after calling setage function
	fmt.Println("person instance after setting age: ", p)

}
