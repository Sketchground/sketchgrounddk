package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Serving on :8080")
	http.ListenAndServe(":8080", nil)
}
