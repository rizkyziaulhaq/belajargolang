package main

import "fmt"

func array() {
	// Array Manual
	var names [3]string

	names[0] = "Rizky"
	names[1] = "Ziaul"
	names[2] = "Haq"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Array Langsung
	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi))
}

