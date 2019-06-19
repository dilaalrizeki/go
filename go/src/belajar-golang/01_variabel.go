package main

import "fmt"

func main() {
	// Metode pendeklarasian variabel "Manifest Typying"
	// var <nama-variabel> <tipe-data>
	// var <nama-variabel> <tipe-data> = <nilai>

	// Metode pendeklarasian variabel "Inference"
	// <nama-variabel> := <nilai>

	// menggunakan var, tanpa tipe data, menggunakan perantara "=" {Manifest Typying}
	var firstName = "john"

	// tanpa var, tanpa tipe data, menggunakan perantara ":=" {Inference}
	lastName := "wick"

	fmt.Println("halo", firstName, lastName)
	
	// Multi variabel
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Println(first, second, third)

	var fourth, fifth, sixth string = "empat", "lima", "enam"
	fmt.Println(fourth, fifth, sixth)

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	fmt.Println(seventh, eight, ninth)

	// inference dengan beda tipe data
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(one, isFriday, twoPointTwo, say)

	// Penulisan variabel menggunakan keyword "new"
	name := new(string)

	fmt.Println(name)   // 0x20818a220
	fmt.Println(*name)  // ""

}