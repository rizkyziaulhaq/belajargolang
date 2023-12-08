package main

import "fmt"

func constant() {
//Multiple constant

	const (
		firstName string = "Rizky"
		lastName = "Ziaul Haq"
		value = 2004
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

}