package main 
import "fmt"
func main(){
	num := 3
	doubleNum := func() int{
		num = num * 2
		return num
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
}