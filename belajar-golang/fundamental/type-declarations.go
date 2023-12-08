package main

import "fmt"

func typeDeclaration() {
	type NoKTP string
	type Married bool

	var noKtpZia NoKTP = "192010238"
	var marriedStatus Married = true
	
	fmt.Println(noKtpZia)
	fmt.Println(marriedStatus)
}