package main

import (
	"github.com/gabielfemi/go-inform/farmerly"
	"log"
	"net/http"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", farmerly.Index)
	log.Println("Server started on: 127.0.0.1:8000")
	_ = http.ListenAndServe(":8000", nil)
}

