package profile

import (
	"database/sql"

	"github.com/GuyARoss/project-orva/pkg/pgdb"
)

type Request struct {
	Creds *pgdb.DbCreds
}

type Profile struct {
	AccountID    string
	BlockAddress string
	CreatedOn    int64
}

func (req *Request) FindByID(id string) (*Profile, error) {
	db := pgdb.CreateSession(req.Creds)

	defer db.Close()

	sqlQuery := `select * from profile where accountID = $1`
	row := db.QueryRow(sqlQuery, id)

	p := &Profile{}
	if err := row.Scan(&p.AccountID, &p.BlockAddress, &p.CreatedOn); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return p, nil
}

func (req *Request) Create(p *Profile) error {
	db := pgdb.CreateSession(req.Creds)
	defer db.Close()

	sqlQuery := `insert into profile VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlQuery, p.AccountID, p.BlockAddress, p.CreatedOn)

	return err
}
