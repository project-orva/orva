package profile

import (
	"database/sql"

	"github.com/GuyARoss/project-orva/pkg/pgdb"
)

type Request struct {
	Creds *pgdb.DbCreds
}

type User struct {
	ID          string
	AccessLevel string
	FirstName   string
	LastName    string
}

func (req *Request) FindByID(id string) (*User, error) {
	db := pgdb.CreateSession(req.Creds)

	defer db.Close()

	sqlQuery := `select * from users where ID = $1`
	row := db.QueryRow(sqlQuery, id)

	u := &User{}
	if err := row.Scan(&u.ID, &u.AccessLevel); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return u, nil
}

func (req *Request) Create(u *User) error {
	db := pgdb.CreateSession(req.Creds)
	defer db.Close()

	sqlQuery := `insert into users VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlQuery, u.ID, u.AccessLevel, u.FirstName, u.LastName)

	return err
}
