package main

import "fmt"
import"time"

type Blog struct {
	Title string
	Date time.Time
	Description string
	Author string
}


func main (){
	
	//
	var blogharsh Blog
	
	var currenttime time.Time
	
	
	fmt.Scan(&blogharsh.Title)
	
	currenttime = time.Now()
	
	blogharsh.Date = currenttime
	
	fmt.Scan(&blogharsh.Description)
	
	fmt.Scan(&blogharsh.Author)
	
	//
	fmt.Println(blogharsh.Title)
	
	//
	fmt.Println(blogharsh.Date)
	
	//
	fmt.Println(blogharsh.Description)
	
	//
	fmt.Println(blogharsh.Author)
	
	
	
	// how to declare a custom type struct varible and assign value on at same time
	
	blogNeel := Blog{
			Title : "Some Blog Title",
			Date : time.Now(),
			Description : "some description",
			Author: "Neel",
			}
			
			
	fmt.Printf("%+v\n",blogNeel)	
			
			
	blogtom := Blog{
			Title : "Tom Story Book",
			Date : time.Now(),
			Description : "Tom kill jarry ",
			Author :"Tom the KING",
			}
			
	fmt.Printf("%+v\n",blogtom)
			
			
 	blogben10 := Blog{
			Title : "Ben10 Story Book",
			Date : time.Now(),
			Description : "Ben10 use fire wall ",
			Author :"Ben10",
			}
			
	fmt.Printf("%+v\n",blogben10)
	
	blogsuperbox := Blog{
			Title : "Super Box Story Book",
			Date : time.Now(),
			Description : "the box hide in the room",
			Author :"some1",
			}
			
	fmt.Printf("%+v\n",blogsuperbox)
	
	
	
}
