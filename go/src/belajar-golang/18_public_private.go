package main

import "fmt"

// modifier public dan private dibedakan dengan melihat huruf awalannya. bernilai public jika awalannya huruf kapital dan bernilai private jika awalannya huruf kecil
// modifier private tidak bisa diakses dari package lain
// modifier public bisa diakses dari package lain dengan cara <nama_package>.<method/fungsi/struct> dengan syarat nama packagenya harus diimport terlebih dahulu

func SayHello() { // modifier public karena huruf awalannya menggunakan huruf kapital
	fmt.Println("Hello")
}

func introduce(name string) { // modifier private karena huruf awalannya menggunakan huruf kecil
	fmt.Println("nama saya", name)
}