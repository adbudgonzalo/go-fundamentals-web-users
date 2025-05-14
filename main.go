package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users", UserServer) 
	
    log.Fatal(http.ListenAndServe(":8080", nil))
}

// endpoint /users
func UserServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, 200, "success in GET")
	case http.MethodPost:
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, 200, "success in POST")
	default:
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, 400, "not found")
	}
}
