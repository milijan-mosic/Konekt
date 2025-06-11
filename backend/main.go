package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	address := "5000"
	urlPrefix := "/api/1.0"

	http.HandleFunc(urlPrefix+"/", helloWorld)

	fmt.Printf("Listening to address: http://localhost:%s\n", address)
	http.ListenAndServe(":"+address, nil)
}
