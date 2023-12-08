package main

import "fmt"

func saySwitch(){
	var name string
	fmt.Println("Masukan Nama :")
	fmt.Scan(&name) 

	switch name {
	case "Rizky":
		fmt.Println("Hello Rizky")
	case "Ziaul":
		fmt.Println("Hello Ziaul")
	case "Haq":
		fmt.Println("Hello Haq")
	default:
		fmt.Println("Hai, boleh kenalan?")
	}

	// Switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}