package WebServer

import (
	"fmt"
	"net/http"
)

func Exec() {
	web()
}

func web() {
	// Takes in routes and use funcs to deal with them
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	// in order to serve, we need to use HTTP to listen and serve on a port
	fmt.Println("Server starting at 127.0.0.1:3000")

	// Next line disabled so program continues execution progression
	// (once you serve, unless multi-threaded, application will pause to listen for web requests
	//http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// output something to browser
	fmt.Fprintf(w, "<h1>Hey, a H1 header!</h1>")
	fmt.Fprintln(w, "Hello browser world")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>This is the ABOUT route handler page.</h1>")
}
