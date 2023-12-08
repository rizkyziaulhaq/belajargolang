package main

import "fmt"

func sayIf(){
	var name string
	fmt.Println("Masukan Nama :")
	fmt.Scan(&name)

	if name == "Zia" {
		fmt.Println("Hello Zia")
	} else if name == "Ardiana" {
		fmt.Println("Hello Ardiana")
	} else if name == "Ul" {
		fmt.Println("Hello Ul")
	} else {
		fmt.Println("Hai, boleh kenalan?")
	}



	//If short statement
	//if length := len(name); length > 5 {
	//	fmt.Println("Terlalu panjang")
	//} else {
	//	fmt.Println("Nama sudah benar")
	//}
}