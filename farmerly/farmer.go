package farmerly

import (
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
	parseView(w, "farmers.html", r)
}

func Category(w http.ResponseWriter, r *http.Request) {
	parseView(w, "category.gohtml", r)
	if r.Method == "POST"{
		_ = r.ParseForm()
		log.Println("POST request")
		log.Println(r.Form)
	}
}

func ViewGen(w http.ResponseWriter, r *http.Request){
	parseView(w, "generated_content.html", r)
}

