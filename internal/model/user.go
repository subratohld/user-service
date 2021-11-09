package model

type User struct {
	Id        int64  `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Mobile    string `db:"mobile" json:"mobile"`
	FirstName string `db:"firstName" json:"firstName"`
	LastName  string `db:"lastName" json:"lastName"`
	Address   string `db:"address" json:"address"`
	Password  string `db:"password" json:"password"`
}

func (u User) Validate() error {
	return nil
}
