package main

// import net package to get the http request response functions
import(
	"fmt"
	"net/http"
)

func main() {
	
	// create a an http handler function to handle the request response
	// create a hanlder function to handle the request commint to root directory or index page  
	http.HandleFunc( "/", handler )

	// create another handler function to handle the request coming on /earth url
	http.HandleFunc( "/earth", handler2 )

	// define a port to listen to the request
	fmt.Println("Server is running")
	http.ListenAndServe(":8080", nil)


}

// the function takes request respoonse as a parameter
func handler(w http.ResponseWriter, r *http.Request){
	// write to the browser on hitting the index page
	fmt.Fprintf( w, " Hello World\n " )
}

func handler2(w http.ResponseWriter, r *http.Request ){
	fmt.Fprintf( w, " Hello Earth\n " )
}
