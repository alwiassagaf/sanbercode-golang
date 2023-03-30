package main

import "fmt"

// SOAL KE 1

func luasPersegiPanjang(panjang float64, lebar float64) float64 {
    return panjang * lebar
}

func kelilingPersegiPanjang(panjang float64, lebar float64) float64 {
    return 2 * (panjang + lebar)
}

func volumeBalok(panjang float64, lebar float64, tinggi float64) float64 {
    return panjang * lebar * tinggi
}

// SOAL KE 2
func introduce(name, gender, profession, age string) string {
    var salutation string
    if gender == "laki-laki" {
        salutation = "Pak "
    } else {
        salutation = "Bu "
    }
    return fmt.Sprintf("%s%s adalah seorang %s yang berusia %s tahun", salutation, name, profession, age)
}

// SOAL KE 3
func buahFavorit(nama string, buah ...string) string {
    return fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %q", nama, buah)
}



func main() {
	// hasil soal 1
    panjang := 12.0
    lebar := 4.0
    tinggi := 8.0

    luas := luasPersegiPanjang(panjang, lebar)
    keliling := kelilingPersegiPanjang(panjang, lebar)
    volume := volumeBalok(panjang, lebar, tinggi)

    fmt.Println(luas)
    fmt.Println(keliling)
    fmt.Println(volume)

    // hasil soal ke 2
    john := introduce("John", "laki-laki", "penulis", "30")
    fmt.Println(john)

    sarah := introduce("Sarah", "perempuan", "model", "28")
    fmt.Println(sarah)

	// hasil soal ke 3
	buah := []string{"semangka", "jeruk", "melon", "pepaya"}
    buahFavoritJohn := buahFavorit("John", buah...)
    fmt.Println(buahFavoritJohn)

// SOAL KE 4
	var dataFilm = []map[string]string{}

    tambahDataFilm := func(judul string, durasi string, genre string, tahun string) {
        film := map[string]string{
            "judul":  judul,
            "durasi": durasi,
            "genre":  genre,
            "tahun":  tahun,
        }
        dataFilm = append(dataFilm, film)
}
// hasil soal ke 4
tambahDataFilm("LOTR", "2 jam", "action", "1999")
tambahDataFilm("avenger", "2 jam", "action", "2019")
tambahDataFilm("spiderman", "2 jam", "action", "2004")
tambahDataFilm("juon", "2 jam", "horror", "2004")

for _, item := range dataFilm {
	fmt.Println(item)
}
}
