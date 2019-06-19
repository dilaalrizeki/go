package main

import "fmt"
import "strings"

// Method adalah fungsi yang menempel pada struct, sehingga hanya bisa di akses lewat variabel objek.

type student struct {
    name  string
    grade int
}

func (s student) sayHello() {
    fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
    return strings.Split(s.name, " ")[i-1]
}

func main() {
    var s1 = student{"john wick", 21}
    s1.sayHello() // mengakses method

    var name = s1.getNameAt(2)
    fmt.Println("nama panggilan :", name)
}