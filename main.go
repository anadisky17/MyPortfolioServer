package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// type person struct {
// 	name    string
// 	profile string
// }

type msg struct {
	Name    string
	Email   string
	Subject string
	Message string
}

// func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	http.FileServer()

// }

func (m msg) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println("Error in dumping request", err)
	}
	fmt.Fprint(w, string(dump))

}

func main() {

	fmt.Println("Starting Web server...")
	// person := person{
	// 	name:    "anadi",
	// 	profile: "Sre/Devops/Webdev",
	// }

	msgw := new(msg)
	http.Handle("/", http.FileServer(http.Dir("./assets")))

	http.Handle("/sendmsg", msgw)

	http.ListenAndServe("0.0.0.0:80", nil)

}
