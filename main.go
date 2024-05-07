package main

import (
	"fmt"
	routes "myapp/route"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	routes.InitilizeRoutes()
}

// Define handler function to handle incoming request
// func homeHandler(w http.ResponseWriter, r *http.Request) { // is a type of W, it genereate response, second one passes the original value so pointer is written.
//  p := mux.Vars(r) //gorrila mux, var is varibale(changeable date) //value is what we type in url
//  fmt.Println("the value given is", p)
//  course := p["course"]
//  // Write the response to client
//  _, err := w.Write([]byte("the course is " + course)) //based on what platform you used, it will be called there or displayed(this w)
//  if err != nil {
//   fmt.Println("error:", err)
//  }

// }

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	fmt.Println("the value given is", p)
	course := p["course"]
	// Write the response to client
	_, err := w.Write([]byte("My name is " + course)) //w.write=var two bcuz it returns err or no of bytes
	if err != nil {
		fmt.Println("error:", err)
	}
}
