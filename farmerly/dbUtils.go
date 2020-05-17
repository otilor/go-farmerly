package farmerly
type Categories struct {
	id int
	name string
}

func fetchCategories() (res []Categories) {
	db := databaseConn()
	fetchCategories, err := db.Query("SELECT * FROM categories")
	isError(err)


	categories := Categories{}
	var result []Categories
	for fetchCategories.Next() {
		var id int
		var name string
		err = fetchCategories.Scan(&id, &name)

		isError(err)

		categories.id = id
		categories.name = name
		result = append(result, categories)
	}
	return result
}