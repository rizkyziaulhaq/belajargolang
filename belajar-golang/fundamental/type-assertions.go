package main

import "fmt"

func random() interface{} {
	return 1000
}

func typeAssertions() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	case bool:
		fmt.Println("Value", value, "is bool")
	default:
		fmt.Println("Unknown")
	}

	//resultString := result.(string)
	//fmt.Println(resultString)

	//resultInt := result.(int) //panic
	//fmt.Println(resultInt)
}