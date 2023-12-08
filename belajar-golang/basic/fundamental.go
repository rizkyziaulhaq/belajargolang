package main

// func typeVariable() {
// 	// := digunakan pd deklarasi awal
// 	id := 1
// 	var npm = "202210715048"
// 	var nama = "Rachi"
// 	var smt = "3"

// 	var firstName string = "Rizky"

// 	var middleName string
// 	middleName = "Ziaul"

// 	// Multiple var
// 	var first, second, third = 1, "Rachi", "Belajar Golang"

// 	// _, = menandakan utk tidak pakai var
// 	_, dua, tiga := 2, "Ern", "Belajar Java"

// 	var (
// 		lastName = "Haq"
// 		junkFood = "🍔"
// 	)

// 	// new = digunakan untuk membuat variabel pointer dengan tipe data tertentu
// 	name := new(string)

// 	// make = biasa digunakan di channel, slice, map

// 	fmt.Printf("Data Mahasiswa : \n ID: %v \n NPM: %s \n Nama: %s \n Semester: %s", id, npm, nama, smt)
// 	fmt.Printf("Hai %s %s %s %s %s!\n", nama, firstName, middleName, lastName, junkFood)
// 	fmt.Println("Hai", nama, firstName, middleName, lastName, junkFood)
// 	fmt.Println(first, second, third)
// 	fmt.Println(dua, tiga)
// 	fmt.Println(*name)
// }

// // Multiple const
// func typeConst() {
// 	const (
// 		today string = "Jumat"
// 		sekarang
// 		isToday = true
// 	)

// 	const (
// 		id   = 1
// 		npm  = "202210715048"
// 		nama = "Rachi"
// 		smt  = "3"
// 	)
// 	fmt.Println(today, sekarang, isToday)
// 	fmt.Println(id, npm, nama, smt)
// }

// func typeOperator() {
// 	// Aritmatika + - * / %
// 	// value := (((2+6)%3)*4 - 2) / 3

// 	// Perbandingan == != < <= > >=
// 	// var isEqual = (value == 2)

// 	// Logika && || !
// 	left, right := false, true

// 	var (
// 		leftAndRight = left && right
// 		leftOrRight  = left || right
// 		leftReverse  = !left
// 	)

// 	// fmt.Printf("Hasil: %v (%t) \n", value, isEqual)
// 	fmt.Printf("left && right \t(%t) \n", leftAndRight)
// 	fmt.Printf("left || right \t(%t) \n", leftOrRight)
// 	fmt.Printf("!left \t\t(%t) \n", leftReverse)
// }

// func typeKondisi() {
// 	point := 10

// 	if point == 10 {
// 		fmt.Println("Lulus dengan nilai sempurna")
// 	} else if point >= 5 {
// 		fmt.Println("Lulus")
// 	} else if point == 4 {
// 		fmt.Println("Hampir lulus")
// 	} else {
// 		fmt.Printf("Tidak lulus, nilai anda %v\n", point)
// 	}

// 	value := 8040.0

// 	if percent := value / 100; percent >= 100 {
// 		fmt.Printf("%.1f%s Perfect!\n", percent, "%")
// 	} else if percent >= 70 {
// 		fmt.Printf("%.1f%s Good!\n", percent, "%")
// 	} else {
// 		fmt.Printf("%.1f%s Not bad!\n", percent, "%")
// 	}

// 	// switch
// 	nilai := 6

// 	switch nilai {
// 	case 10:
// 		fmt.Println("Perfect!")
// 	case 8, 7:
// 		fmt.Println("Awesome")
// 	default:
// 		{
// 			fmt.Println("Not bad")
// 			fmt.Println("You can better me!")
// 		}
// 		// fmt.Println("Not bad")
// 	}

// 	// switch dengan gaya if - else
// 	hasil := 10

// 	switch {
// 	case hasil == 10:
// 		fmt.Println("Perfect!")
// 	case (hasil <= 8) && (hasil >= 6):
// 		fmt.Println("Awesome")
// 		fallthrough // dipakei memaksa proses pengecekan diteruskan ke satu case selanjutnya dengan tanpa menghiraukan nilai kondisinya
// 	default:
// 		{
// 			fmt.Println("Not bad")
// 			fmt.Println("You need to learn more!")
// 		}
// 	}

// 	// Seleksi kondisi bersarang
// 	angka := 8

// 	if angka >= 7 {
// 		switch angka {
// 		case 10, 9:
// 			fmt.Println("Perfect!")
// 		default:
// 			{
// 				fmt.Println("Noice!")
// 				fmt.Println("Keep strong!")
// 			}
// 		}
// 	} else {
// 		if angka == 5 {
// 			fmt.Println("Not bad!")
// 		} else if angka == 3 {
// 			fmt.Println("Keep trying!")
// 		} else {
// 			fmt.Println("You can do it!")
// 			if angka == 0 {
// 				fmt.Println("Try harder!")
// 			}
// 		}
// 	}
// }

// func typePerulangan() {
// 	// cara 1
// 	for i := 0; i <= 5; i++ {
// 		// fmt.Println("Angka ke-", i)
// 	}

// 	// cara 2
// 	a := 0
// 	for a <= 5 {
// 		// fmt.Println("Angka", a)
// 		a++
// 	}

// 	// cara 3
// 	b := 0
// 	for {
// 		// fmt.Println("Angka", b)

// 		b++
// 		if b == 5 {
// 			break // memberhentikan perulangan
// 		}
// 	}

// 	// cara 4
// 	xs := "123" // string
// 	for i, v := range xs {
// 		fmt.Println("Index:", i, "Value:", v)
// 	}

// 	ys := [5]int{10, 20, 30, 40, 50} // array
// 	for _, v := range ys {
// 		fmt.Println("Value:", v)
// 	}

// 	zs := ys[0:2] // slice
// 	for _, v := range zs {
// 		fmt.Println("Value:", v)
// 	}

// 	kvs := map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
// 	for k, v := range kvs {
// 		fmt.Println("Key:", k, "Value:", v)
// 	}

// 	for range kvs {
// 		fmt.Println("Done")
// 	}

// 	// penggunaan break - continue
// 	for c := 1; c <= 10; c++ {
// 		if c%2 == 1 {
// 			continue
// 		}

// 		if c > 8 {
// 			break
// 		}

// 		fmt.Println("Angka:", c)
// 	}

// 	// perulangan bersarang
// 	for d := 0; d < 5; d++ {
// 		for e := d; e < 5; e++ {
// 			fmt.Print(e, " ")
// 		}

// 		fmt.Println()
// 	}

// 	// pemanfaatan label dalam perulangan
// outerLoop:
// 	for f := 0; f < 5; f++ {
// 		for g := 0; g < 5; g++ {
// 			if f == 2 {
// 				break outerLoop
// 			}
// 			fmt.Println("Matriks: [", f, "][", g, "]")
// 		}
// 	}
// }

// func typeArray() {
// 	// cara 1
// 	var name [3]string
// 	name[0] = "Rizky"
// 	name[1] = "Ziaul"
// 	name[2] = "Haq"

// 	// cara 2 - horizontal
// 	var fruits = [4]string{"Apel", "Mangga", "Jambu", "Pisang"}

// 	// cara 3 - vertikal
// 	var buah = [4]string{
// 		"Apel",
// 		"Mangga",
// 		"Jambu",
// 		"Pisang",
// 	}

// 	// perulangan elemen array menggunakan keyword for
// 	for i := 0; i < len(buah); i++ {
// 		fmt.Printf("Elemen %v : %s\n", i, buah[i])
// 	}

// 	// perulangan elemen array menggunakan keyword for - range
// 	for i, buah := range buah {
// 		fmt.Printf("Elemen: %v %s\n", i, buah)
// 	}

// 	// penggunaan variabel _ dalam for - range
// 	for _, buah := range buah {
// 		fmt.Printf("Nama buah: %s\n", buah)
// 	}

// 	// jika yang dibutuhkan hanya indeks elemen-nya saja, bisa gunakan 1 buah variabel setelah keyword for.
// 	for i, _ := range buah {
// 		fmt.Printf("Elemen: %v\n", i)
// 	}

// 	// alokasi elemen array menggunakan keyword make
// 	cars := make([]string, 2)
// 	cars[0] = "Yamaha"
// 	cars[1] = "Honda"

// 	// inisialisasi nilai awal array tanpa jumlah elemen
// 	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	// array multidimensi
// 	numbers1 := [2][3]string{[3]string{"Rachi", "Zia", "Miralis"}, [3]string{"Ern", "Jahe", "Amel"}}
// 	numbers2 := [2][3]int{{3, 2, 3}, {3, 4, 5}}

// 	fmt.Println(name[0], name[1], name[2])
// 	fmt.Println("Jumlah element\t\t: ", len(fruits))
// 	fmt.Println("Isi semua element\t: ", fruits)
// 	fmt.Println("Data array\t: ", (numbers))
// 	fmt.Println("Jumlah element\t: ", len(numbers))
// 	fmt.Println("numbers1", numbers1)
// 	fmt.Println("numbers2", numbers2)
// 	fmt.Println(cars)
// }

// func typeSlice() {
// 	// cara 1
// 	animal := []string{"Hiu", "Burung", "Badak", "Elang"}
// 	listAnimal := append(animal, "Jerapah")

// 	// cara 2
// 	buahBuahan := []string{"Mangga", "Melon", "Pepaya", "Pisang"}
// 	newBuah := buahBuahan[1:3]

// 	// mengubah data elemen slice baru, yang terbentuk dari slice lama
// 	fruits := []string{"apple", "grape", "banana", "melon"}

// 	aFruits := fruits[0:3]
// 	bFruits := fruits[1:4]

// 	aaFruits := aFruits[1:2]
// 	baFruits := bFruits[0:1]

// 	baFruits[0] = "pinnaple" // buah "grape" diubah menjadi "pinnaple"

// 	// fungsi - copy
// 	dst := make([]string, 2)
// 	src := []string{"Richter", "Rachi", "Zata", "Ziaul"}
// 	n := copy(dst, src)

// 	// dst := []string{"Mangga", "Mangga", "Mangga", "Mangga"}
// 	// src := []string{"Pisang", "Jambu", "Apel"}
// 	// n := copy(dst, src)

// 	fmt.Println(fruits)   // [apple grape banana melon]
// 	fmt.Println(aFruits)  // [apple grape banana]
// 	fmt.Println(bFruits)  // [grape banana melon]
// 	fmt.Println(aaFruits) // [grape]
// 	fmt.Println(baFruits) // [grape]
// 	fmt.Println("Nama hewan: ", animal, "->", listAnimal)
// 	fmt.Println("Nama buah:", newBuah)
// 	fmt.Println(dst)
// 	fmt.Println(src)
// 	fmt.Println(n)
// }

// func typeMap() {
// 	// penggunaan map
// 	var chicken map[string]int
// 	chicken = map[string]int{}

// 	chicken["Januari"] = 50
// 	chicken["Februari"] = 40

// 	// inisialisasi map
// 	data := map[string]int{}
// 	data["One"] = 1

// 	// cara horizontal
// 	// chicken = map[string]int{"Januari": 50, "Februari": 40}

// 	// cara vertikal
// 	// chicken = map[string]int{
// 	// 	"Januari":  50,
// 	// 	"Februari": 40,
// 	// }

// 	// iterasi item map menggunakan for - range
// 	tester := map[string]int{
// 		"Januari":  50,
// 		"Februari": 40,
// 		"Maret":    34,
// 		"April":    67,
// 	}

// 	// for key, val := range tester {
// 	// 	fmt.Println(key, " \t:", val)
// 	// }

// 	// menghapus item map
// 	delete(tester, "Januari")

// 	// mendeteksi keberadaan item dengan key tertentu
// 	// _, isExist := tester["April"]
// 	// if isExist {
// 	// 	// fmt.Println(value)
// 	// } else {
// 	// 	// fmt.Println("Item is not exist")
// 	// }

// 	// kombinasi slice & map
// 	nyoba := []map[string]string{
// 		{"Name": "Rizky Ziaul Haq", "Age": "18", "Gender": "Male", "Color": "Blue"},
// 		{"Address": "Jl.Pahlawan", "No": "60"},
// 		{"Community": "Lenovo"},
// 	}

// 	for _, coba := range nyoba {
// 		fmt.Println(coba["Name"], coba["Age"], coba["Gender"], coba["Color"], coba["Address"], coba["No"], coba["Community"])
// 	}
// 	// fmt.Println(len(tester), "items:", tester)
// 	// fmt.Println("Januari", chicken["Januari"]) // 50
// 	// fmt.Println("Mei", chicken["Mei"])         // 0 karena blm ada item yg tersimpan dgn key Mei
// 	// fmt.Println(data)
// }
