package main 
import "fmt"
func main(){
	num := 3
	doubleNum := func() int{
		return num := num * 2
	}
	fmt.Println(doubleNum)
}