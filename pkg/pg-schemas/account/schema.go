package account

import (
	"database/sql"
	"github.com/GuyARoss/project-orva/pkg/pgdb"
)

type Request struct {
	Creds *pgdb.DbCreds
}

type Account struct {
	Type int32
	AccessLevel int32
	Name string
	ID string
}

func (req *Request) FindByID(id string) (*Account, error) {
	db := pgdb.CreateSession(req.Creds)

	defer db.Close()

	sqlQuery := `select * from account where id = $1`
	row := db.QueryRow(sqlQuery, id)

	a := &Account{}
	if err := row.Scan(&a.ID, &a.AccessLevel, &a.Name, &a.Type); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return a, nil
}