package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Selamat datang di halaman utama kami"))
}

func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selamat datang di halaman tentang kami"))
}

func Product(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product page : %d", idNumb)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/about", About)
	mux.HandleFunc("/product", Product)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
