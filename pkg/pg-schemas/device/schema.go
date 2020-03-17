package profile

import (
	"database/sql"

	"github.com/GuyARoss/project-orva/pkg/pgdb"
)

type Request struct {
	Creds *pgdb.DbCreds
}

type Device struct {
	ID          string
	AccessLevel int
	Owner       string
	Name        string
	RegistedOn  int64
}

func (req *Request) FindByID(id string) (*Device, error) {
	db := pgdb.CreateSession(req.Creds)

	defer db.Close()

	sqlQuery := `select * from devices where ID = $1`
	row := db.QueryRow(sqlQuery, id)

	d := &Device{}
	if err := row.Scan(&d.ID, &d.AccessLevel, &d.Owner, &d.Name, &d.RegistedOn); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return d, nil
}

func (req *Request) Create(d *Device) error {
	db := pgdb.CreateSession(req.Creds)
	defer db.Close()

	sqlQuery := `insert into devices VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlQuery, d.ID, d.AccessLevel, d.Owner, d.Name, d.RegistedOn)

	return err
}
