package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type SlackMsg struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// read post args,
			// TODO: Do something with the form value content.
			name := r.FormValue("Name")
			email := r.FormValue("Email")
			phone := r.FormValue("Phone")
			message := r.FormValue("Message")

			// url for webhook to slack...
			// https://hooks.slack.com/services/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
			// curl -X POST --data-urlencode 'payload={"channel": "#websitecontacts", "username": "webhookbot", "text": "This is posted to #websitecontacts and comes from a bot named webhookbot.", "icon_emoji": ":ghost:"}' https://hooks.slack.com/services/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

			// payload={"text": "A very important thing has occurred! <https://alert-system.com/alerts/1234|Click here> for details!"}
			data := bytes.NewBuffer(nil)
			enc := json.NewEncoder(data)

			msg := SlackMsg{Text: fmt.Sprintf("Yay! A person wants to be contacted\nName: %v\nEmail: %v\nPhone: %v\nMessage: %v", name, email, phone, message)}

			err := enc.Encode(msg)
			if err != nil {
				panic(err)
			}

			resp, err := http.Post("https://hooks.slack.com/services/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", "application/json", data)
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
	fmt.Println("Serving on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
