package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// SlackMsg represents the content of the slack message sent when contact us if filled out
type SlackMsg struct {
	Text string `json:"text"`
}

func main() {
	port := os.Getenv("SG_ADDR")
	if port == "" {
		port = ":10001"
	}

	hookURL := os.Getenv("SG_SLACK_HOOK")
	if hookURL == "" {
		panic("Set SG_SLACK_HOOK environment variable to a valid slack url")
	}

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// read post args,
			// TODO: Do something with the form value content.
			name := r.FormValue("Name")
			email := r.FormValue("Email")
			phone := r.FormValue("Phone")
			message := r.FormValue("Message")

			data := bytes.NewBuffer(nil)
			enc := json.NewEncoder(data)

			msg := SlackMsg{Text: fmt.Sprintf("Yay! A person wants to be contacted\nName: %v\nEmail: %v\nPhone: %v\nMessage: %v", name, email, phone, message)}

			err := enc.Encode(msg)
			if err != nil {
				panic(err)
			}

			resp, err := http.Post(hookURL, "application/json", data)
			if err != nil {
				panic(err)
			}
			if resp.StatusCode != http.StatusOK {
				panic("wrong status")
			}

		}

		tmpl, err := template.ParseFiles("contact.html")
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	})

	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Serving on " + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
