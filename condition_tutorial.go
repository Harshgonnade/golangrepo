package main

import "fmt"

func main(){

	// variable decleration statement
	var a int
	
	// function calling statement 
	// input from user 
	fmt.Scan(&a)
	
	fmt.Println ("====================================================================================")
	
	// 1st type condition statement
	//When you have only one condition then use type 1 condition 
	// syntax:- if <condition>{}
	//ex:-
	if a>0 {
	fmt.Println("The value is greater than 0")
	}
	
	fmt.Println("====================================================================================")
	
	// 2nd type condition statement
	//When you have two condition then use typr 2 condition
	// syntax:- if <condition>{
	//          }else{}
	//ex:-
	if a>0 {
	fmt.Println("The value is greater than 0")
	}else{
	fmt.Println("The value is less than 0")
	}
	
	fmt.Println("====================================================================================")
	
	// 3rd type condition statement 
	//When you have more than two condition then use type 3 conition
	//syntax:- if <condition>{
	//  }else if <condition>{
	// }else {}
	// next condition is applied in same line after first codition curly bracket end  
	//ex:-
	if a>0 {
	fmt.Println("The number you have enter is positive")
	}else if a<0 {
	fmt.Println("The number you have enter is Negative")
	}else{
	fmt.Println("The number you have enter is Nutral/ Zero")
	}
	
	fmt.Println("====================================================================================")
}

