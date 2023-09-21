package repository

import (
	"app/internal/entity"
	"context"
)

type AuthorRepository interface {
	Insert(ctx context.Context, author entity.Author) (entity.Author, error)
	Update(ctx context.Context, author entity.Author) (entity.Author, error)
	Delete(ctx context.Context, id int64) (entity.Author, error)
	List(ctx context.Context) ([]entity.Author, error)
	Find(ctx context.Context, id int64) (entity.Author, error)
}
