package main

import "fmt"
import "runtime"

// Goroutine bisa disebut sebagai mini thread

func print(till int, message string) {
    for i := 0; i < till; i++ {
        fmt.Println((i + 1), message)
    }
}

func main() {
    runtime.GOMAXPROCS(2)

    go print(5, "halo") // u/ pemanggilan goroutine, harus diawali keyword "go" sebelum fungsinya
    print(5, "apa kabar")

    var input string
    fmt.Scanln(&input) // u/ blocking proses
}