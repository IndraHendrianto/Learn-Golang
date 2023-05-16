package main

import "fmt"

func sayHello() {
	fmt.Println("hello world")
}

func main() {
	sayHello()
	sayHello()
	sayHello()
	for i := 0; i < 10; i++ {
		sayHello()
	}

}
