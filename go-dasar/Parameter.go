package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Antoldi"
	sayHelloTo("Indra", "Hendrianto")
	sayHelloTo(firstName, "Baco")
}
