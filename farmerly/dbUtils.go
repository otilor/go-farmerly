package farmerly

import (
	"fmt"
	"net/url"
)

type Categories struct {
	Id         int
	Name       string
	uniqueHash string
}

func fetchCategories() (res []Categories) {
	db := databaseConn()
	fetchCategories, err := db.Query("SELECT * FROM categories")
	isError(err)

	categories := Categories{}
	var result []Categories
	for fetchCategories.Next() {
		var Id int
		var name, hash string
		err = fetchCategories.Scan(&Id, &name, &hash)

		isError(err)

		categories.Id = Id
		categories.Name = name
		categories.uniqueHash = hash
		result = append(result, categories)
	}
	return result
}

func verifyData(details url.Values) {
	if len(details["category"][0]) != 0 && len(details["username"][0]) != 0 {
		fmt.Println("Submitted successfully, your details are")
		fmt.Println(details)
	}

}