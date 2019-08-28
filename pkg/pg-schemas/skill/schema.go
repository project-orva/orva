package skill

import (
	"database/sql"

	"github.com/GuyARoss/project-orva/pkg/pgdb"
)

// Request used for pg sessions
type Request struct {
	Creds *pgdb.DbCreds
}

// Skill skill type
type Skill struct {
	ID          string
	Endpoint    string
	AccessLevel int32
}

// FindByID finds a skill by ID
func (req *Request) FindByID(id string) (*Skill, error) {
	db := pgdb.CreateSession(req.Creds)

	defer db.Close()

	sqlQuery := `select * from skills where id = $1`
	row := db.QueryRow(sqlQuery, id)

	skill := &Skill{}

	switch err := row.Scan(&skill.ID, &skill.Endpoint, &skill.AccessLevel); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		return skill, nil
	default:
		panic(err)
	}
}

// Create creates a new skill
func (req *Request) Create(skill *Skill) error {
	db := pgdb.CreateSession(req.Creds)
	defer db.Close()

	sqlQuery := `insert into skills VALUES ($1, $2, $3)`
	_, err := db.Exec(sqlQuery, skill.ID, skill.Endpoint, skill.AccessLevel)

	return err
}

// Update updates a skill
func (req *Request) Update(skill *Skill) error {
	db := pgdb.CreateSession(req.Creds)
	defer db.Close()

	sqlQuery := `update skills set endpoint = $1, accesslevel = $2 where id = $3`
	_, err := db.Exec(sqlQuery, skill.Endpoint, skill.AccessLevel, skill.ID)

	return err
}

// Delete deletes a skill
func (req *Request) Delete(skill *Skill) error {
	db := pgdb.CreateSession(req.Creds)
	defer db.Close()

	sqlQuery := `delete from skills where id = $1`
	_, err := db.Exec(sqlQuery, skill.ID)

	return err
}
