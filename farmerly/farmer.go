package farmerly

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
	parseView(w, "farmers.html", r)
}

func Category(w http.ResponseWriter, r *http.Request) {
	parseView(w, "category.html", r)
}

func ViewGen(w http.ResponseWriter, r *http.Request){
	parseView(w, "generated_content.html", r)
}

