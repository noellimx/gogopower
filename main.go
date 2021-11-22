package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "woosh")
}

func main() {
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Starting server @ port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greeting(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
