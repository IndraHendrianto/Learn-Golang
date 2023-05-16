package main

import "fmt"

func main() {
	name := "indra"

	switch name {
	case "indra":
		fmt.Println("Hello Indra")
		fmt.Println("Selamat Pagi")
	case "Baco":
		fmt.Println("Hello Baco")
		fmt.Println("Selamat Siang")
	default:
		fmt.Println("Hello semua")

	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")

	}

	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("Nama terlalu panjang")
	case lenght > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
