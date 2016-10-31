package main

import(
	"fmt"
	"os"
	"log"
	"io/ioutil"
)

func main() {
	
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	// write to a text file
	file.WriteString("This is a random text")

	file.Close()

	// read teh created file

	stream, err := ioutil.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	// convert the file in bytes to string
	readString := string(stream)

	fmt.Println(readString)
}