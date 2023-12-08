package main

import (
	"fmt"
	"math"
	"strings"
)

func printMessage(message string, arr []string) {
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// fungsi dengan return value/nilai balik
// var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

// func randomWithRange(min, max int) int {
// 	value := randomizer.Int()%(max-min+1) + min
// 	return value
// }

// penggunaan keyword return untuk menghentikan proses dalam fungsi
// func divideNumber(m, n int) {
// 	if n == 0 {
// 		fmt.Printf("Invalid divider. %d cannot divided by %d\n", m, n)
// 		return
// 	}

// 	res := m / n
// 	fmt.Printf("%d / %d\n", m, n, res)
// }

// penerapan fungsi multiple return
func calculate(d float64) (float64, float64) {
	area := math.Pi * math.Pow(d/2, 2) // hitung luas
	circumference := math.Pi * d       // hitung keliling

	return area, circumference // kembalikan 2 nilai
}

// fungsi variadic
func kalkulator(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	avg := float64(total) / float64(len(numbers))
	return avg
}

func yourHobies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, ", ")

	fmt.Printf("\nHello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s", hobbiesAsString)
}

// fungsi sebagai parameter
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

// Pointer
func change(original *int, value int) {
	*original = value
}

// struct
type MahasiswaUBJ struct {
	Id       int
	Npm      string
	Nama     string
	Semester int
	Prodi    string
}

func structUBJ() {
	var _ = MahasiswaUBJ{
		Id:       001,
		Npm:      "20231071782",
		Nama:     "Zia",
		Semester: 6,
		Prodi:    "Desain",
	}

	var maba = MahasiswaUBJ{}
	maba.Id = 002
	maba.Npm = "20231071783"
	maba.Nama = "Rachi"
	maba.Semester = 4
	maba.Prodi = "Industri"

	mhs := []MahasiswaUBJ{
		{
			Id:       1101,
			Npm:      "202210715048",
			Nama:     "Rizky Ziaul Haq",
			Semester: 3,
			Prodi:    "Informatika",
		},
		{
			Id:       1102,
			Npm:      "202210715049",
			Nama:     "Ardiana",
			Semester: 3,
			Prodi:    "Akuntansi",
		},
		{
			Id:       1103,
			Npm:      "202210715050",
			Nama:     "Birjon",
			Semester: 5,
			Prodi:    "Manajemen",
		},
	}

	fmt.Println("\nId	: ", maba.Id)
	fmt.Println("Npm	: ", maba.Npm)
	fmt.Println("Nama	: ", maba.Nama)
	fmt.Println("Smt	: ", maba.Semester)
	fmt.Println("Prodi	: ", maba.Prodi)

	// fmt.Printf("\n%v", maba)
	fmt.Printf("\n%v", mhs)

	for i, m := range mhs {
		fmt.Printf("\nindex: %d data: %v", i, m)
	}
}

// slice & struct
type person struct {
	name string
	age  int
}

func structSlice() {
	var allStudents = []person{
		{name: "\n\nWick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	var semuaSiswa = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}

	for _, student := range semuaSiswa {
		fmt.Println(student)
	}
}

type siswa struct {
	nama string
	umur int
}

func (s siswa) sayHello() {
	fmt.Println("Hello", s.nama)
}

func (s siswa) getNameAt(i int) string {
	return strings.Split(s.nama, " ")[i-1]
}

func tryMethod() {
	var siba = siswa{"Rizky Ziaul Haq", 18}
	siba.sayHello()

	name := siba.getNameAt(2)
	fmt.Println("Nama panggilan :", name)
}
