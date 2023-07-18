package entity

type AuthUser struct {
	Email     string
	FirstName string
	LastName  string
	Salt      string
	Password  string
	Id        int64
}
