package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Zia"
	} else {
		return "Hello " + name
	}
}

func returnValue() {
	result := getHello("Ardiana")
	fmt.Println(result)

	fmt.Println(getHello(""))
}