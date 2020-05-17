// Package farmerly implements a simple farming-algorithm.
// It maps farmers to content based on their selection

package farmerly

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
