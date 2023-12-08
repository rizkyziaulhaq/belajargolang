package main

import "fmt"

func booleanOperator() {
	var ujian = 80
	var absensi = 75

	var lulusUjian bool = ujian >= 80
	var lulusAbsensi bool = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus bool = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)
}