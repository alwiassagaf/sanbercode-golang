package main

import (
	"fmt"
	"sanbercode-golang-batch-43/library"
)

func main() {
	//jawaban Soal 1
	s := library.SegitigaSamaSisi{Alas: 4, Tinggi: 3}
	fmt.Println("Luas segitiga sama sisi:", s.Luas())
	fmt.Println("Keliling segitiga sama sisi:", s.Keliling())

	p := library.PersegiPanjang{Panjang: 5, Lebar: 3}
	fmt.Println("Luas persegi panjang:", p.Luas())
	fmt.Println("Keliling persegi panjang:", p.Keliling())

	t := library.Tabung{JariJari: 5, Tinggi: 10}
	fmt.Println("Volume tabung:", t.Volume())
	fmt.Println("Luas permukaan tabung:", t.LuasPermukaan())

	b := library.Balok{Panjang: 5, Lebar: 3, Tinggi: 2}
	fmt.Println("Volume balok:", b.Volume())
	fmt.Println("Luas permukaan balok:", b.LuasPermukaan())

	 //jawaban Soal 2
	 p1 := library.Phone{Name: "iPhone 12", Brand: "Apple", Year: 2020, Colors: []string{"Black", "White", "Red", "Green", "Blue"}}
	 p2 := library.Phone{Name: "Galaxy S21", Brand: "Samsung", Year: 2021, Colors: []string{"Phantom Gray", "Phantom White", "Phantom Violet", "Phantom Pink"}}
 
	 Phones := []library.PhoneInterface{p1, p2}
 
	 for _, p := range Phones {
		 fmt.Println(p.Display())
	 }

	 //jawaban Soal 3
	 fmt.Println(library.LuasPersegi(4, true))
	 fmt.Println(library.LuasPersegi(8, false))
	 fmt.Println(library.LuasPersegi(0, true))
	 fmt.Println(library.LuasPersegi(0, false))

	  //jawaban Soal 4
	  angka1 := []int{6, 8}
	  angka2 := []int{12, 14}
	  prefix := "hasil penjumlahan dari "
	  hasil :=library.Gabungkan(prefix, angka1, angka2)
	  fmt.Println(hasil)


}
