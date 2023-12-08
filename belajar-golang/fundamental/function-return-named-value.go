package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "Rizky"
	middleName = "Ziaul"
	lastName = "Haq"

	return
}

func returnNamedValue() {
	a, b, c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}