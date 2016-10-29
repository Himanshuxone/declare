package main
import "fmt"

// create a function using pointers
// passing pointers to the functions

func main() {
	
	// y := new(int) is logically equal to var i int; y = func() *int{ return &i } 
	// below new(int retunrs the memory address of y
	// can generate a pointer using new key word
	yPtr := new(int)

	// to pass the pointer to the use "&" keyowrd 
	changeXVAL(yPtr)
	// print the changed value of x
	fmt.Println("y =",*yPtr)

}

// passing a pointer to the function
// to pass teh type of the pointer use "*" keyowrd before return type
func changeXVAL(yPtr *int){
	// to get the value of a pointer use "*" keyword and to get the address of a pointer use "&" keyword
	*yPtr = 100

	// above teh value of x will change since we are changing the value on the memory address that x points to

}