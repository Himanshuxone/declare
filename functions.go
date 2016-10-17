package main

import (
	"fmt"
	"os"
)

func main() {
	// create a float array and pass it to a function
	listnums := []float64{1,2,3,4,5}

	// calling the function 
	fmt.Println( sumvalues(listnums) )
	os.Exit(1) // to exit the program import os library and use os.Exit(1)
}

/*
	create a function sumvalues outside of main loop through the nums array, calculate the sun and return the float

*/
func sumvalues(nums []float64) float64{ // func <funcname>( function arguents <arg type> ) return <type> => is the definition of the function 
	fmt.Println(nums)
	sum := 0.0
	for _,val := range nums {
		sum += val
	}
	return sum
	
}