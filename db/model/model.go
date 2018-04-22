package model

type Model interface {
	CreateComment(comment string) error

}

// Comment provides comment body
type Comment struct {
	string body
}