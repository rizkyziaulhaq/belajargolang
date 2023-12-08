package main

import "fmt"

func Variabel(){
	var name string

	name = "Rizky Ziaul"
	fmt.Println(name)

	name = "Rizky Haq"
	fmt.Println(name)

	var friedName = "Faisal"
	fmt.Println(friedName)

	// Variabel int
	var age = 16
	fmt.Println(age)

	// := hanya untuk deklarasi awal
	country := "Irlandia"
	fmt.Println(country)

	// Multiple variabel agar enak dilihat
	var (
		firstName = "Rizky"
		lastName = "Ziaul Haq"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}