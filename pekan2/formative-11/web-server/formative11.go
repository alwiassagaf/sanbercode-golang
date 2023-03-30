package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jariJari := 7.0
		tinggi := 10.0

		volume := math.Pi * math.Pow(jariJari, 2) * tinggi
		luasAlas := math.Pi * math.Pow(jariJari, 2)
		kelilingAlas := 2 * math.Pi * jariJari

		fmt.Fprintf(w, "jariJari: %v, tinggi: %v, volume: %v, luas alas: %v, keliling alas: %v",
			jariJari, tinggi, volume, luasAlas, kelilingAlas)
	})

	fmt.Println("Starting server at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
