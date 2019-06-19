package main
import "fmt"

func main() {
	// Slice adalah array yg tidak dituliskan elemennya
	var Hewan = [] string {"Ular", "Ayam", "Bebek", "Singa", "Harimai"}

	for i := 0; i < len(Hewan); i++ {
		fmt.Println(Hewan[i])
	}
	
	// ada 5 fungsi yg bisa digunakan dalam slice
	// len(), cap(), append(), copy()

	// fungsi cap() dan len()
	var aHewan = Hewan[0:3]
	fmt.Println(len(aHewan)) // nilai 3
	fmt.Println(cap(aHewan)) // nilai 5 karena dimulai dari index 0

	var bHewan = Hewan[2:5]
	fmt.Println(len(bHewan)) // nilai 3
	fmt.Println(cap(bHewan)) // nilai 3 karena dimulai dari index ke 2 sampai index 4

	// append() = menambahkan elemen di slice pada posisi paling akhir
	var cHewan = append(Hewan, "Semut Hitam")
	fmt.Println(cHewan)

	// copy() = men-copy elemen slice pada src ke dst
	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2
}