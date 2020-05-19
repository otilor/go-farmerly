// Package farmerly implements a simple farming-algorithm.
// It maps farmers to content based on their selection

package farmerly

func sortAccordingToCategory(category string) {
	db := databaseConn()
	getPostsAccordingToCategory, err := db.Query("SELECT * FROM posts WHERE category = ?",category)
	isError(err)


	for  getPostsAccordingToCategory.Next(){

	}

}