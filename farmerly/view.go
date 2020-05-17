package farmerly


import (
	"html/template"
	"net/http"
)

func parseView(w http.ResponseWriter, tmpl string, r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		panic(err.Error())
	}

	err = t.Execute(w, nil)
	if err != nil {
		panic(err.Error())
	}
}


func ViewGen(w http.ResponseWriter, r *http.Request) {
	parseView(w, "generated_content.gohtml", r)
}
