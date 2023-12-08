package main

import "fmt"

//Alias atau type declaration
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Monyet" {
		return "..."
	} else if name == "Anjing" {
		return "..."
	} else if name == "Babi" {
		return "..."
	} else {
		return name
	}
}

func functionAsParameter(){
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Monyet", spamFilter)
	sayHelloWithFilter("Zia", spamFilter)
	sayHelloWithFilter("Babi", spamFilter)

	//filter := spamFilter
	//sayHelloWithFilter("Anjing",filter)
}