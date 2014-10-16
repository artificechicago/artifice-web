package main

import (
	"net/http"
	"html/template"
	"github.com/sendgrid/sendgrid-go"
	"log"	
)

var signup = template.Must(template.ParseFiles("website/sign-up.html"))

func sendMail() {
	sg := sendgrid.NewSendGridClient("", "") //This will fail
	message := sendgrid.NewMail()
	message.AddToName("James Crooks")
	message.AddTo("crooks1379@gmail.com")
	message.SetSubject("SendGrid Goroutine Test")
	message.SetText("WIN... concurrently!")
	message.SetFrom("no-reply@artificechicago.org")

	err := sg.Send(message)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Email sent!")
	}
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log.Println("POST'd")
		go sendMail()
	}
	signup.Execute(w, nil)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("website")))
	http.HandleFunc("/sign-up.html", handleSignUp)
	panic(http.ListenAndServe(":3000", nil))

}
