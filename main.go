package main

import (
	"github.com/gabielfemi/go-inform/farmerly"
	"net/http"
)

func main() {
	http.HandleFunc("/", farmerly.Index)
}

