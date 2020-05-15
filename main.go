package main

import (
	"github.com/gabielfemi/go-inform/farmerly"
	"log"
	"net/http"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", farmerly.Index)
	http.HandleFunc("/category", farmerly.Category)
	http.HandleFunc("/show", farmerly.ViewGen)
	log.Println("Server started on: 127.0.0.1:8000")
	_ = http.ListenAndServe(":8000", nil)
}

