package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpIndra NoKTP = "123456789"
	var marriedStatus Married = true
	fmt.Println(ktpIndra)
	fmt.Println(marriedStatus)
}
