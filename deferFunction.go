package main
import "fmt"

func main() {
	// defer key word is used to execute any function at the last or after the execution of all contents in the main
	// here defer will execute printOne() function after printTwo() even the printOne function is written before the latter functions
	// hence the defer is used for wrapping things up like closing a databse connection after the query executed successfully
	// defer makes it confirm that the function will execute at the end after everything executed succesfully
	defer printOne()
	printTwo()
}

func printOne(){ fmt.Println(1) }
func printTwo(){ fmt.Println(2) }