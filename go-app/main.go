package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hey There, This is just a sample GO App")
    })

    http.ListenAndServe(":5002", nil)
}
