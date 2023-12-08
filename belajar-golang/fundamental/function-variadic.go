package main

import "fmt"

func sumAll(numbers ...int) int{
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func variadic() {
	total := sumAll(20, 30)
	fmt.Println(total)

	slice := []int{11,89}
	total = sumAll(slice...)
	fmt.Println(total)
}