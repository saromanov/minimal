package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/saromanov/minimal/db/model"
	"github.com/satori/go.uuid"
)

// Operations provides implementation of basic
// operations interface for db handling
type Operations struct {
	DB *sql.DB
}

// CreateAuthor provides inserting of new author
func (d *Operations) CreateAuthor(author *model.Author) (*uuid.UUID, error) {
	db := d.DB
	_, err := db.Exec(authorTable)
	if err != nil {
		log.Printf("%q: %s\n", err, authorTable)
		return nil, fmt.Errorf("unable to prepare table: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("unable to prepare inserting into db: %v", err)
	}
	stmt, err := tx.Prepare("insert into Author(Id, Name, Description) values(?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("unable to prepare inserting into db: %v", err)
	}
	uud := uuid.NewV4()
	defer stmt.Close()
	_, err = stmt.Exec(uud, author.Name, author.Description, author.Link, author.JoinedAt)
	if err != nil {
		return nil, fmt.Errorf("unable to insert into db: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("unable to commit into db: %v", err)
	}
	return &uud, nil
}
