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

	// functon has paramters a, b is of integer type and it will return the value of type int as a+b = int
	sum := func(a, b int) int { return a+b } (3, 4)

	// arrays

	var favNum[5] float64
	favNum[0] = 165
	favNum[1] = 3.141
	favNum[2] = 43.9
	favNum[3] = 226
	favNum[4] = 5

	for j, value := range favNum {
		fmt.Println(value, j)
	}

	// := assignment operator is used to create an array of length 5 
	// another way to create an array using assignment operator	 
	favNums1 := [5]float64 {1,7,6,4,5}
	for _, value := range favNums1 {
		fmt.Println(value);
	}
}