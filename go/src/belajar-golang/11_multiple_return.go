package main

import "fmt"
import "math"

func main() {
    var diameter float64 = 15
    var area, circumference = calculate(diameter)

    fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
    fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

func calculate(d float64) (float64, float64) {
    // hitung luas
    var area = math.Pi * math.Pow(d / 2, 2)
    // hitung keliling
    var circumference = math.Pi * d

    // kembalikan 2 nilai
    return area, circumference
}

// Predined return value
// func calculate(d float64) (area float64, circumference float64) {}
// variabel nilai balik sudah ditentukan di awal (area dan circumference)
// jadi setelah return tidak perlu lagi ditambahkan variabel nilai baliknya
// nilai balik didapat dari kode diatas return