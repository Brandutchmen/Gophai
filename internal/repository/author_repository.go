package repository

import (
	"app/internal/graph/model"
	"context"
)

type AuthorRepository interface {
	Insert(ctx context.Context, author model.Author) (model.Author, error)
	Update(ctx context.Context, author model.Author) (model.Author, error)
	Delete(ctx context.Context, id int64) (model.Author, error)
	List(ctx context.Context) ([]model.Author, error)
	Find(ctx context.Context, id int64) (model.Author, error)
}
