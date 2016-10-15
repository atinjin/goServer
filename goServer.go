package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)


type LoginRequest struct {
	userId string
	password string
	service int
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Golang Server:%s", r.URL.Path[1:])
	fmt.Fprintf(w, "Your auth header is %s", r.Header.Get("Authorization"))
	if(r.URL.Path[1:] == "/session") {
		dec := json.NewDecoder(r.Body)
		var req LoginRequest
		err := dec.Decode(&req)
		if err != nil {
			log.Fatal("Error $s", err)
		}
		defer r.Body.Close()
		resBody, err := json.Marshal(req)
		fmt.Fprintf(w, "Request Body is %s", resBody )
	}
}

func main() {
	fmt.Print("Hello, Golang");
	http.HandleFunc("/test", handler)
	http.HandleFunc("/session", handler)

	http.ListenAndServe(":8080", nil)
}
