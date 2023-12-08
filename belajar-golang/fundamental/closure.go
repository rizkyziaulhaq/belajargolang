package main

import "fmt"

func closure() {
	name := "Rachi"
	counter := 0

	increment := func() {
		name := "Zia"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
