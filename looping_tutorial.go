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
	fmt.Print("Enter number:-	");
	fmt.Scan(&a)
	
	// this loop is used for multiple operation
	// syntax for loop is:- for <initialization>;<condition>;<post>{}
	// like this 1) for i:=1 ; i<=20 ; i++{} 
	for i := 1 ; i <= 10 ; i++ {
		fmt.Println(a,"x",i,"=",a * i)
	}
	
	//
	fmt.Println("Integer Array");
	numbers := []int{7, 9, 1, 2, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i],i)

	}
	
	//
	fmt.Println("String Array");
	arr := []string{"a", "b", "c", "d", "e", "f"}

	for index, a := range arr {
		fmt.Println(index, a)
	}
	
}
