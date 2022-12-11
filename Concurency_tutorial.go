package main
 import "fmt"
 //import "time"
 
 // Function Decleration statement for Even No
 func PrintEvenNo(upto int,done chan bool){
 		
 		//syntax of fro loop:-for <initialization>;<condition>;<post>{-----}
 		for i := 0 ; i <= upto ; i+=2 {
 		
 			// time.Sleep(time.Millisecond * 50 )
 		
 			fmt.Println("even: ",i)
 		}
 		
 		fmt.Println("sending true on print even chan")
 		//true value put in the done or sending signal to done 
 		done<-true
 		
 		fmt.Println("sending complete true on print even chan")
 }
 
 // Function Decleration satatement for Odd No
 func PrintOddNo(upto int, done chan bool){
 		
 		//syntax of fro loop:-for <initialization>;<condition>;<post>{-----}
 		for i := 1 ; i <= upto ; i+=2 {
 		
 		// time.Sleep(time.Millisecond * 100 )
 		
 			fmt.Println("Odd: ",i)
 		}
 		//true value put in the done or sending signal to done 
 		done<-true
 }
 
 Main function body
 func main() {
 
 	// Variable decleration statement
 	var doneEven chan bool
 
 	//assign the value of doneEven	
 	doneEven = make(chan bool)
 	
 	fmt.Println("execution started")
 	
 	//Print satatement for go routing 
 	go PrintEvenNo(100,doneEven)
 	
 	
 	
 	fmt.Println("-----------------------------------------------")

	//variable decleration & assign the value at a same time 
	doneOdd := make( chan bool)
	
	//Print statement for go-routing
	go PrintOddNo(1000,doneOdd)
		
		
	fmt.Println("waiting on even channel")
 	
 	//
 	<-doneEven
 	
 	fmt.Println("======================== even printing complete")
 	fmt.Println("-----------------------------------------------")
 	fmt.Println("waiting on odd channel")
 
 	//
 	<-doneOdd
 	
 	fmt.Println("======================== odd printing complete")
 	
 	fmt.Println("Excecution End") 
 	
 	
 }
