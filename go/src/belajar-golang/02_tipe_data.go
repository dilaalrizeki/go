package main
import "fmt"

func main() {
	// Tipe data ada 3 jenis = numerik (desimal & non-desimal), string, dan boolean.

	// Numerik desimal = float32 dan float64
	// Numerik non-desimal = uint (8,16,32,64), byte, int (8,16,32,64), rune

	// String ditandai dengan diapit oleh ""
	// bool berisi true dan false

	var positiveNumber uint8 = 89
	var decimalNumber = 2.62
	var benar bool = true
	var pesan string = "Halo"

	fmt.Println(positiveNumber, decimalNumber, benar, pesan)

}