package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	names := []string{"Indra", "Hendrianto", "Antoldi", "Baco"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for i, value := range names {
		fmt.Println("orang ke", i, "adalah", value)
	}

	person := make(map[string]string)
	person["name"] = "Indra"
	person["title"] = "QAE"

	for key, value := range person {
		fmt.Println(key, "=", value)

	}

}
