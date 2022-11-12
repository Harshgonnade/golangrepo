package main

import "fmt"

func main(){

	// veriable decleration statement
	var a int
	
	// function calling statement 
	// input from user 
	fmt.Scan(&a)
	
	fmt.Println ("====================================================================================")
	
	// 1st type condition statement
	// syntax:- if <condition>{}
	if a>0 {
	fmt.Println("The value is greater than 0")
	}
	
	fmt.Println("====================================================================================")
	
	// 2nd type condition statement
	// syntax:- if <condition>{
	// 	}else{}
	if a>0 {
	fmt.Println("The value is greater than 0")
	}else{
	fmt.Println("The value is less than 0")
	}
	
	fmt.Println("====================================================================================")
	
	// 3rd type condition statement 
	//syntax:- if <condition>{
	//  }else if <condition>{
	// }else {}
	// next condition is applied in same line after first codition curly bracket end  
	if a>0 {
	fmt.Println("The number you have enter is positive")
	}else if a<0 {
	fmt.Println("The number you have enter is Negative")
	}else{
	fmt.Println("The number you have enter is Nutral/ Zero")
	}
	
	fmt.Println("====================================================================================")
}

