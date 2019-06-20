package main
import "fmt"

func main() {
	// Ada 3 Jenis operator = Aritmatika, Perbandingan dan Logika
	// #1 Aritmatika ( + - * / % )
	var nilai = (((6 + 6) % 3) + 11 - 2) / 3
	fmt.Println(nilai)
	// #2 Perbandingan (==, !=, <, >, <=, >=
	var isEqual = (nilai == 2)
	fmt.Println("apakah",nilai, "== 2 : ", isEqual)
	// #3 Logika (||, &&, !)
	a,b := true, false
	operatorDan := a && b
	fmt.Printf("a && b = (%t)  \n", operatorDan)
}
