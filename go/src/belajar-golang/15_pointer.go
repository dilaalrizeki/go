// Pointer adalah alamat memory
package main

import "fmt"

func main() {
    // Penulisan pointer ditandai dengan *var number *int
    // var name *string
    // var number = 4
    // variabel pointer tidak bisa menampung nilai yg bukan pointer

    var numberA int = 4
    var numberB *int = &numberA // tanda "&" u/ mengambil nilai pointer dari variabel biasa yg biasa disebut dengan referencing

    fmt.Println("numberA (value)   :", numberA)  // 4
    fmt.Println("numberA (address) :", &numberA) // 0xc04200e210

    fmt.Println("numberB (value)   :", *numberB) // 4
    fmt.Println("numberB (address) :", numberB)  // 0xc04200e210

}