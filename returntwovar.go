package main 

import "fmt"

func main() {
	number1, number2 := returnthem(5)
	fmt.Println(number1, number2)
}

// returning two values from the function
func returnthem(number int) (int, int){
	return number + 1, number + 2
} 