// Package farmerly implements a simple farming-algorithm.
// It maps farmers to content based on their selection

package farmerly

import (
	"html/template"
	"net/http"
)

func parseEmptyView(w http.ResponseWriter, tmpl string, r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	isError(err)

	err = t.Execute(w, nil)
	isError(err)
}

func parseView(w http.ResponseWriter, tmpl string, pageVars []Categories, r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	isError(err)

	err = t.Execute(w, pageVars)
	isError(err)
}

func ViewGen(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		uniqueHash, ok := r.URL.Query()["uniqueHash"]
		if ok{
			findUserWithHash(uniqueHash)
			parseEmptyView(w , "farmers.gohtml", r)
		}
		parseEmptyView(w, "generated_content.gohtml", r)
	} else {
	}

}

func findUserWithHash(hash []string) {
	uniqueHash := hash[0]
	db := databaseConn()
	findUserWithHash, err:= db.Query("SELECT users WHERE uniqueHash = ?", uniqueHash)
	isError(err)



}
