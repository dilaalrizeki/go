package main
import "fmt"

func main() {
	var ultah map[string]string
	ultah = map[string]string{}
	// bisa langsung = var ultah = map[string]string{}

	ultah["januari"] = "Dilarida Alrizeki"
	ultah["September"] = "Naftaline Aditya Khoirunnisya"

	fmt.Println("Ulang tahun pada bulan januari adalah", ultah["januari"])

	fmt.Println("======================================================")

	// inisialisasi map secara langsung dengan metode vertikal dan horizontal
	// cara vertikal
	var hero = map[int]string{1:"Ringo", 2:"Alpha", 3:"Joule", 4:"Glaive"}
	fmt.Println(hero[2])
	fmt.Println("======================================================")

	// cara horizontal
	var hero2 = map[int]string{
		1:"Ringo",
		2:"Alpha",
		3:"Joule",
		4:"Glaive",
	}
	fmt.Println(hero2[1])
	fmt.Println("======================================================")

	// inisialisasi map ada 3 
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int) = menghasilkan data pointer, pengambilannya menggunakna *


	// Iterasi map menggunakan for - range
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	
	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}
	fmt.Println("Banyak data", len(chicken))
	fmt.Println("======================================================")

	// menghapus item di map menggunakan delete(), key sensitive
	delete(chicken, "januari") // key januari dihapus
	fmt.Println(chicken)
	fmt.Println("Banyak data",len(chicken))
}