package main
// importing many packages will be done using () as below
// importing string packages for using string functions
import( 
	"fmt" 
	"strings" 
	"sort"
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

 	// create a csv list of values to join and apply more string operations
 	csvString := "1,2,3,4,5"

 	// split the string from comma
 	fmt.Println( strings.Split( csvString, "," ) )

 	// create an array of string type

 	listOfLetters := []string{"c","b","a"}

 	// sort the liost of letters
 	sort.Strings(listOfLetters)
 	fmt.Println(listOfLetters)

 	// join the letters to using comma
 	fmt.Println(strings.Join(listOfLetters, ","))

 	// create a list of letters by joining using comma
 	listOfNums := strings.Join([]string{"1","2","3"}, ",")
 	fmt.Println(listOfNums)

} 