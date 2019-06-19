package main
import "fmt"

// Deklarasi struct
type student struct {
    name string
    grade int
}
// Keyword "type" digunakan u/ pendeklarasian struct

func main() {
    // membuat variabel objek
    var s1 student // var <nama_variabel> <nama_struct>

    s1.name = "john wick"
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("grade :", s1.grade)
}