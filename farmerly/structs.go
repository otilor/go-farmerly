package farmerly

type UserFromDatabase struct {
	id         int
	name       string
	category   string
	uniqueHash string
}

type Categories struct {
	Id         int
	Name       string
	uniqueHash string
}
