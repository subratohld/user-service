package repository

import (
	db "github.com/subratohld/sqldb"
	"github.com/subratohld/user-service/internal/model"
)

type User interface {
	Save(user *model.User) (int64, error)
}

func NewUserRepo(db db.Sql) User {
	return &user{
		db: db,
	}
}

type user struct {
	db db.Sql
}

func (usr *user) Save(user *model.User) (int64, error) {
	query := `INSERT INTO tbl_user(EMAIL,MOBILE,FIRST_NAME,LAST_NAME,ADDRESS,PASSWORD)
				values(:email, :mobile, :firstName, :lastName, :address, :password);`
	res, err := usr.db.NamedExec(query, user)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
