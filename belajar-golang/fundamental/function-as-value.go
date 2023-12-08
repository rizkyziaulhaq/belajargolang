package main

import "fmt"

func getGoodBye(name string) string{
	return "Good Bye " + name
}

func functionAsValue(){
	//Function jadi variable
	goodbye := getGoodBye
	fmt.Println(goodbye("Zia"))	
}