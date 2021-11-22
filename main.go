package main

import "fmt"

func main() {
	fmt.Printf("Hello World 1!")
}

func greeting(name string) string {
	return fmt.Sprintf("Hi %s", name)
}