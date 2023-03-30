package main

import (
	"errors"
	"fmt"
	"time"
	"sort"
	"sync"
)

//soal 1
func tampilkanKalimat(kalimat string, tahun int) {
	defer fmt.Println("Terima kasih telah menggunakan aplikasi ini")

	fmt.Println(kalimat)
	fmt.Println(tahun)
	fmt.Println("Silahkan mulai menggunakan aplikasi ini")
}

//soal 2
func kelilingSegitigaSamaSisi(sisi int, tampilkanKalimat bool) (string, error) {
	if sisi == 0 {
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		if tampilkanKalimat {
			return "", err
		} else {
			// panic(err)
		}
	}

	keliling := sisi * 3

	if tampilkanKalimat {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	} else {
		return fmt.Sprintf("%d", keliling), nil
	}
}

// soal 3
func tambahAngka(nilai int, total *int) {
	*total += nilai
}

func cetakAngka(total *int) {
	fmt.Println("Total angka:", *total)
}

// soal 4
var phones = []string{}

func addPhone(p *[]string, phone string) {
	*p = append(*p, phone)
}

// soal 5
var phone = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

// soal 6
var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

func getMovies(moviesChannel chan string, movies ...string) {
    for _, movie := range movies {
        moviesChannel <- movie
    }
    close(moviesChannel)
}


func main() {
	// jawaban soal ke 1
	kalimat := "Golang Backend Development"
	tahun := 2021
	tampilkanKalimat(kalimat, tahun)

	// jawaban soal ke 2
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()

	hasil, err := kelilingSegitigaSamaSisi(4, true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(hasil)
	}

	hasil, err = kelilingSegitigaSamaSisi(8, false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(hasil)
	}

	hasil, err = kelilingSegitigaSamaSisi(0, true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(hasil)
	}

	hasil, err = kelilingSegitigaSamaSisi(0, false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(hasil)
	}

	// jawaban soal ke 3
	angka := 1
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	defer fmt.Println("Nilai angka terakhir:", angka)

	// jawaban soal ke 4
	fmt.Println("Golang Backend Development")
	fmt.Println("2021")
	fmt.Println("Silahkan mulai menggunakan aplikasi ini")

	addPhone(&phones, "Xiaomi")
	addPhone(&phones, "Asus")
	addPhone(&phones, "IPhone")
	addPhone(&phones, "Samsung")
	addPhone(&phones, "Oppo")
	addPhone(&phones, "Realme")
	addPhone(&phones, "Vivo")

	sort.Strings(phones)

	fmt.Println("Daftar telepon yang tersedia:")
	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}

	// jawaban soal ke 5
	wg := sync.WaitGroup{}
	wg.Add(len(phone))

	for i, phone := range phone {
		go func(i int, phone string) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("%d. %s\n", i+1, phone)
		}(i, phone)
	}

	wg.Wait()
	fmt.Println("Selesai menampilkan data")

	// jawaban soal ke 6
	moviesChannel := make(chan string)
    go getMovies(moviesChannel, movies...)
    for value := range moviesChannel {
        fmt.Println(value)
    }
}
