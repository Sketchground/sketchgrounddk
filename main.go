package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// read post args,
			// TODO: Do something with the form value content.
			fmt.Println(r.FormValue("Name"))
			fmt.Println(r.FormValue("Email"))
			fmt.Println(r.FormValue("Phone"))
			fmt.Println(r.FormValue("Message"))
		}

		tmpl, err := template.ParseFiles("contact.html")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, nil)
	})

	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Serving on :8080")
	http.ListenAndServe(":8080", nil)
}
