package main

import "fmt"

var BuildVersion = "v.0.1.0"

func Hello() error {
	fmt.Println("hello, world!")
	return nil
}

func main() {
	Hello()
	fmt.Printf("version: %s\n", BuildVersion)
}
