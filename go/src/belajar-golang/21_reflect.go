package main

import "fmt"
import "reflect"

// Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkan memanipulasinya.

func main() {
    var number = 23
    var reflectValue = reflect.ValueOf(number) 

    fmt.Println("tipe  variabel :", reflectValue.Type()) // .Type() digunakan untuk cek tipe data

    if reflectValue.Kind() == reflect.Int { // cek apa tipe datanya int
        fmt.Println("nilai variabel :", reflectValue.Int())
    }
}