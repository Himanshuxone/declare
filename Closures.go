package main 

import "fmt"

func main() {
	
	list := []float64{1,2,6,7,9}
	// fmt.Println(factorial(5))
	fmt.Println(sum(list))

}

func factorial(num int) int{

	if num == 0 {
		return 1
	}else {
		return num * factorial(num - 1)
	}

}

func sum(list []float64) float64{

	var sum float64;
	for _,value := range list {
		sum += value	
	}
	return sum;
}