package main

import (
	"fmt"
	"net/http"

	ms "github.com/anadisky17/MyPortfolioServer/internal"
)

// type person struct {
// 	name    string
// 	profile string
// }

type msk int

// func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	http.FileServer()

// }

func (M msk) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// dump, err := httputil.DumpRequest(r, true)
	// if err != nil {
	// 	fmt.Println("Error in dumping request", err)
	// }
	// fmt.Fprint(w, string(dump))
	// Seeting default reciever vaule
	msg := new(ms.Msg)
	msg.Name = "Aparna J"
	msg.Message = "Test Body of Email"
	msg.RecipientEmail = "aparna.joshi.1994@gmail.com"
	msg.Subject = "Test Subject"
	r.ParseForm()
	msg.Name = r.Form.Get("Name")
	msg.RecipientEmail = r.Form.Get("Email")
	msg.Message = fmt.Sprint(r.Form.Get("Message") + "  " + "Sender Email is " + r.Form.Get("Email"))
	msg.Subject = r.Form.Get("Subject")

	ms.SendMail(*msg)

	fmt.Fprint(w, "message sent succesfully!!")

}

func main() {

	fmt.Println("Starting Web server...")
	// person := person{
	// 	name:    "anadi",
	// 	profile: "Sre/Devops/Webdev",
	// }

	msgw := new(msk)
	http.Handle("/", http.FileServer(http.Dir("./assets")))

	http.Handle("/sendmsg", msgw)

	http.ListenAndServe("0.0.0.0:80", nil)

}
