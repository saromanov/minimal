package sqlite

import _ "github.com/mattn/go-sqlite3"

var (
	authorTable = `CREATE TABLE IF NOT EXISTS Author(
	Id BLOB PRIMARY KEY,
	Name varchar(1024) not null
	Description varchar(1024) not null
)`

	commentTable = `CREATE TABLE IF NOT EXISTS Comment(
		Id BLOB PRIMARY KEY,
		ThreadId BLOB not null,
		Body text not null,
		Author BLOB not null,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP not null,
		ReplyTo BLOB default null,
	)`
)
