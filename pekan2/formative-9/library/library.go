package library

import (
	"math"
    "fmt"      
)
// soal ke 1
type SegitigaSamaSisi struct {
    Alas, Tinggi int
}

type PersegiPanjang struct {
    Panjang, Lebar int
}

type Tabung struct {
    JariJari, Tinggi float64
}

type Balok struct {
    Panjang, Lebar, Tinggi int
}

type BangunDatar interface {
    Luas() float64
    Keliling() float64
}

type BangunRuang interface {
    Volume() float64
    LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() float64 {
    return float64(s.Alas * s.Tinggi / 2)
}

func (s SegitigaSamaSisi) Keliling() float64 {
    return float64(3 * s.Alas)
}

func (p PersegiPanjang) Luas() float64 {
    return float64(p.Panjang * p.Lebar)
}

func (p PersegiPanjang) Keliling() float64 {
    return float64(2 * (p.Panjang + p.Lebar))
}

func (t Tabung) Volume() float64 {
    return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
    return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
    return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
    return 2 * (float64(b.Panjang*b.Lebar) + float64(b.Panjang*b.Tinggi) + float64(b.Lebar*b.Tinggi))
}

// soal ke 2
type Phone struct {
    Name    string
    Brand   string
    Year    int
    Colors  []string
}

type PhoneInterface interface {
    Display() string
}

func (p Phone) Display() string {
    return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: %v\n", p.Name, p.Brand, p.Year, p.Colors)
}

// Soal 3
func LuasPersegi(Sisi int, ShowText bool) interface{} {
    if Sisi == 0 {
        if ShowText {
            return "Maaf anda belum menginput sisi dari persegi"
        } else {
            return nil
        }
    }
    luas := Sisi * Sisi
    if ShowText {
        return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", Sisi, luas)
    } else {
        return luas
    }

}

// Soal 4
func Gabungkan(prefix, angka1, angka2 interface{}) string {
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
        result = result[:len(result)-3] 
        result += fmt.Sprintf(" = %d", total)
        return result
    }
    return ""
}

