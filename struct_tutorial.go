package main


import "fmt"
import"time"

//A struck is a user define type to store a collection of different fields into a single field
//Struct decleration statement
//The 'type' keyword is introduce a new type , followed by the name 'Blog' and the keyword struct  
//The struct contains a collection of various fields inside the curly braces 
//Each field possesses a name and a type
//syntax:-type custom-name struct
type Blog struct {
	Title string
	Date time.Time
	Description string
	Author string
}

// Main function body 
func main (){
	
	//Initialize blog using struct
	//syntax:-var name custome-name
	var blogharsh Blog
	
	//Initlalize time statement
	var currenttime time.Time
	
	//Input statement for Title
	//Assess data in field using dot notation of the form structurname.fieldname
	fmt.Println("Enter the title")
	fmt.Scan(&blogharsh.Title)
	
	currenttime = time.Now()
	blogharsh.Date = currenttime
	
	//Input statement for Description
	fmt.Println("Enter the description")
	fmt.Scan(&blogharsh.Description)
	
	//Input statement for Author
	fmt.Println("Enter the author")
	fmt.Scan(&blogharsh.Author)
	
	//Print statement for Title
	fmt.Println(blogharsh.Title)
	
	//Print statement for Data
	fmt.Println(blogharsh.Date)
	
	//Print statement for Description
	fmt.Println(blogharsh.Description)
	
	//Print statement for Author
	fmt.Println(blogharsh.Author)
	
	
	
	// how to declare a custom type struct varible and assign value on at same time
	blogNeel := Blog{
			Title : "Some Blog Title",
			Date : time.Now(),
			Description : "some description",
			Author: "Neel",
			}
			
	//Print all term in separate line by using (\n)		
	fmt.Printf("%+v\n",blogNeel)	
			
	// how to declare a custom type struct varible and assign value on at same time		
	blogtom := Blog{
			Title : "Tom Story Book",
			Date : time.Now(),
			Description : "Tom kill jarry ",
			Author :"Tom the KING",
			}
			
	fmt.Printf("%+v\n",blogtom)
			
	// how to declare a custom type struct varible and assign value on at same time		
 	blogben10 := Blog{
			Title : "Ben10 Story Book",
			Date : time.Now(),
			Description : "Ben10 use fire wall ",
			Author :"Ben10",
			}
			
	fmt.Printf("%+v\n",blogben10)
	
	// how to declare a custom type struct varible and assign value on at same time
	blogsuperbox := Blog{
			Title : "Super Box Story Book",
			Date : time.Now(),
			Description : "the box hide in the room",
			Author :"some1",
			}
			
	fmt.Printf("%+v\n",blogsuperbox)
	
	
	
}
