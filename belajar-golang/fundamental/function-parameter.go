package main

import "fmt"

func sayHelloTo(firstName string, middleName string, lastName string, age int) {
	fmt.Println("Hello", firstName, middleName, lastName, age)
}

func parameter() {
	sayHelloTo("Rizky", "Ziaul", "Haq", 18)
}