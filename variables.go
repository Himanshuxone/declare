package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var age int = 26
	var favnum float64 = 1.6189002
	randNum := 1
	fmt.Println(randNum)
	fmt.Println(favnum)
	fmt.Println(age)

	// for loops
	fmt.Println("loop starts")
	i := 1;
	for i < 5 {
		fmt.Println(i)
		i++
	}
}