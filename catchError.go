package main
import "fmt"

func main() {
	fmt.Println(safeDiv(6,0));
	fmt.Println(safeDiv(6,3));
}

// create a function safeDiv to divide two numbers to generate an exception divide the number by zero and catch the exception using recover() function

func safeDiv(num1, num2 int) int{
	// to execute the recover() function at the end of completing division use defer keyword
	// below function is used to catch an error if occurs
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2;
	return solution;

} 