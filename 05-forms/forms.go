package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		fmt.Printf("Email: %s, Subject: %s, Message: %s", details.Email, details.Subject, details.Message)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	log.Panic(http.ListenAndServe(":8080", nil))
}
