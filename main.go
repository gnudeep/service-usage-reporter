package main

import (
    "fmt"
    "net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, this is a simple message from the microservice!")
}

func main() {
    http.HandleFunc("/", messageHandler)
    http.ListenAndServe(":8080", nil)
}

