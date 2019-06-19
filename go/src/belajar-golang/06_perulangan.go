package main
import "fmt"

func main() {
	// Perulangan di Golang hanya for
	// Penulisan 1
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// Penulisan 2 (konsep seperti while)
	var j = 0
	for j < 5 {
		fmt.Println("Angka", j)
		j++
	}

	// Penulisan 3 tanpa kondisi, dihasilkan perulangan tanpa henti dan berhenti dengan break
	var k = 0

	for {
		fmt.Println("Angka", k)

		k++
		if k == 5 {
			break
		}
	}

	// Penulisan 4 = for - range u/ array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for l, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", l, fruit)
	}

	// Keyword break dan continue
	for m := 1; m <= 10; m++ {
		if m % 2 == 1 {
			continue
		}
	
		if m > 8 {
			break
		}
	
		fmt.Println("Angka", m)
	}
}