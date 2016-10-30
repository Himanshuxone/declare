package main
import "fmt"

func main() {
	// if you dont the paramters pass the key value form 
	// rect1 := Rectangle{leftX:0,topX50, heightX:10, widthX: 10}
	
	// else you can directly pass the values in the same order 
	rect1 := Rectangle{ 0, 50, 10, 10 }
	fmt.Println("Rectangle is", rect1.widthX, "wide")
	fmt.Println("Area of rectangle is", rect1.area())

}

// define a data type rectangle which is a structure
type Rectangle struct{
	// declare the variables and type of those variables like in a go way
	leftX float64
	topX float64
	heightX float64
	widthX float64

}


// create a function to calculate area of a rectangle
/**
 *  A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, 
 *  and associates the method with the receiver's base type. The receiver is specified via an extra parameter section preceding the method name.
 */

// below is not a functions it is a method with a reciever rectangle struct here
// we pass the reference to the struct using "*" keword which will pass the values to the method
// using methods we can pass the value of a structure from outside main which is a receiver
func (rect *Rectangle) area() float64{
	return rect.widthX * rect.heightX
}