package main 
import "fmt"
func main() {
	sliceNum := []int { 5,4,3,2,1 }
	fmt.Println( sliceNum[2:] ); // get the value of slice starting from index 2
	fmt.Println( sliceNum[:3] ); // get the value of slice starting from 0 index upto index 3
}