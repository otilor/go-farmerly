// Package farmerly implements a simple farming-algorithm.
// It maps farmers to content based on their selection

package farmerly

func sortAccordingToCategory(category string) Posts{
	db := databaseConn()
	getPostsAccordingToCategory, err := db.Query("SELECT * FROM posts WHERE category = ?",category)
	isError(err)

	posts := Posts{}
	for  getPostsAccordingToCategory.Next(){
		var id int
		var title, body, category string
		err := getPostsAccordingToCategory.Scan(&id, &title, &body, &category)
		isError(err)
		posts.Id = id
		posts.Title = title
		posts.Body = body
		posts.Category = category
	}
	return posts
}