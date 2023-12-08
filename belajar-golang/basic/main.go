package main

import (
	"fmt"
	"strings"
)

func main() {
	// typeStruct()
	// typeVariable()
	// typeConst()
	// typeOperator()
	// typeKondisi()
	// typePerulangan()
	// typeArray()
	// typeSlice()
	// typeMap()

	names := []string{"Rizky", "Ziaul", "Haq"}
	printMessage("\nHai", names)

	// var randomValue int

	// randomValue = randomWithRange(2, 10)
	// fmt.Println("\nRandom number: ", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("Random number: ", randomValue, "\n")
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("Random number: ", randomValue, "\n\n")

	// divideNumber(10, 2)
	// divideNumber(4, 0)
	// divideNumber(8, 4)

	var diameter float64 = 15
	area, circumference := calculate(diameter)

	fmt.Printf("\n\nLuas lingkaran:\t\t\t %.2f \n", area)
	fmt.Printf("Keliling lingkaran:\t\t %.2f", circumference)

	var avg = kalkulator(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	msg := fmt.Sprintf("\n\nRata-rata: %.2f", avg)
	fmt.Println(msg)

	// yourHobies("Rachi", "Membaca", "Menulis") // cara 1
	hobbies := []string{"Membaca", "Mengetik"} // cara 2
	yourHobies("Zia", hobbies...)

	// closure mencari nilai terkecil dan terbesar
	getMinMax := func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	numbers := []int{12, 23, 34, 53, 4, 72, 33}
	var min, max = getMinMax(numbers)
	fmt.Printf("\n\nData : %v\nMin  : %v\nMax  : %v\n", numbers, min, max)

	// IFFE Closure
	nilai := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var nilaiBaru = func(min int) []int {
		var r []int
		for _, e := range nilai {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)
	fmt.Println("\nNilai asli: ", nilai)
	fmt.Println("Nilai filter: ", nilaiBaru)

	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "a")
	})
	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("\ndata asli \t\t:", data)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan]

	// Pointer
	pointerNumber := 5
	fmt.Println("\nbefore:", pointerNumber)

	change(&pointerNumber, 10)
	fmt.Println("after: ", pointerNumber)

	// struct
	structUBJ()
	structSlice()

	// method
	tryMethod()
}
