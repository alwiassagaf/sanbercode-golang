package main

import (
	"fmt"
	"strconv"
)

func main() {
	// SOAL KE 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjang * lebar
	var kelilingPersegiPanjang int = 2 * (panjang + lebar)
	var luasSegitiga int = alas * tinggi / 2

	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)

// SOAL KE 2
	var nilaiJohn = 80
	var nilaiDoe = 50
	var indeksJohn string
	var indeksDoe string

	if nilaiJohn >= 80 {
		indeksJohn = "A"
	} else if nilaiJohn >= 70 {
		indeksJohn = "B"
	} else if nilaiJohn >= 60 {
		indeksJohn = "C"
	} else if nilaiJohn >= 50 {
		indeksJohn = "D"
	} else {
		indeksJohn = "E"
	}

	if nilaiDoe >= 80 {
		indeksDoe = "A"
	} else if nilaiDoe >= 70 {
		indeksDoe = "B"
	} else if nilaiDoe >= 60 {
		indeksDoe = "C"
	} else if nilaiDoe >= 50 {
		indeksDoe = "D"
	} else {
		indeksDoe = "E"
	}

	fmt.Println("Indeks nilai John: ", indeksJohn)
	fmt.Println("Indeks nilai Doe: ", indeksDoe)

	// SOAL KE 3
	var tanggal = 14
	var bulan = 3
	var tahun = 2001

	bulanString := []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	fmt.Printf("%d %s %d\n", tanggal, bulanString[bulan], tahun)

	// SOAL KE 4
	var tahunKelahiran = 2001

    if tahunKelahiran >= 1995 && tahunKelahiran <= 2015 {
        fmt.Println("Saya termasuk dalam Generasi Z")
    } else if tahunKelahiran >= 1980 && tahunKelahiran <= 1994 {
        fmt.Println("Saya termasuk dalam Generasi Y (Millennials)")
    } else if tahunKelahiran >= 1965 && tahunKelahiran <= 1979 {
        fmt.Println("Saya termasuk dalam Generasi X")
    } else if tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
        fmt.Println("Saya termasuk dalam Baby Boomer")
    } else {
        fmt.Println("Tahun kelahiran tidak valid")
	}
}
