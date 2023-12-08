package main

import (
	"fmt"
)

func forUntuk(){
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
		
	}

	// For Range
	slice := []string{"Rizky", "Ziaul", "Haq"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Zia"
	person["title"] = "Student"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}