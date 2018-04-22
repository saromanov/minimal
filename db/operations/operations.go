package operations

import (
	"github.com/saromanov/minimal/db/model"
	"github.com/satori/go.uuid"
)

// Operations provides all db operations
type Operations interface {
	CreateAuthor(*model.Author) error
	DeleteAuthor(uuid.UUID) error
	GetAuthor(uuid.UUID) (*model.Author, error)

	CreateComment(*model.Comment) error
	DeleteComment(uuid.UUID) error
	GetComment(uuid.UUID) (*model.Comment, error)
}
