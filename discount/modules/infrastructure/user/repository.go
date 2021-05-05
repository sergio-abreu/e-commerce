package user

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/discount/modules/domain/user"
)

var ErrUserNotFound = errors.New("user not found")

type Repository struct {
	db *sql.DB
}

func (r *Repository) GetById(id string) (user.User, error) {
	query := "SELECT id, date_of_birth FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)
	var aUser user.User
	err := row.Scan(&aUser.Id, &aUser.DateOfBirth)
	if err == sql.ErrNoRows {
		return aUser, ErrUserNotFound
	}
	if err != nil {
		return aUser, errors.Wrap(err, "failed to query user database")
	}
	return aUser, nil
}
