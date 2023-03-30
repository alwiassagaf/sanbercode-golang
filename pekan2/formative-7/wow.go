package main

import "fmt"
// soal ke 1
type buah struct {
    nama       string
    warna      string
    adaBijinya bool
    harga      int
}
// soal ke 2
type segitiga struct {
    alas   int
    tinggi int
}

type persegi struct {
    sisi int
}

type persegiPanjang struct {
    panjang int
    lebar   int
}
// soal ke 3
type phone struct {
    name   string
    brand  string
    year   int
    colors []string
}
// soal ke 4
type movie struct {
    title    string
    genre    string
    duration int
    year     int
}



// ke 2
func (s segitiga) luas() float64 {
	return 0.5 * float64(s.alas) * float64(s.tinggi)
}

func (p persegi) luas() float64 {
	return float64(p.sisi) * float64(p.sisi)
}

func (pp persegiPanjang) luas() float64 {
	return float64(pp.panjang) * float64(pp.lebar)
}
// ke 3
func (p *phone) tambahWarna(color string) {
	p.colors = append(p.colors, color)
}
// ke 4
var dataFilm = []movie{}
func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
    film := movie{title, genre, duration, year}
    *dataFilm = append(*dataFilm, film)
}

func main() {
    //jawaban Soal 1
    buah1 := buah{"Nanas", "Kuning", false, 9000}
    buah2 := buah{"Jeruk", "Oranye", true, 8000}
    buah3 := buah{"Semangka", "Hijau & Merah", true, 10000}
    buah4 := buah{"Pisang", "Kuning", false, 5000}

    fmt.Println(buah1)
    fmt.Println(buah2)
    fmt.Println(buah3)
    fmt.Println(buah4)

    //jawaban Soal 2
    segitiga1 := segitiga{4, 6}
    persegi1 := persegi{8}
    persegiPanjang1 := persegiPanjang{6, 10}

    fmt.Println("Luas segitiga:", segitiga1.luas())
    fmt.Println("Luas persegi:", persegi1.luas())
    fmt.Println("Luas persegi panjang:", persegiPanjang1.luas())

    //jawaban Soal 3
    p1 := phone{
        name:   "iPhone",
        brand:  "Apple",
        year:   2021,
        colors: []string{"Silver", "Gold", "Graphite"},
    }
    p1.tambahWarna("Blue")
    fmt.Println(p1)

    // jawaban Soal 4
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
    tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
    tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
    tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for _, film := range dataFilm {
		fmt.Printf("%v (%v)\n", film.title, film.genre)
		fmt.Printf("Genre: %v\n", film.duration)
		fmt.Printf("Duration: %v minutes\n", film.duration)
		fmt.Printf("Year: %v\n\n", film.year)
	}
}
