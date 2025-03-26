package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("record not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Delete(ctx context.Context, postID int64) error
		Update(ctx context.Context, post *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		GetByPostID(ctx context.Context, postID int64) ([]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Comments: &CommentStore{db},
	}
}
