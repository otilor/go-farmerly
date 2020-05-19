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

func parseView(w http.ResponseWriter, tmpl string, pageVars interface{}, r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	isError(err)

	err = t.Execute(w, pageVars)
	isError(err)
}

func ViewGen(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		uniqueHash, ok := r.URL.Query()["uniqueHash"]
		if ok{
			user := findUserWithHash(uniqueHash)

			parseView(w, "farmers.gohtml", user, r)
		}else {
			parseEmptyView(w, "generated_content.gohtml", r)
		}

	} else {
	}

}

func findUserWithHash(hash []string) []UserFromDatabase{
	uniqueHash := hash[0]
	db := databaseConn()
	findUserWithHash, err:= db.Query("SELECT * FROM users WHERE userHash = ?", uniqueHash)
	isError(err)

	usr := UserFromDatabase{}
	var res []UserFromDatabase
	 for findUserWithHash.Next() {
	 	var id int
	 	var name, category, hash string
	 	err = findUserWithHash.Scan(&id, &name, &category, &hash)
	 	isError(err)

	 	usr.Id = id
	 	usr.Name = name
	 	usr.Category = category
	 	usr.UniqueHash = hash
	 }
	 res = append(res, usr)
	return res

}
