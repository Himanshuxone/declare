package main
import "fmt"

// create a function using pointers
// passing pointers to the functions

func main() {
	
	x := 0

	// to pass the pointer to the use "&" keyowrd 
	changeXVAL(&x)
	// print the changed value of x
	fmt.Println("x =",x)

	// print the address of x that it points to
	fmt.Println("Memory address of x is",&x)

}

// passing a pointer to the function
// to pass teh type of the pointer use "*" keyowrd before return type
func changeXVAL(x *int){
	// to get the value of a pointer use "*" keyword and to get the address of a pointer use "&" keyword
	*x = 2

	// above teh value of x will change since we are changing the value on the memory address that x points to

}