package main

import (
	"fmt"
	"math"
)

// Soal 1
type segitigaSamaSisi struct{
    alas, tinggi int
}

type persegiPanjang struct{
    panjang, lebar int
}

type tabung struct{
    jariJari, tinggi float64
}

type balok struct{
    panjang, lebar, tinggi int
}

type bangunDatar interface{
    luas() int
    keliling() int
}

type bangunRuang interface{
    volume() float64
    luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
    return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
    return 3 * s.alas
}

func (p persegiPanjang) luas() int {
    return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
    return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
    return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
    return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
    return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
    return 2 * (float64(b.panjang*b.lebar) + float64(b.panjang*b.tinggi) + float64(b.lebar*b.tinggi))
}

// Soal 2
type phone struct {
    name    string
    brand   string
    year    int
    colors  []string
}

type phoneInterface interface {
    display() string
}

func (p phone) display() string {
    return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: %v\n", p.name, p.brand, p.year, p.colors)
}

// Soal 3
func luasPersegi(sisi int, showText bool) interface{} {
    if sisi == 0 {
        if showText {
            return "Maaf anda belum menginput sisi dari persegi"
        } else {
            return nil
        }
    }
    luas := sisi * sisi
    if showText {
        return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
    } else {
        return luas
    }
}

// Soal 4
func gabungkan(prefix, angka1, angka2 interface{}) string {
    total := 0
    if p, ok := prefix.(string); ok {
        result := p
        if a1, ok := angka1.([]int); ok {
            for _, num := range a1 {
                total += num
                result += fmt.Sprintf("%d + ", num)
            }
        }
        if a2, ok := angka2.([]int); ok {
            for _, num := range a2 {
                total += num
                result += fmt.Sprintf("%d + ", num)
            }
        }
        result = result[:len(result)-3] // Menghapus tanda "+" dan spasi terakhir
        result += fmt.Sprintf("= %d", total)
        return result
    }
    return ""
}

func main() {
    //jawaban Soal 1
	s := segitigaSamaSisi{alas: 4, tinggi: 3}
    fmt.Println("Luas segitiga sama sisi:", s.luas())
    fmt.Println("Keliling segitiga sama sisi:", s.keliling())

    p := persegiPanjang{panjang: 5, lebar: 2}
    fmt.Println("Luas persegi panjang:", p.luas())
    fmt.Println("Keliling persegi panjang:", p.keliling())

    t := tabung{jariJari: 3.5, tinggi: 7}
    fmt.Println("Volume tabung:", t.volume())
    fmt.Println("Luas permukaan tabung:", t.luasPermukaan())

    b := balok{panjang: 6, lebar: 4, tinggi: 5}
    fmt.Println("Volume balok:", b.volume())
    fmt.Println("Luas permukaan balok:", b.luasPermukaan())

    //jawaban Soal 2
	p1 := phone{name: "iPhone 12", brand: "Apple", year: 2020, colors: []string{"Black", "White", "Red", "Green", "Blue"}}
    p2 := phone{name: "Galaxy S21", brand: "Samsung", year: 2021, colors: []string{"Phantom Gray", "Phantom White", "Phantom Violet", "Phantom Pink"}}
    
    phones := []phoneInterface{p1, p2}
    
    for _, p := range phones {
        fmt.Println(p.display())
    }

    //jawaban Soal 3
    fmt.Println(luasPersegi(4, true))
    fmt.Println(luasPersegi(8, false))
    fmt.Println(luasPersegi(0, true))
    fmt.Println(luasPersegi(0, false))

    //jawaban Soal 4
    prefix := "hasil penjumlahan dari "
    angka1 := []int{6, 8}
    angka2 := []int{12, 14}
    fmt.Println(gabungkan(prefix, angka1, angka2))
}
