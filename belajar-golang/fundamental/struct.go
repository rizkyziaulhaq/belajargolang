package main

import "fmt"

type Customer struct {
	Name, Address, Kabar string
	Age                  int
}

//struct-function
func (customer Customer) sayHi(name string) {
	fmt.Println("Hai", name, "nama saya", customer.Name)
}

//struct-function
func (a Customer) saySup() {
	fmt.Println("Hei, darimana asalmu?", a.Address)
}

//struct-function
func (b Customer) sayAge() {
	fmt.Println("Berapa usiamu?", b.Age)
}

func (c Customer) sayLife() {
	fmt.Println("Apakah kamu sehat?", c.Kabar)
}

func sayStruct() {
	var zia Customer
	zia.Name = "Rizky Ziaul Haq"
	zia.Address = "Bekasi Timur"
	zia.Kabar = "Alhamdulillah, Sehat"
	zia.Age = 18

	zia.sayHi("Rey")
	zia.saySup()
	zia.sayLife()
	zia.sayAge()

	fmt.Println(zia)

	rachi := Customer{
		Name:    "Rachi",
		Address: "Indonesia",
		Age:     17,
	}
	fmt.Println(rachi)

	//Cara yg ketiga ini rentan error disarankan pake 1 atau 2
	//reysan := Customer{"Reysan", "Vinland", 16}
	//fmt.Println(reysan)
}
