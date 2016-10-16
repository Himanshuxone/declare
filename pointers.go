package main
import "fmt"
func main() {
	i, j := 43, 2701

	p := &i // p points to i
	fmt.Println(*p) // to access the value of pointer put *p to access its value it means here read i through the pointer
	*p = 21 // changes the value of i through pointer p
	fmt.Println(i)

	p := &j		 // point to j
	*p = *p / 37 // divide j through the pointer 
	fmt.Println(j) // see the new value of j
}