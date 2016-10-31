package main
// importing many packages will be done using () as below
// importing string packages for using string functions
import( 
	"fmt" 
	"strings" 
)

func main() {
 	
 	// assign a string to a variable

 	sampleString := "Hello World"
 	// check if a string contains a character
 	fmt.Println(strings.Contains(sampleString, "lo"))

 	// calculate the first index of a charachter in a string
 	fmt.Println(strings.Index(sampleString, "lo"))

 	// cacluate the count/no of a appearence of a character n a string
 	fmt.Println(strings.Count(sampleString, "l"))

 	// replace characters in a string
 	// here replace "l" character with "x" starting with third index
 	fmt.Println(strings.Replace(sampleString, "l", "x", 3))

} 