package main

import (
	"fmt"
	"math"
)
//SOAL KE 1
func updateLingkaran(luasLingkaran *float64, kelilingLingkaran *float64, jariJari float64) {
	*luasLingkaran = math.Pi * jariJari * jariJari
	*kelilingLingkaran = 2 * math.Pi * jariJari
}
//SOAL KE 2
func introduce(sentence *string, name string, gender string, profession string, age string) {
	if gender == "laki-laki" {
		*sentence = "Pak " + name + " adalah seorang " + profession + " yang berusia " + age + " tahun"
	} else {
		*sentence = "Bu " + name + " adalah seorang " + profession + " yang berusia " + age + " tahun"
	}
}
//SOAL KE 3
func addFruits(fruits *[]string) {
	*fruits = append(*fruits, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
}
//SOAL KE 4
func tambahDataFilm(title string, duration string, genre string, year string, dataFilm *[]map[string]string) {
	film := make(map[string]string)
	film["title"] = title
	film["duration"] = duration
	film["genre"] = genre
	film["year"] = year

	*dataFilm = append(*dataFilm, film)
}

func main() {
	// JAWABAN SOAL KE 1
	var luasLingkaran float64
	var kelilingLingkaran float64

	updateLingkaran(&luasLingkaran, &kelilingLingkaran, 6.0)

	fmt.Printf("Luas lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling lingkaran: %.2f\n", kelilingLingkaran)

	// JAWABAN SOAL KE 2
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	// JAWABAN SOAL KE 3
	var buah []string

	addFruits(&buah)

	for i, b := range buah {
		fmt.Printf("%d. %s\n", i+1, b)
	}

	// JAWABAN SOAL KE 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for _, film := range dataFilm {
		fmt.Printf("Title: %s\n", film["title"])
		fmt.Printf("Duration: %s\n", film["duration"])
		fmt.Printf("Genre: %s\n", film["genre"])
		fmt.Printf("Year: %s\n\n", film["year"])
	}
}
