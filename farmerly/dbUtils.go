// Package farmerly implements a simple farming-algorithm.
// It maps farmers to content based on their selection

package farmerly

func FetchCategories() (res []Categories) {
	db := databaseConn()
	FetchCategories, err := db.Query("SELECT * FROM categories ORDER by id")
	isError(err)

	categories := Categories{}
	var result []Categories
	for FetchCategories.Next() {
		var Id int
		var name string
		err = FetchCategories.Scan(&Id, &name)

		isError(err)

		categories.Id = Id
		categories.Name = name
		result = append(result, categories)
	}
	return result
}
func FetchUsers() (res []Users) {
	db := databaseConn()
	FetchUsers, err := db.Query("SELECT * FROM users ORDER by id")
	isError(err)

	users := Users{}
	var result []Users
	for FetchUsers.Next() {
		var Id int
		var name string
		var category_name string
		var userHash string
		err = FetchUsers.Scan(&Id, &name, &category_name, &userHash)

		isError(err)

		users.Id = Id
		users.Name = name
		users.Name = name
		users.Category = category_name
		users.userHash = userHash
		result = append(result, users)
	}
	return result
}

func GetPosts() (res []Posts){
	db := databaseConn()
	FetchUsers, err := db.Query("SELECT * FROM posts ORDER by id")
	isError(err)

	posts := Posts{}
	var result []Posts
	for FetchUsers.Next() {
		var Id int
		var title string
		var body string
		var category string
		err = FetchUsers.Scan(&Id, &title, &body, &category)

		isError(err)

		posts.Id = Id
		posts.Title = title
		posts.Body = body
		posts.Category = category
		result = append(result, posts)
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
