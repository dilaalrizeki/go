package main
import "fmt"

func main() {
	// pendeklarasian array tipe ke-1 : langsung mengisi nilai pada indeks
	var names [4] string
	names[0] = "Dila"
	names[1] = "Rida"
	names[2] = "Al"
	names[3] = "Rizeki"

	fmt.Println(names[0],names[1],names[2],names[3])
	fmt.Println("==============");

	// pendeklarasian array tipe ke-2 : langsung diberikan nilai pada saat pendeklarasian

	var buah = [4] string {"apel", "jeruk", "alpukat", "melon"}
	fmt.Println("Jumlah element \t: ", len(buah))
	fmt.Println("Buah \t \t: ", buah)

	// pendeklarasian array tipe ke-3 : Gaya vertikal dan Horizontal

	var hewan[4] string
	hewan =[4] string{"Tikus", "Gajah", "Singa", "Kucing"} // Gaya Horizontal
	hewan =[4] string{ // Gaya Vertikal
		"Kucing",
		"Singa",
		"Gajah",
		"Tikus",
	}
}