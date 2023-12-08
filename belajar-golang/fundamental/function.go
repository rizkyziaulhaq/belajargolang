package main

import "fmt"

func sayHello(){
	fmt.Println("Hello Wolrd")
}

func sayFunction() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}