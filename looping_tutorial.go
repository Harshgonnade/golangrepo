package main 

import "fmt"

func main () {
	// veriable decleration statement
	// syntax:- var <name of variable><type of variable>
	// OR syntax:-<variable name>:=
	// a:=
	var a int 
	
	// function calling statement
	// input from user 
	// syntax:- fmt.Scan<&variable name>
	fmt.Scan(&a)
	
	// this loop is used for multiple operation
	// syntax for loop is:- for <initialization>;<condition>;<post>{}
	// like this 1) for i:=1 ; i<=20 ; i++{} 
	for i := 1 ; i <= 10 ; i++ {
		fmt.Println(a * i)
	}
	
}
