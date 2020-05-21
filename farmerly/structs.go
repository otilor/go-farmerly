package farmerly

type UserFromDatabase struct {
	Id         int
	Name       string
	Category   string
	UniqueHash string
}

type Categories struct {
	Id         int
	Name       string
	uniqueHash string
}
type Users struct {
	Id int
	Name string
	Category string
	userHash string
}
type Posts struct {
	Id int
	Title string
	Body string
	Category string
}