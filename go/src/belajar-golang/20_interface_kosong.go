// #1
// Interface kosong atau interface{} adalah tipe data yang bisa menampung segala jenis data, bahkan array, bisa pointer bisa tidak (konsep ini disebut dengan dynamic typing).
package main

import "fmt"

func main() {
    var secret interface{}

    secret = "ethan hunt"
    fmt.Println(secret)

    secret = []string{"apple", "manggo", "banana"}
    fmt.Println(secret)

    secret = 12.4
    fmt.Println(secret)
}
// Keyword interface digunakan untuk pembuatan interface. Tetapi ketika ditambahkan kurung kurawal ({}) di belakang-nya (menjadi interface{}), maka kegunaannya akan berubah, yaitu sebagai tipe data.

// ==========================================================================

// // #2 Casting variabel interface kosong
// package main

// import "fmt"
// import "strings"

// func main() {
//     var secret interface{}

//     secret = 2
//     var number = secret.(int) * 10 // casting interface ke int
//     fmt.Println(secret, "multiplied by 10 is :", number)

//     secret = []string{"apple", "manggo", "banana"}
//     var gruits = strings.Join(secret.([]string), ", ")
//     fmt.Println(gruits, "is my favorite fruits")
// }

// Variabel bertipe interface{} bisa ditampilkan ke layar sebagai string dengan memanfaatkan fungsi print, seperti fmt.Println(). Tapi perlu diketahui bahwa nilai yang dimunculkan tersebut bukanlah nilai asli, melainkan bentuk string dari nilai aslinya.
// Oleh karena itu jika membutuhkan nilai asli dari interface harus di casting ke tipe aslinya.

// ==========================================================================

// // #3 Casting variabel interface kosong ke objek pointer

// package main

// import "fmt"

// type person struct {
//     name string
//     age  int
// }

// var secret interface{} = &person{name: "wick", age: 27}
// var name = secret.(*person).name // casting pointer
// fmt.Println(name)

