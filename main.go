package main

import (
    "fmt"
    "net/http"
    "html"
)

func main() {

    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world")
    })

    log.Fatal(http.ListenAndServe(":8080", null))
}