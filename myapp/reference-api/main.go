package main

import (
	"log"
	"net/http"
	"github.com/practice-cloud-run-functions/myapp/reference-api/handler"
)

func main() {
	http.HandleFunc("/api/users", handler.GetUsers)
	log.Println("Reference API running on :8080")
	http.ListenAndServe(":8080", nil)
}
