package main
import "fmt"
// include math library to calcualte power of a variable and getting pi constant
import "math"

func main() {
	// passing the values top struct of rectangle
	rect := Rectangle{20,50}
	// passing the values top struct of circle
	circ := Circle{4}

	fmt.Println("Area of rectangle is", getArea(rect))
	fmt.Println("Circle Area =", getArea(circ))
}

// create an interface which will be implemented in other methods
// the interface will be created just like struct
// if a metyhod uses this interface it must implemenmts all the properties and methods defined inside it
type Shape interface{
	// defining a method with return type of float
	area() float64
}


// due to interface all the functions and methods are connected to each other

// create a struct of rectangle and cirlce which will implement interface method area
type Rectangle struct{
	height float64
	width float64
}

type Circle struct{
	radius float64
}

// pass the rectangle struct as the receiver to the method
func (r Rectangle) area() float64{
	return r.width * r.height
}

// calculate the area of circle
func (c Circle) area() float64{
	return math.Pi * math.Pow(c.radius, 2)
}

// create a common getArea function to get the area
func getArea(shape Shape) float64{
	return shape.area()
}