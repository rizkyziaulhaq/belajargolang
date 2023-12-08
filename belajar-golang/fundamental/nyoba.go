package main

import "fmt"

type CusTomer struct {
	name, jk, alamat string
	umur             int
}

//func (a CusTomer) sayHai(name string) {
//	fmt.Println("Hai, nama kamu siapa?", a.name)
//}

func nyobaAja() {
	rachi := CusTomer{
		name:   "Rachi",
		jk:     "Pria",
		alamat: "Jl.Pisang",
		umur:   18,
	}

	fmt.Println(rachi)
}

