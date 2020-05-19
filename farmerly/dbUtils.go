// Package farmerly implements a simple farming-algorithm.
// It maps farmers to content based on their selection

package farmerly



func fetchCategories() (res []Categories) {
	db := databaseConn()
	fetchCategories, err := db.Query("SELECT * FROM categories")
	isError(err)

	categories := Categories{}
	var result []Categories
	for fetchCategories.Next() {
		var Id int
		var name string
		err = fetchCategories.Scan(&Id, &name)

		isError(err)

		categories.Id = Id
		categories.Name = name
		result = append(result, categories)
	}
	return result
}

func addUser(name string, category string) {
	db := databaseConn()
	userHash := generateUserHash()
	addUser, err := db.Prepare("INSERT INTO users(name, category_name, userHash) VALUES (?, ?, ?)")
	isError(err)
	addUser.Exec(name, category, userHash)
}
