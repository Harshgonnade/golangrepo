package main

import "fmt"

// struct decleration statement
type Person struct{
	Name string
	Age int
}

//Function decleration statement 
// * for pointer declear
func(p *Person) SetName(newName string){
	p.Name = newName
}

//struct decleration statement
type Plannet struct{
		Plannet1 string
		Plannet2 string
}

//Function declertation statement 
func(p1 *Plannet) SetPlannetName(changeName string){
	p1.Plannet1 = changeName
}

//Main function Body 
func main(){

	//variable decleration & assign the value at same time 
	per := Person{
		Name:"Haarsh",
		Age: 23,
	}

	//Print statement before setName
	fmt.Println(per)

	//Calling setName function 
	(&per).SetName("Harsh")

	//Print statement after setName
	fmt.Println(per)

	fmt.Println("----------------------------------------")

	//variable decleration & assign the value at same time 
	universe := &Plannet{
		Plannet1:"Earths",
		Plannet2:"Moon",
	}

	//Print statement before setPlannetName
	fmt.Println(universe)

	//Calling setPlannetName function
	universe.SetPlannetName("Earth")

	//Print statement after setPlannetName
	fmt.Println(universe)
}
