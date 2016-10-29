package main
import "fmt"
func main() {
	demoPanic();
}

// define a function to try catch an error

func demoPanic(){
	// use defer keyword to run the function after execution of enclosing function completely
	defer func() {
		fmt.Println(recover())
	}()

	// create an error which will be catched by recover() which is used to catch exceptions
	panic(" PANIC !.......... ")
}