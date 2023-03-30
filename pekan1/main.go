package main 

import (
    "fmt"
    "strings"
    "strconv"
)
	 

func main() {
    // SOAL KE 1
    word1, word2, word3, word4, word5 := "Bootcamp", "Digital", "Skill", "Sanbercode", "Golang"
    fmt.Println(word1, word2, word3, word4, word5)

    // SOAL KE 2
    hi := "hello word"
    hi = strings.Replace(hi, "word", "Golang", 1)
    fmt.Println(hi)

    // SOAL KE 3
    var kataPertama = "saya"
    var kataKedua = "senang"
    var kataKetiga = "belajar"
    var kataKeempat = "golang"

    kataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
    output := fmt.Sprintf("%s %s %s %s", kataPertama, kataKedua, kataKetiga, strings.ToUpper(kataKeempat))
    fmt.Println(output)

    // SOAL KE 4
    angkaPertama := "8"
    angkaKedua := "5"
    angkaKetiga := "6"
    angkaKeempat := "7"

    a, _ := strconv.Atoi(angkaPertama)
    b, _ := strconv.Atoi(angkaKedua)
    c, _ := strconv.Atoi(angkaKetiga)
    d, _ := strconv.Atoi(angkaKeempat)

    total := a + b + c + d
    fmt.Println(total)

    // SOAL KE 5
    kalimat := "halo halo bandung"
    angka := 2021
    kalimat = "Hi" + kalimat[2:8] + strings.ToUpper(string(kalimat[8])) + kalimat[9:]
    hasil = fmt.Sprintf("%s - %d", kalimat, angka)
    fmt.Println(hasil)
}
