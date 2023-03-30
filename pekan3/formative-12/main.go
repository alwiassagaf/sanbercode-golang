package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID uint
}

var nilaiMahasiswa = []NilaiMahasiswa{}

func main() {
	http.HandleFunc("/api/nilai-mahasiswa", nilaiMahasiswaHandler)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func nilaiMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getNilaiMahasiswaHandler(w, r)
	case "POST":
		postNilaiMahasiswaHandler(w, r)
	default:
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
	}
}

func getNilaiMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(nilaiMahasiswa) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	jsonData, err := json.Marshal(nilaiMahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(jsonData))
}

func postNilaiMahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || username != "admin" || password != "password" {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var nilaiMahasiswaBaru NilaiMahasiswa
	err := json.NewDecoder(r.Body).Decode(&nilaiMahasiswaBaru)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if nilaiMahasiswaBaru.Nilai > 100 {
		http.Error(w, "Nilai should be at most 100.", http.StatusBadRequest)
		return
	}

	if nilaiMahasiswaBaru.Nilai >= 80 {
		nilaiMahasiswaBaru.IndeksNilai = "A"
	} else if nilaiMahasiswaBaru.Nilai >= 70 {
		nilaiMahasiswaBaru.IndeksNilai = "B"
	} else if nilaiMahasiswaBaru.Nilai >= 60 {
		nilaiMahasiswaBaru.IndeksNilai = "C"
	} else if nilaiMahasiswaBaru.Nilai >= 50 {
		nilaiMahasiswaBaru.IndeksNilai = "D"
	} else {
		nilaiMahasiswaBaru.IndeksNilai = "E"
	}

	nilaiMahasiswaBaru.ID = uint(len(nilaiMahasiswa) + 1)

	nilaiMahasiswa = append(nilaiMahasiswa, nilaiMahasiswaBaru)

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(nilaiMahasiswaBaru)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonData))
}
