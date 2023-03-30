package main

import (
	"fmt"
	"strings"
)

func main() {
	// SOAL KE 1
	for i := 1; i <= 20; i++ {
        if i % 2 == 0 {
            fmt.Printf("%d - Berkualitas\n", i)
        } else if i % 3 == 0 && i % 2 == 1 {
            fmt.Printf("%d - I Love Coding\n", i)
        } else {
            fmt.Printf("%d - Santai\n", i)
        }
    }

// SOAL KE 2
for i := 1; i <= 7; i++ {
	fmt.Println("#" + strings.Repeat("#", i-1))
}

// SOAL KE 3
kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	var hasil []string

	for _, kata := range kalimat {
		if kata != "aku" && kata != "dan" {
			hasil = append(hasil, kata)
		}
	}

	fmt.Println(hasil)

	// SOAL KE 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, v := range sayuran {
		fmt.Printf("%d. %s\n", i+1, v)
	}

	// SOAL KE 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	fmt.Println("Ukuran kotak:")
	for k, v := range satuan {
		fmt.Printf("%s %d ", k, v)
	}

}