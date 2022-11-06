package main

import "fmt"

// program for addition
// function  declearation line for addition
func add (a int,b int)(int){
	var sum int 
	
	sum = a + b
	
	return sum
}

// program for substraction
// function  declearation line for substraction
func sub (a int,b int)(int){
	var sub int 
	
	sub = a - b
	
	return sub
}

// program for multiplication
// function  declearation line for multiplication
func mul (a int,b int)(int){
	var mul int 
	
	mul = a * b
	
	return mul
}

// program for diivision
// function  declearation line for division
func div (a int,b int)(float64){
	var div float64
	
	div = float64(a) / float64(b)
	
	return div
}

///================================================== main func ====================================================///

func main () {
    // decleration variable num1 & num2
    var num1 int
    var num2 int
    
    // take input from user
    // syntax: fmt.scan(&variable name)
    fmt.Println("enter the value of num1")
    fmt.Scan(&num1)
    fmt.Println("enter the value of num2")
    fmt.Scan(&num2)
    
    // variable decleration line for addition 
    var additionresult int
    
    additionresult = add(num1 , num2)
    
    fmt.Println("Addition is ",additionresult)
    
    //=====
    
     // variable decleration line for substraction
     var subresult int
    
    subresult = sub(num1 , num2)
    
    fmt.Println("Substraction is ",subresult)
    
    //========
    
    // variable decleration line for multiplication
    var mulresult int 
    
    mulresult = mul(num1,num2)
    
    fmt.Println("Multiply is ",mulresult)
    
    //=====
    
    // variable decleration line for division
    var divresult float64
    
    divresult = div(num1,num2)
    
    fmt.Println("Division is ",divresult)
}
