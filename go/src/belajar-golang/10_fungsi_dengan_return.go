package main

import ( // import dengan banyak package bisa di tuliskan dibawah. Bisa juga di tuliskan satu persatu dengan keyword import
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())
    var randomValue int

    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
}

func randomWithRange(min, max int) int { // setelah tutup parameter diikuti dengan tipe data nilai kembalian (return) = (parm, parm tipe_data_param) tipe_data_return
	// func nama_func(a int, b string) returnType = jika parameter memiliki tipe data yg berbeda
    var value = rand.Int() % (max - min + 1) + min
    return value
}