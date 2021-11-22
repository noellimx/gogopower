package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloWorldHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "woosh")
}

func main() {
	http.HandleFunc("/helloworld", helloWorldHandler)
	listeningPort := os.Getenv("PORT")
	log.Printf("Starting server @ port: %s", listeningPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", listeningPort), nil))
}

func greeting(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
