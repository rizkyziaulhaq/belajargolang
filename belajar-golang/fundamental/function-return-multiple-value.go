package main

import "fmt"

func getFullName() (string, string, string)  {
	return "Rizky", "Ziaul", "Haq"
}


func returnMultipleValue()  {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}